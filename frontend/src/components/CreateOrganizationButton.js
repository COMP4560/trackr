import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useCreateOrganization } from "../hooks/useCreateOrganization";
import LoadingButton from "@mui/lab/LoadingButton";
import Button from "@mui/material/Button";
import MenuItem from "@mui/material/MenuItem";
import Typography from "@mui/material/Typography";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import DialogTitle from "@mui/material/DialogTitle";
import AddRoundedIcon from "@mui/icons-material/AddRounded";
import formatError from "../utils/formatError";

const CreateOrganizationButton = ({ sx, menuItem, icon }) => {
  const navigate = useNavigate();
  const [errorDialogOpen, setErrorDialogOpen] = useState(false);
  const [createOrganization, createOrganizationContext] =
    useCreateOrganization();

  const handleOnClick = () => {
    createOrganization()
      .then((result) => {
        navigate("/organizations/settings/" + result.data.id);
        menuItem();
      })
      .catch(() => {
        setErrorDialogOpen(true);
      });
  };

  const handleCloseErrorDialog = () => {
    setErrorDialogOpen(false);
    menuItem();
  };

  return (
    <>
      {menuItem ? (
        <MenuItem
          disabled={createOrganizationContext.isLoading}
          onClick={() => {
            handleOnClick();
          }}
          sx={{ py: 1, pr: 8, alignItems: "start" }}
        >
          <Typography
            variant="subtitle2"
            color="#2a2a2a"
            sx={{ fontWeight: 400 }}
          >
            Create Organization
          </Typography>
        </MenuItem>
      ) : (
        <LoadingButton
          sx={sx}
          startIcon={icon && <AddRoundedIcon />}
          loading={createOrganizationContext.isLoading}
          type="submit"
          variant="contained"
          onClick={handleOnClick}
          disableElevation
        >
          Create Organization
        </LoadingButton>
      )}

      <Dialog open={errorDialogOpen} onClose={handleCloseErrorDialog}>
        <DialogTitle>Error</DialogTitle>
        <DialogContent>
          <DialogContentText>
            {formatError(createOrganizationContext.error)}
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button autoFocus onClick={handleCloseErrorDialog}>
            Okay
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};

export default CreateOrganizationButton;
