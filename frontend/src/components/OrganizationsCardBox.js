import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
// import CircularProgress from "@mui/material/CircularProgress";
// import CenteredBox from "./CenteredBox";
// import ErrorIcon from "@mui/icons-material/Error";
// import ErrorBoundary from "./ErrorBoundary";
// import LoadingBoundary from "./LoadingBoundary";
import CreateOrganizationButton from "./CreateOrganizationButton";
// import SearchBar from "../components/SearchBar";

import OrganizationsCardList from "./OrganizationsCardList";

const OrganizationsCardBox = () => {
  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        mb: 3,
      }}
    >
      <Box
        sx={{
          display: "flex",
          flexDirection: "row",
          mb: 2,
          justifyContent: "space-between",
          pb: 2,
          borderBottom: "2px solid #ededed",
        }}
      >
        <Typography variant="h6">Organizations</Typography>
        <CreateOrganizationButton
          sx={{
            alignSelf: "flex-end",
            ml: 2,
          }}
        />
      </Box>

      <Box
        sx={{
          display: "flex",
          flexDirection: "row",
          flexWrap: "wrap",
          // originally gap was 15px, 30px now to align cards
          gap: "30px",
        }}
      >
        <OrganizationsCardList />
      </Box>
    </Box>
  );
};

export default OrganizationsCardBox;
