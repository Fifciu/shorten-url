import { sendToApi } from "./base";

const MODULE_BASE_PATH = 'authentication'; // without / at the begginig

export const authenticationService = {
  async register(fullname: string, email: string, password: string) {
    return sendToApi<undefined>(`${MODULE_BASE_PATH}/register`, {
      method: 'POST',
      credentials: 'same-origin',
      body: JSON.stringify({
        fullname,
        email,
        password
      })
    }, false);
  },

  async login(email: string, password: string) {
    await sendToApi<undefined>(`${MODULE_BASE_PATH}/login`, {
      method: 'POST',
      credentials: 'same-origin',
      body: JSON.stringify({
        email,
        password
      })
    }, false);
  },

  async refresh() {
    await sendToApi<undefined>(`${MODULE_BASE_PATH}/refresh`, {
      method: 'POST',
      credentials: 'same-origin',
    }, false);
  },

  async logout() {
    await sendToApi<undefined>(`${MODULE_BASE_PATH}/logout`, {
      method: 'POST'
    }, false);
  },
};
