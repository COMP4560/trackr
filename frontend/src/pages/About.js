import { Suspense, lazy } from "react";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import Link from "@mui/material/Link";
import CenteredBox from "../components/CenteredBox";
import CircularProgress from "@mui/material/CircularProgress";
import Logo from "../components/Logo";

const LicenseList = lazy(() => import("../components/LicenseList"));

const About = () => {
  return (
    <Box sx={{ display: "flex", flexDirection: "column" }}>
      <Box
        sx={{
          display: "flex",
          flexDirection: "row",
          alignItems: "center",
          mb: 1,
        }}
      >
        <Logo
          sxIcon={{ fontSize: 90, mr: 2, ml: -1.1 }}
          sxText={{ fontSize: 66 }}
        />
      </Box>
      <Typography variant="h7" sx={{ mb: 2, color: "#707070" }}>
        A platform created as part of the course{" "}
        <Link href="https://sci.umanitoba.ca/cs/wp-content/uploads/sites/3/2022/10/comp4560.pdf">
          COMP 4560 – Industrial Project
        </Link>{" "}
        from the{" "}
        <Link href="https://sci.umanitoba.ca/cs/">
          Department of Computer Science, University of Manitoba
        </Link>{" "}
        designed to make it easy for makers to store and visualize data recorded
        from their internet-connected devices like Arduinos and Raspberry Pis.
      </Typography>

      <Typography variant="h5">Contributors</Typography>
      <Typography variant="h7" sx={{ mb: 1.5, color: "#707070" }}>
        These are the students who have helped design, build, and shape the
        platform:
      </Typography>

      <Box
        sx={{
          borderRadius: 1,
          background: "#e5e5e5",
          width: "fit-content",
          px: 2,
          py: 0.5,
          fontSize: 15,
          mb: 0.5,
        }}
      >
        Fall 2022
      </Box>
      <List
        sx={{
          mb: 1,
          listStyleType: "disc",
          pl: 2.5,
          "& .MuiListItem-root": {
            display: "list-item",
            py: 0.3,
          },
        }}
      >
        <ListItem>Aidan Andrew-Hodgert</ListItem>
        <ListItem>Michael Kolisnyk</ListItem>
        <ListItem>Vlad Reinis</ListItem>
      </List>

      <Typography variant="h5">Licenses</Typography>
      <Typography variant="h7" sx={{ color: "#707070" }}>
        These are the licenses for the libraries that we use:
      </Typography>

      <Suspense
        fallback={
          <CenteredBox sx={{ mt: 1 }}>
            <CircularProgress />
          </CenteredBox>
        }
      >
        <LicenseList />
      </Suspense>
    </Box>
  );
};

export default About;
