import { createElement } from "react";
import Table from "./Table";
import Graph from "./Graph";
import PieChart from "./PieChart";
import CenteredBox from "../CenteredBox";
import Typography from "@mui/material/Typography";
import ErrorIcon from "@mui/icons-material/ErrorOutline";

export const VisualizationTypes = [Table, Graph, PieChart];

export const createVisualizationElement = (visualization) => {
  const metadata = JSON.parse(visualization.metadata);

  for (let i = 0; i < VisualizationTypes.length; i++) {
    if (VisualizationTypes[i].name === metadata.name) {
      return createElement(
        VisualizationTypes[i].view,
        {
          visualization,
          visualizationType: VisualizationTypes[i],
          metadata,
        },
        {}
      );
    }
  }

  return (
    <CenteredBox>
      <ErrorIcon sx={{ fontSize: 50, mb: 1.5 }} />
      <Typography variant="h7" sx={{ userSelect: "none", mb: 2 }}>
        Failed to load visualization, invalid type: {metadata.name}
      </Typography>
    </CenteredBox>
  );
};
