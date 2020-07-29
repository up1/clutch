import ListExperiments from "./list-experiments";
import { StartAbortExperiment, StartLatencyExperiment, StartMobileFaultInjectionExperiment } from "./start-experiment";

const register = function register() {
  return {
    developer: {
      name: "Lyft",
      contactUrl: "mailto:hello@clutch.sh",
    },
    path: "experimentation",
    group: "Experimentation",
    displayName: "Experimentation",
    routes: {
      listExperiments: {
        path: "list",
        displayName: "List Experiments",
        description: "List Experiments.",
        component: ListExperiments,
      },
      startAbortExperiment: {
        path: "startabort",
        displayName: "Start an Abort Experiment",
        description: "Start an Abort Experiment.",
        component: StartAbortExperiment,
      },
      startLatencyExperiment: {
        path: "startlatency",
        displayName: "Start a Latency Experiment",
        description: "Start a Latency Experiment.",
        component: StartLatencyExperiment,
      },
      startAbortExperiment: {
        path: "startabort",
        displayName: "Start an Abort Experiment",
        description: "Start an Abort Experiment.",
        component: StartAbortExperiment,
      },
      startMobileFaultInjectionExperiment: {
        path: "startMobileFaultInjectionExperiment",
        displayName: "Start a MFI Experiment",
        description: "Start a MFI Experiment",
        component: StartMobileFaultInjectionExperiment,
      },
    },
  };
};

export default register;
