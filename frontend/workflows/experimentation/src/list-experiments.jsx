import React, { useState } from "react";
import { client, Row, Table } from "@clutch-sh/core";
import { Button, Container } from "@material-ui/core";

import { StartAbortExperiment, StartLatencyExperiment } from "./start-experiment";

function renderListItem(experiment) {
  return (
    <Row
      key={experiment.id}
      data={[
        experiment.identifier,
        experiment.targets,
        experiment.type,
        experiment.description,
        experiment.status,
      ]}
    />
  );
}

const ListExperiments = () => {
  const [experiments, setExperiments] = useState();

  if (!experiments) {
    client.post("/v1/experiments/get").then(response => {
      setExperiments(response?.data?.experiments || []);
    });

    return (
      <Table headings={["Identifier", "Targets", "Type", "Description", "Status"]} />
    );
  }

  return (
    <Container>
      <Table
        data={experiments}
        headings={["Identifier", "Targets", "Type", "Description", "Status"]}
      >
        {experiments.map(e => {
          return renderListItem(e);
        })}
      </Table>
      <Button onClick={StartAbortExperiment}>Start Abort Experiment</Button>
      <Button onClick={StartLatencyExperiment}>Start Latency Experiment</Button>
    </Container>
  );
};

export default ListExperiments;
