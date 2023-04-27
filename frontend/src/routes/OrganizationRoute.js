import { createContext, useEffect, useRef } from "react";
import { useQueryClient } from "react-query";
import { useParams } from "react-router-dom";
import CenteredBox from "../components/CenteredBox";
import Typography from "@mui/material/Typography";
import CircularProgress from "@mui/material/CircularProgress";
import ErrorIcon from "@mui/icons-material/ErrorOutline";
import ErrorBoundary from "../components/ErrorBoundary";
import LoadingBoundary from "../components/LoadingBoundary";
import OrganizationsAPI from "../api/OrganizationsAPI";

export const OrganizationRouteContext = createContext();

const OrganizationRoute = ({ element }) => {
  const params = useParams();
  const organizationId = parseInt(params.organizationId);
  const errorBoundaryRef = useRef();
  const queryClient = useQueryClient();

  useEffect(() => {
    queryClient.prefetchQuery([OrganizationsAPI.QUERY_KEY, organizationId], () =>
      OrganizationsAPI.getOrganization(organizationId)
    );
    // queryClient.prefetchQuery([FieldsAPI.QUERY_KEY, projectId], () =>
    //   FieldsAPI.getFields(projectId)
    // );
    // queryClient.prefetchQuery([VisualizationsAPI.QUERY_KEY, projectId], () =>
    //   VisualizationsAPI.getVisualizations(projectId)
    // );

    errorBoundaryRef.current.reset();
  }, [queryClient, organizationId]);

  return (
    <ErrorBoundary
      ref={errorBoundaryRef}
      fallback={({ error }) => (
        <CenteredBox>
          <ErrorIcon sx={{ fontSize: 100, mb: 3 }} />
          <Typography
            variant="h5"
            sx={{ mb: 10, userSelect: "none", textAlign: "center" }}
          >
            {error}
          </Typography>
        </CenteredBox>
      )}
    >
      <OrganizationRouteContext.Provider value={organizationId}>

        <LoadingBoundary
          fallback={
            <CenteredBox>
              <CircularProgress />
            </CenteredBox>
          }
        >
          {element}
        </LoadingBoundary>
      </OrganizationRouteContext.Provider>
    </ErrorBoundary>
  );
};

export default OrganizationRoute;
