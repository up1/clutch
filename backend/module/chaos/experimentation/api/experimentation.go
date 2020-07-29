package api

// <!-- START clutchdoc -->
// description: Experimentation Framework Service implementation. Supports a CRUD API for managing experiments.
// <!-- END clutchdoc -->

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	mobilefaultinjectionv1 "github.com/lyft/clutch/backend/api/chaos/mobilefaultinjection/v1"

	//experimentationv1 "github.com/lyft/clutch-preview/backend/api/chaos/experimentation/v1"
	serverexperimentation "github.com/lyft/clutch/backend/api/chaos/serverexperimentation/v1"
	"reflect"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/uber-go/tally"
	"go.uber.org/zap"
	"golang.org/x/net/context"

	experimentation "github.com/lyft/clutch/backend/api/chaos/experimentation/v1"
	"github.com/lyft/clutch/backend/module"
	"github.com/lyft/clutch/backend/service"
	"github.com/lyft/clutch/backend/service/chaos/experimentation/experimentstore"
)

const (
	Name = "clutch.module.chaos.experimentation.api"
)

// Service contains all dependencies for the API service.
type Service struct {
	experimentStore       experimentstore.ExperimentStore
	converterRegistry     *ConverterRegistry
	logger                *zap.SugaredLogger
	createExperimentsStat tally.Counter
	getExperimentsStat    tally.Counter
	deleteExperimentsStat tally.Counter
}

type ConverterRegistry struct {
	services     map[string]func(experiment *experimentation.Experiment)experimentation.ExperimentListViewModel // map of types to services.
	serviceTypes []reflect.Type           // keep an ordered slice of registered service types.
}

func NewServiceRegistry() *ConverterRegistry {
	return &ConverterRegistry{
		services: make(map[string]func(experiment *experimentation.Experiment)experimentation.ExperimentListViewModel),
	}
}

func castToModel(experiment *experimentation.Experiment) experimentation.ExperimentListViewModel {
	testSpecification := &serverexperimentation.ServerTestSpecification{}
	ptypes.UnmarshalAny(experiment.GetTestConfig(), testSpecification)

	viewModel := experimentation.ExperimentListViewModel{}

	var clusterPair *serverexperimentation.ClusterPairTarget
	var description string

	switch testSpecification.GetConfig().(type) {
	case *serverexperimentation.ServerTestSpecification_Abort:
		abort := testSpecification.GetAbort()
		clusterPair = abort.GetClusterPair()
		description = fmt.Sprintf("%d abort %.2f%%", abort.GetHttpStatus(), abort.GetPercent())
	case *serverexperimentation.ServerTestSpecification_Latency:
		latency := testSpecification.GetLatency()
		clusterPair = latency.GetClusterPair()
		description = fmt.Sprintf("%dms delay %.2f%%", latency.GetDurationMs(), latency.GetPercent())
	default:
		break
	}

	viewModel.Identifier = string(experiment.GetId())
	viewModel.Targets = fmt.Sprintf("%s cluster to %s cluster",
		clusterPair.GetUpstreamCluster(),
		clusterPair.GetDownstreamCluster())
	viewModel.Type = "Server-side Fault Injection Test"
	viewModel.Description = description
	viewModel.Status = "Stopped"

	return viewModel
}

func castMobileFaultInjectionTestToModel(experiment *experimentation.Experiment) experimentation.ExperimentListViewModel {
	test := &mobilefaultinjectionv1.Test{}
	ptypes.UnmarshalAny(experiment.GetTestConfig(), test)

	viewModel := experimentation.ExperimentListViewModel{}
	viewModel.Identifier = string(experiment.GetId())
	viewModel.Targets = test.Endpoint
	viewModel.Type = "Mobile Fault Injection Tests"
	viewModel.Status = "Running"
	return viewModel
}

// New instantiates a Service object.
func New(_ *any.Any, logger *zap.Logger, scope tally.Scope) (module.Module, error) {
	store, ok := service.Registry[experimentstore.Name]
	if !ok {
		return nil, errors.New("could not find service")
	}

	experimentStore, ok := store.(experimentstore.ExperimentStore)
	if !ok {
		return nil, errors.New("service was not the correct type")
	}

	apiScope := scope.SubScope("experimentation")
	return &Service{
		experimentStore:       experimentStore,
		converterRegistry:     NewServiceRegistry(),
		logger:                logger.Sugar(),
		createExperimentsStat: apiScope.Counter("create_experiments"),
		getExperimentsStat:    apiScope.Counter("get_experiments"),
		deleteExperimentsStat: apiScope.Counter("delete_experiments"),
	}, nil
}

func (s *Service) Register(r module.Registrar) error {
	experimentation.RegisterExperimentsAPIServer(r.GRPCServer(), s)
	proto.RegisterType((*serverexperimentation.ServerTestSpecification)(nil),
		"clutch.chaos.serverexperimentation.v1.ServerTestSpecification")
	proto.RegisterType((*mobilefaultinjectionv1.Test)(nil),
		"clutch.chaos.mobilefaultinjection.v1.Test")
	s.converterRegistry.services["type.googleapis.com/clutch.chaos.serverexperimentation.v1.ServerTestSpecification"] = castToModel
	s.converterRegistry.services["type.googleapis.com/clutch.chaos.mobilefaultinjection.v1.Test"] = castMobileFaultInjectionTestToModel
	return r.RegisterJSONGateway(experimentation.RegisterExperimentsAPIHandler)
}

// CreateExperiments adds experiments to the experiment store.
func (s *Service) CreateExperiments(ctx context.Context, req *experimentation.CreateExperimentsRequest) (*experimentation.CreateExperimentsResponse, error) {
	s.createExperimentsStat.Inc(1)
	err := s.experimentStore.CreateExperiments(ctx, req.Experiments)
	if err != nil {
		return nil, err
	}

	return &experimentation.CreateExperimentsResponse{}, nil
}

// GetExperiments returns all experiments from the experiment store.
func (s *Service) GetExperiments(ctx context.Context, _ *experimentation.GetExperimentsRequest) (*experimentation.GetExperimentsResponse, error) {
	s.getExperimentsStat.Inc(1)
	experiments, err := s.experimentStore.GetExperiments(ctx)
	if err != nil {
		return &experimentation.GetExperimentsResponse{}, err
	}

	var items[] *experimentation.ExperimentListViewModel
	for _, experiment := range experiments {
		typeURL := experiment.TestConfig.TypeUrl
		viewModel := s.converterRegistry.services[typeURL](experiment)
		items = append(items, &viewModel)
	}

	return &experimentation.GetExperimentsResponse{Experiments: items}, nil
}

// DeleteExperiments deletes experiments from the experiment store.
func (s *Service) DeleteExperiments(ctx context.Context, req *experimentation.DeleteExperimentsRequest) (*experimentation.DeleteExperimentsResponse, error) {
	s.deleteExperimentsStat.Inc(1)
	err := s.experimentStore.DeleteExperiments(ctx, req.Ids)
	if err != nil {
		return nil, err
	}

	return &experimentation.DeleteExperimentsResponse{}, nil
}
