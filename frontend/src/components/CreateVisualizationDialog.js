import { VisualizationTypes } from "./visualizations/Visualizations";
import { createElement, useState } from "react";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogTitle from "@mui/material/DialogTitle";
import DialogContentText from "@mui/material/DialogContentText";
import Box from "@mui/material/Box";
import ToggleButtonGroup from "@mui/material/ToggleButtonGroup";
import ToggleButton from "@mui/material/ToggleButton";
import Button from "@mui/material/Button";
import CreateFieldDialog from "./CreateFieldDialog";
import VisualizationsEditor from "./visualizations/VisualizationsEditor";

// This component is used to create a popup menu used for adding a new visualization to the project.

const CreateVisualizationDialog = ({ onClose }) => {
  const [primaryDialog, setPrimaryDialog] = useState();
  const [secondaryDialog, setSecondaryDialog] = useState();
  const [visualizationType, setVisualizationType] = useState();

  const handlePrimaryBack = () => {
    setPrimaryDialog();
  };

  const handleSecondaryBack = () => {
    setSecondaryDialog();
  };

  const handleAddField = () => {
    setSecondaryDialog(
      createElement(CreateFieldDialog, {
        onClose: onClose,
        onBack: handleSecondaryBack,
      })
    );
  };

  const handleSelectVisualization = () => {
    setPrimaryDialog(
      createElement(VisualizationsEditor, {
        onClose: onClose,
        onBack: handlePrimaryBack,
        onAddField: handleAddField,
        visualizationType,
      })
    );
  };

  const handleChangeVisualization = (event, visualizationType) => {
    setVisualizationType(visualizationType);
  };

  return (
    <>
      <Box sx={{ display: !secondaryDialog ? "block" : "none" }}>
        {primaryDialog}
      </Box>
      {secondaryDialog}
      {!secondaryDialog && !primaryDialog && (
        <>
          <DialogTitle sx={{ display: "flex", alignItems: "center" }}>
            New Visualization
          </DialogTitle>
          <DialogContent
            sx={{
              display: "flex",
              flexDirection: "column",
            }}
          >
            <DialogContentText sx={{ mb: 2 }}>
              You can visualize the data collected by your IoT devices by
              picking a unique visualization from the list below.
            </DialogContentText>

            <Box
              sx={{
                display: "flex",
                flexDirection: "row",
                gap: "12px",
                flexWrap: "wrap",
              }}
            >
              <ToggleButtonGroup
                fullWidth
                color="primary"
                value={visualizationType}
                exclusive
                onChange={handleChangeVisualization}
                sx={{ mr: 1 }}
              >
                {VisualizationTypes.map((visualizationType) => (
                  <ToggleButton
                    key={visualizationType.name}
                    value={visualizationType}
                    sx={{ p: 3, display: "flex", flexDirection: "column" }}
                  >
                    {createElement(visualizationType.icon, { sx: { mb: 1 } })}
                    {visualizationType.name}
                  </ToggleButton>
                ))}
              </ToggleButtonGroup>
            </Box>
          </DialogContent>
          <DialogActions sx={{ pb: 3, pr: 3 }}>
            <Button onClick={onClose}>Cancel</Button>
            <Button
              variant="contained"
              disableElevation
              autoFocus
              disabled={!visualizationType}
              onClick={handleSelectVisualization}
            >
              Next
            </Button>
          </DialogActions>
        </>
      )}
    </>
  );
};

export default CreateVisualizationDialog;
