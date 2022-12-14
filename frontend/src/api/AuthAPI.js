import axios from "axios";

class AuthAPI {
  static #BASE_URL = process.env.REACT_APP_API_PATH + "/api/auth";
  static QUERY_KEY = "auth";

  static isLoggedIn = () => {
    return axios.get(this.#BASE_URL + "/", { withCredentials: true });
  };

  static login = (email, password, rememberMe) => {
    return axios.post(
      this.#BASE_URL + "/login",
      { email, password, rememberMe },
      { withCredentials: true }
    );
  };

  static register = (email, password, firstName, lastName) => {
    return axios.post(
      this.#BASE_URL + "/register",
      { email, password, firstName, lastName },
      { withCredentials: true }
    );
  };

  static logout = () => {
    return axios.get(this.#BASE_URL + "/logout", { withCredentials: true });
  };
}
export default AuthAPI;
