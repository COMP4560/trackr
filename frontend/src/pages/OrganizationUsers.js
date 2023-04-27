import ErrorBoundary from "../components/ErrorBoundary";
import CenteredBox from "../components/CenteredBox";
import ErrorIcon from "@mui/icons-material/Error";
import Typography from "@mui/material/Typography";
import LoadingBoundary from "../components/LoadingBoundary";
import CircularProgress from "@mui/material/CircularProgress";
import OrganizationUsersTable from "../components/OrganizationUsersTable";

const OrganizationUsers = () => {
  return (
    <ErrorBoundary
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
      <LoadingBoundary
        fallback={
          <CenteredBox>
            <CircularProgress />
          </CenteredBox>
        }
      >
        <OrganizationUsersTable />
      </LoadingBoundary>
    </ErrorBoundary>
  );
};
export default OrganizationUsers;
