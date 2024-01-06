import { sendToApi } from "./base";

const MODULE_BASE_PATH = 'users';

export const usersService = {
  async changeEmail(current_password: string, new_email: string) {
    return sendToApi(`${MODULE_BASE_PATH}/email`, {
      method: 'PATCH',
      body: JSON.stringify({
        current_password,
        new_email,
      })
    });
  },

  async changePassword(current_password: string, new_password: string) {
    return sendToApi(`${MODULE_BASE_PATH}/password`, {
      method: 'PATCH',
      body: JSON.stringify({
        current_password,
        new_password,
      })
    });
  }
};
