import { useState, useEffect } from "react";
import ProjectsAPI from "../api/ProjectsAPI";

export const useProjects = () => {
  const [projects, setProjects] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState();

  useEffect(() => {
    ProjectsAPI.getProjects()
      .then((result) => {
        setLoading(false);
        setError();
        setProjects(result.data.projects);
      })
      .catch((error) => {
        if (error?.response?.data?.error) {
          setError(error.response.data.error);
        } else {
          setError("Failed to load projects: " + error.message);
        }
      });
  }, []);

  return [projects, setProjects, loading, error];
};
