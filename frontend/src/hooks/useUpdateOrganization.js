import { useMutation, useQueryClient } from "react-query";
import OrganizationsAPI from "../api/OrganizationsAPI";

export const useUpdateOrganization = () => {
  const queryClient = useQueryClient();
  const mutation = useMutation(
    ({ id, name, description}) => {
      return OrganizationsAPI.updateOrganization(id, name, description);
    },
    {
      onSuccess: (result, { id, name, description }) => {
        queryClient.resetQueries(OrganizationsAPI.QUERY_KEY, { exact: true });
        queryClient.setQueryData([OrganizationsAPI.QUERY_KEY, id], (oldData) => {
          if (oldData) {
            oldData.data = {
              ...oldData.data,
              name,
              description,
              apiKey: result.data.apiKey,
              updatedAt: new Date().toISOString(),
            };
          }

          return oldData;
        });
      },
    }
  );

  return [mutation.mutateAsync, mutation];
};
