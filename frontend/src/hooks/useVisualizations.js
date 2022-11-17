import { useState, useEffect } from "react";
import VisualizationsAPI from "../api/VisualizationsAPI";

export const useVisualizations = (projectId) => {
  const [visualizations, setVisualizations] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState();

  useEffect(() => {
    VisualizationsAPI.getVisualizations(projectId)
      .then((result) => {
        setLoading(false);
        setError();
        setVisualizations(result.data.visualizations);
      })
      .catch((error) => {
        if (error?.response?.data?.error) {
          setError(error.response.data.error);
        } else {
          setError("Failed to load visualizations: " + error.message);
        }

        setLoading(false);
        setVisualizations([]);
      });

  }, [projectId]);

  return [visualizations, setVisualizations, loading, error];
};
