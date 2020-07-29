package api

// <!-- START clutchdoc -->
// description: Experimentation Framework Service implementation. Supports a CRUD API for managing experiments.
// <!-- END clutchdoc -->

import (
	"errors"
	"github.com/golang/protobuf/proto"
	serverexperimentationv1 "github.com/lyft/clutch/backend/api/chaos/serverexperimentation/v1"
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
	services     map[reflect.Type]Service // map of types to services.
	serviceTypes []reflect.Type           // keep an ordered slice of registered service types.
}

func NewServiceRegistry() *ConverterRegistry {
	return &ConverterRegistry{
		services: make(map[reflect.Type]Service),
	}
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
	proto.RegisterType((*serverexperimentationv1.ServerTestSpecification)(nil),
		"clutch.chaos.serverexperimentation.v1.ServerTestSpecification")
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

	return &experimentation.GetExperimentsResponse{Experiments: experiments}, nil
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
