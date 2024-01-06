import { sendToApi } from "./base";

const MODULE_BASE_PATH = 'authentication'; // without / at the begginig

export const authenticationService = {
  async register(fullname: string, email: string, password: string) {
    return sendToApi(`${MODULE_BASE_PATH}/register`, {
      method: 'POST',
      credentials: 'same-origin',
      body: JSON.stringify({
        fullname,
        email,
        password
      })
    });
  },

  async login(email: string, password: string) {
    await sendToApi(`${MODULE_BASE_PATH}/login`, {
      method: 'POST',
      credentials: 'same-origin',
      body: JSON.stringify({
        email,
        password
      })
    });
  },

  async refresh() {
    await sendToApi(`${MODULE_BASE_PATH}/refresh`, {
      method: 'POST',
      credentials: 'same-origin',
    });
  },

  async logout() {
    await sendToApi(`${MODULE_BASE_PATH}/logout`, {
      method: 'POST'
    });
  },
};
