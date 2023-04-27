import Container from "@mui/material/Container";
import OrganizationsCardBox from "../components/OrganizationsCardBox";
import ErrorBoundary from "../components/ErrorBoundary";
import CenteredBox from "../components/CenteredBox";
import ErrorIcon from "@mui/icons-material/Error";
import Typography from "@mui/material/Typography";
import LoadingBoundary from "../components/LoadingBoundary";
import CircularProgress from "@mui/material/CircularProgress";

// Page that you can open through the navbar
// Page includes a list of organisations currently created

const Organizations = () => {
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
        <Container sx={{ mt: 3, pb: 4 }}>
          <OrganizationsCardBox />
        </Container>
      </LoadingBoundary>
    </ErrorBoundary>
  );
};
export default Organizations;
