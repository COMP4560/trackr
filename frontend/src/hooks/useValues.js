import { useState, useEffect } from "react";
import ValuesAPI from "../api/ValuesAPI";

export const useValues = (fieldId, order, offset, limit) => {
  order = order || "asc";
  offset = offset || 0;
  limit = limit || 0;

  const [values, setValues] = useState([]);
  const [totalValues, setTotalValues] = useState(0);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState();

  useEffect(() => {
    ValuesAPI.getValues(fieldId, order, offset, limit)
      .then((result) => {
        setLoading(false);
        setError();
        setValues(result.data.values);
        setTotalValues(result.data.totalValues);
      })
      .catch((error) => {
        if (error?.response?.data?.error) {
          setError(error.response.data.error);
        } else {
          setError("Failed to load values: " + error.message);
        }
      });
  }, [fieldId, order, offset, limit]);

  return [values, totalValues, loading, error];
};
