import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useCreateProject } from "../hooks/useCreateProject";
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

const AddUserButton = ({ sx, menuItem, icon }) => {
  const navigate = useNavigate();
  const [errorDialogOpen, setErrorDialogOpen] = useState(false);
  const [createProject, createProjectContext] = useCreateProject();

  const handleOnClick = () => {
    createProject()
      .then((result) => {
        navigate("/projects/settings/" + result.data.id);
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
          disabled={createProjectContext.isLoading}
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
            Add User
          </Typography>
        </MenuItem>
      ) : (
        <LoadingButton
          sx={sx}
          startIcon={icon && <AddRoundedIcon />}
          loading={createProjectContext.isLoading}
          type="submit"
          variant="contained"
          onClick={handleOnClick}
          disableElevation
        >
          Add User
        </LoadingButton>
      )}

      <Dialog open={errorDialogOpen} onClose={handleCloseErrorDialog}>
        <DialogTitle>Error</DialogTitle>
        <DialogContent>
          <DialogContentText>
            {formatError(createProjectContext.error)}
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

export default AddUserButton;
