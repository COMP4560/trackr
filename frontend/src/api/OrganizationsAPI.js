import axios from "axios";

class OrganizationsAPI {
  static #BASE_URL = process.env.REACT_APP_API_PATH + "/api/organizations";
  static QUERY_KEY = "organizations";

  static getOrganizations = () => {
    return axios.get(this.#BASE_URL + "/", {
      withCredentials: true,
    });
  };

  static getOrganization = (organizationId) => {
    return axios.get(this.#BASE_URL + "/" + organizationId, {
      withCredentials: true,
    });
  };

  static createOrganization = () => {
    return axios.post(this.#BASE_URL + "/", {}, { withCredentials: true });
  };

  // INPORTANT
  // change variables to relevant variables to organizations
  static updateOrganization = (id, name, description) => {
    return axios.put(
      this.#BASE_URL + "/",
      { id, name, description },
      { withCredentials: true }
    );
  };

  static deleteOrganization = (organizationId) => {
    return axios.delete(this.#BASE_URL + "/" + organizationId, {
      withCredentials: true,
    });
  };
}

export default OrganizationsAPI;
