import { useQuery } from "react-query";
import OrganizationsAPI from "../api/OrganizationsAPI";

export const useOrganization = (organizationId) => {
  const { data } = useQuery(
    [OrganizationsAPI.QUERY_KEY, organizationId],
    () => OrganizationsAPI.getOrganization(organizationId),
    {
      suspense: true,
    }
  );

  return data.data;
};
