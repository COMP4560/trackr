import { useState, useEffect } from "react";
import ProjectsAPI from "../api/ProjectsAPI";

export const useProject = (projectId) => {
  const [project, setProject] = useState();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState();

  useEffect(() => {
    ProjectsAPI.getProject(projectId)
      .then((result) => {
        setLoading(false);
        setError();
        setProject(result.data);
      })
      .catch((error) => {
        if (error?.response?.data?.error) {
          setError(error.response.data.error);
        } else {
          setError("Failed to load project: " + error.message);
        }
      });

  }, [projectId]);

  return [project, setProject, loading, error];
};
