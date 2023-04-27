import { useMutation } from "react-query";
import OrganizationsAPI from "../api/OrganizationsAPI";

export const useCreateOrganization = () => {
  const mutation = useMutation(() => OrganizationsAPI.createOrganization());

  return [mutation.mutateAsync, mutation];
};
