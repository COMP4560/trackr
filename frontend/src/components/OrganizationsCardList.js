import { useNavigate } from "react-router-dom";
import { useOrganizations } from "../hooks/useOrganizations";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import CenteredBox from "./CenteredBox";
import GroupsIcon from "@mui/icons-material/Groups";
import Card from "@mui/material/Card";
import OrganizationsMenuButton from "./OrganizationsMenuButton";
import TextButton from "./TextButton";

const OrganizationsCardList = () => {
  const organizations = useOrganizations();
  const navigate = useNavigate();

  return organizations.length === 0 ? (
    <CenteredBox>
      <Typography variant="h7" sx={{ color: "gray" }}>
        You currently have no organizations.
      </Typography>
    </CenteredBox>
  ) : (
    organizations.map((organization) => (
      <Card
        key={organization.id}
        sx={{
          display: "flex",
          flexDirection: "row",
          minWidth: {
            xs: "100%",
            sm: "49%",
            md: "23%",
          },
          maxWidth: {
            xs: "unset",
            sm: "unset",
            md: "200px",
          },
          flex: 1,
          borderRadius: 1,
          background: "#ebf3ff",
          boxShadow: "0 1px 1px 1px rgb(9 30 66 / 10%)",
        }}
      >
        {/*<CardActionArea>*/}
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            background: "white",
            mt: 5,
            px: 2,
            pb: 3,
            flex: 1,
          }}
        >
          <GroupsIcon
            sx={{
              fontSize: 35,
              color: "#3887ff",
              //   backgroundColor: "#bed8ff",
              borderRadius: 1,
              mt: "-35px",
              mb: 1,
            }}
          />
          {/* Name of the organization */}
          <Typography variant="h6" sx={{ fontSize: "18px", mb: 3 }}>
            {organization.name}
          </Typography>

          <TextButton
            onClick={() => navigate("/organizations/users/" + organization.id)}
          >
            <Typography variant="h6" sx={{ fontSize: "13px", flex: 1 }}>
              {/* Users: {organization.numMembers} */}
              Users: 0
            </Typography>
          </TextButton>
          <Box sx={{ display: "flex", flexDirection: "row" }}>
            <Typography
              onClick={() =>
                navigate("/organizations/projects/" + organization.id)
              }
              variant="h6"
              sx={{
                fontSize: "13px",
                flex: 1,
                borderRadius: 1,
                userSelect: "none",
                py: 0.2,
                "&:hover": { background: "#ebf3ff", cursor: "pointer" },
              }}
            >
              {/* Projects: {organization.numProjects} */}
              Projects: 0
            </Typography>
            <OrganizationsMenuButton
              sx={{ zIndex: "" }}
              organization={organization}
            />
          </Box>
        </Box>
      </Card>
    ))
  );
};

export default OrganizationsCardList;
