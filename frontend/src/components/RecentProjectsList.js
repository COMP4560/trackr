import { useNavigate } from "react-router-dom";
import { useProjects } from "../hooks/useProjects";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import CircularProgress from "@mui/material/CircularProgress";
import CenteredBox from "../components/CenteredBox";
import ErrorIcon from "@mui/icons-material/Error";
import AccountTreeRoundedIcon from "@mui/icons-material/AccountTreeRounded";
import TextButton from "./TextButton";

const RecentProjectsList = () => {
  const [projects, , loading, error] = useProjects();
  const navigate = useNavigate();

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        mb: 3,
      }}
    >
      <Typography
        variant="h6"
        sx={{ flexGrow: 1, pb: 2, borderBottom: "2px solid #ededed", mb: 3 }}
      >
        Recent Projects
      </Typography>

      <Box
        sx={{
          display: "flex",
          flexDirection: "row",
          flexWrap: "wrap",
          gap: "15px",
        }}
      >
        {error ? (
          <CenteredBox>
            <ErrorIcon sx={{ fontSize: 50, mb: 2 }} />
            <Typography
              variant="h7"
              sx={{ mb: 10, userSelect: "none", textAlign: "center" }}
            >
              {error}
            </Typography>
          </CenteredBox>
        ) : loading ? (
          <CenteredBox>
            <CircularProgress />
          </CenteredBox>
        ) : projects.length === 0 ? (
          <CenteredBox>
            <Typography variant="h7" sx={{ color: "gray" }}>
              You currently have no projects.
            </Typography>
          </CenteredBox>
        ) : (
          projects.map((project) => (
            <Box
              key={project.id}
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
                <AccountTreeRoundedIcon
                  sx={{
                    fontSize: 35,
                    color: "#3887ff",
                    backgroundColor: "#bed8ff",
                    borderRadius: 1,
                    mt: "-20px",
                    mb: 1,
                  }}
                />

                <TextButton onClick={() => navigate("/projects/" + project.id)}>
                  <Typography variant="h6" sx={{ fontSize: "18px" }}>
                    {project.name}
                  </Typography>
                </TextButton>

                <Typography
                  variant="h7"
                  sx={{
                    fontSize: "13px",
                    mt: 1,
                    color: "#999999",
                    userSelect: "none",
                  }}
                >
                  QUICK LINKS
                </Typography>

                <TextButton
                  onClick={() => navigate("/projects/fields/" + project.id)}
                >
                  <Typography variant="h7" sx={{ fontSize: "13px", flex: 1 }}>
                    Fields
                  </Typography>

                  <Typography
                    variant="h7"
                    sx={{
                      fontSize: "13px",
                      background: "#00000011",
                      px: 1,
                      borderRadius: 100,
                    }}
                  >
                    {project.numberOfFields}
                  </Typography>
                </TextButton>

                <TextButton
                  onClick={() => navigate("/projects/api/" + project.id)}
                >
                  <Typography variant="h7" sx={{ fontSize: "13px", flex: 1 }}>
                    API
                  </Typography>
                </TextButton>

                <TextButton
                  onClick={() => navigate("/projects/settings/" + project.id)}
                >
                  <Typography variant="h7" sx={{ fontSize: "13px", flex: 1 }}>
                    Settings
                  </Typography>
                </TextButton>
              </Box>
            </Box>
          ))
        )}
      </Box>
    </Box>
  );
};

export default RecentProjectsList;
