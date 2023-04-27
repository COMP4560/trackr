import { useQuery } from "react-query";
import OrganizationsAPI from "../api/OrganizationsAPI";

export const useOrganizations = () => {
  const { data } = useQuery(
    OrganizationsAPI.QUERY_KEY,
    OrganizationsAPI.getOrganizations,
    {
      suspense: true,
    }
  );

  return data.data.organizations;
};
