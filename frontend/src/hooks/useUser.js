import { useState, useEffect } from "react";
import UsersAPI from "../api/UsersAPI";

export const useUser = () => {
  const [user, setUser] = useState();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState();

  useEffect(() => {
    UsersAPI.getUser()
      .then((result) => {
        setLoading(false);
        setError();
        setUser(result.data);
      })
      .catch((error) => {
        if (error?.response?.data?.error) {
          setError(error.response.data.error);
        } else {
          setError("Failed to load user: " + error.message);
        }

        setLoading(false);
        setUser();
      });
  }, []);

  return [user, setUser, loading, error];
};
