import { sendToApi } from "./base";

const MODULE_BASE_PATH = '/authentication';

export const authenticationService = {
  async register(fullname: string, email: string, password: string) {
    return sendToApi<undefined>(`${MODULE_BASE_PATH}/register`, {
      method: 'POST',
      body: JSON.stringify({
        fullname,
        email,
        password
      })
    });
  },

  async login(email: string, password: string) {
    await sendToApi<undefined>(`${MODULE_BASE_PATH}/login`, {
      method: 'POST',
      body: JSON.stringify({
        email,
        password
      })
    });
  },

  async refresh() {
    await sendToApi<undefined>(`${MODULE_BASE_PATH}/refresh`, {
      method: 'POST'
    });
  },

  async logout() {
    await sendToApi<undefined>(`${MODULE_BASE_PATH}/logout`, {
      method: 'POST'
    });
  },
};
