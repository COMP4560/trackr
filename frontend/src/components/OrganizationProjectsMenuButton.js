import { useState } from "react";
// import { createElement } from "react";
// import { useNavigate } from "react-router-dom";
import IconButton from "@mui/material/IconButton";
import MoreVert from "@mui/icons-material/MoreVert";
import Menu from "@mui/material/Menu";
import Link from "@mui/material/Link";
import MenuItem from "@mui/material/MenuItem";
import ListItemText from "@mui/material/ListItemText";
import ListItemIcon from "@mui/material/ListItemIcon";
import DeleteIcon from "@mui/icons-material/Delete";
import Divider from "@mui/material/Divider";
import Dialog from "@mui/material/Dialog";
import Typography from "@mui/material/Typography";
// import ContentCopyIcon from "@mui/icons-material/ContentCopy";
// import DeleteProjectDialog from "./DeleteProjectDialog";
import GroupsIcon from "@mui/icons-material/Groups";

const OrganizationProjectsMenuButton = ({ project, noSettings }) => {
  const [anchorEl, setAnchorEl] = useState(null);
  const [dialogOpen, setDialogOpen] = useState(false);
  const [dialogElement] = useState();
  // const [dialogElement, setDialogElement] = useState();
  // const navigate = useNavigate();

  const openDropdownMenu = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const closeDropdownMenu = () => {
    setAnchorEl(null);
  };

  const closeDialog = () => {
    setDialogOpen(false);
  };

  // const openDeleteDialog = () => {
  //   setDialogOpen(true);
  //   setAnchorEl(null);
  //
  //   setDialogElement(
  //     createElement(DeleteProjectDialog, { project, onClose: closeDialog }, {})
  //   );
  // };
  //
  // const handleCopyAPIKey = () => {
  //   navigator.clipboard.writeText(project.apiKey);
  //   setAnchorEl(null);
  // };
  //
  // const handleSettingsProject = (event) => {
  //   event.preventDefault();
  //
  //   navigate("/projects/settings/" + project.id);
  // };

  return (
    <>
      <IconButton onClick={openDropdownMenu}>
        <MoreVert />
      </IconButton>

      <Menu
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={closeDropdownMenu}
      >
        {!noSettings && (
          <Link
            href={"/projects/settings/" + project.id}
            underline="none"
            sx={{ color: "unset" }}
          >
            <MenuItem>
              <ListItemIcon>
                <GroupsIcon fontSize="small" />
              </ListItemIcon>
              <ListItemText>Add user</ListItemText>
            </MenuItem>
          </Link>
        )}
        <Divider />
        <MenuItem>
          <ListItemIcon>
            <GroupsIcon fontSize="small" color="error" />
          </ListItemIcon>
          <ListItemText>
            <Typography color="error">Remove user</Typography>
          </ListItemText>
        </MenuItem>
        <MenuItem>
          <ListItemIcon>
            <DeleteIcon fontSize="small" color="error" />
          </ListItemIcon>
          <ListItemText>
            <Typography color="error"> Delete project</Typography>
          </ListItemText>
        </MenuItem>
      </Menu>

      <Dialog open={dialogOpen} onClose={closeDialog}>
        {dialogElement}
      </Dialog>
    </>
  );
};

export default OrganizationProjectsMenuButton;
