import { ProjectRouteContext } from "../../routes/ProjectRoute";
import { useContext, lazy } from "react";
import { useFields } from "../../hooks/useFields";
import Box from "@mui/material/Box";
import CircularProgress from "@mui/material/CircularProgress";
import Typography from "@mui/material/Typography";
import ErrorIcon from "@mui/icons-material/Error";
import CenteredBox from "../CenteredBox";
import VisualizationMenuButton from "../VisualizationMenuButton";
import ErrorBoundary from "../ErrorBoundary";
import LoadingBoundary from "../LoadingBoundary";

const PieChartSubView = lazy(() => import("./PieChartSubView"));
const PieChartView = ({ visualizationType, visualization, metadata }) => {
  const { fieldId } = visualization;

  const projectId = useContext(ProjectRouteContext);
  const fields = useFields(projectId);
  const fieldName = fields.find((field) => field.id === fieldId)?.name;

  return (
    <>
      <Box
        sx={{
          display: "flex",
          flexDirection: "row",
          alignItems: "center",
          py: 1.5,
          px: 2,
          borderBottom: "1px solid #0000001f",
        }}
      >
        <Box sx={{ flexGrow: 1 }}>{fieldName}</Box>
        <Box>
          <VisualizationMenuButton
            visualizationType={visualizationType}
            visualization={visualization}
            metadata={metadata}
          />
        </Box>
      </Box>

      <ErrorBoundary
        fallback={({ error }) => (
          <CenteredBox>
            <ErrorIcon sx={{ fontSize: 50, mb: 1.5 }} />
            <Typography variant="h7" sx={{ userSelect: "none", mb: 2 }}>
              {error}
            </Typography>
          </CenteredBox>
        )}
      >
        <LoadingBoundary
          fallback={
            <CenteredBox>
              <CircularProgress />
            </CenteredBox>
          }
        >
          <PieChartSubView visualization={visualization} metadata={metadata} />
        </LoadingBoundary>
      </ErrorBoundary>
    </>
  );
};

export default PieChartView;
