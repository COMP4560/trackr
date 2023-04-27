import { useUpdateOrganization } from "../hooks/useUpdateOrganization";
import { useContext, useState } from "react";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Fade from "@mui/material/Fade";
import Alert from "@mui/material/Alert";
import TextField from "@mui/material/TextField";
import LoadingButton from "@mui/lab/LoadingButton";
import Moment from "react-moment";
import formatError from "../utils/formatError";
import {useOrganization} from "../hooks/useOrganization";
import {OrganizationRouteContext} from "../routes/OrganizationRoute";

const EditOrganization = () => {
  const [error, setError] = useState();
  const [success, setSuccess] = useState();
  const [updateOrganization, updateOrganizationContext] = useUpdateOrganization(); //make this useUpdateOrganization
  const organizationId = useContext(OrganizationRouteContext);
  const organization = useOrganization(organizationId);

  const handleSubmit = (event) => {
    event.preventDefault();
    const data = new FormData(event.currentTarget);

    updateOrganization({
      id: organization.id,
      name: data.get("name"),
      description: data.get("description"),
    })
      .then(() => {
        setSuccess("Organization settings updated successfully.");
        setError();
      })
      .catch((error) => {
        setError(formatError(error));
      });
  };

  return (
    <Container sx={{ mt: 3 }}>
      <Box
        sx={{
          display: "flex",
          flexDirection: "row",
          alignItems: "center",
          mb: 3,
        }}
      >
        <Typography
          variant="h5"
          sx={{ flexGrow: 1, display: "flex", alignItems: "center" }}
        >
          Settings
        </Typography>
      </Box>

      <Box
        component="form"
        onSubmit={handleSubmit}
        noValidate
        sx={{ mt: 3, display: "flex", flexDirection: "column" }}
      >
        {(error || success) && (
          <Fade in={error || success ? true : false}>
            <Alert
              severity={error ? "error" : "success"}
              sx={{ mb: 3, mt: -1 }}
            >
              {error || success}
            </Alert>
          </Fade>
        )}

        <TextField
          label="Name"
          name="name"
          error={error ? true : false}
          required
          defaultValue={organization.name}
        />
        <Typography variant="caption" sx={{ mt: 1, mb: 2.5, color: "gray" }}>
          The name of your organization used to identify it.
        </Typography>

        <TextField
          label="Description"
          name="description"
          required
          error={error ? true : false}
          multiline
          rows={4}
          defaultValue={organization.description}
        />
        <Typography variant="caption" sx={{ mt: 1, mb: 2.5, color: "gray" }}>
          The description used to briefly describe your organization.
        </Typography>

        <Divider />

        <Box
          sx={{
            display: "flex",
            flexDirection: "row",
            alignItems: "baseline",
          }}
        >
          <LoadingButton
            loading={updateOrganizationContext.isLoading}
            type="submit"
            variant="contained"
            disableElevation
            sx={{
              my: 2,
              mr: 1.5,
              maxWidth: 180,
              flexGrow: 1,
            }}
          >
            Save Changes
          </LoadingButton>

          <Typography
            variant="caption"
            sx={{
              mt: 1,
              mb: 3,
            }}
          >
            Last modified{" "}
            <Moment fromNow ago>
              {organization.updatedAt}
            </Moment>{" "}
            ago
          </Typography>
        </Box>
      </Box>
    </Container>
  );
};

export default EditOrganization;
