export function isLowerCase(str: string) {
  return str === str.toLowerCase() &&
         str !== str.toUpperCase();
}

export function isUpperCase(str: string) {
  return str === str.toUpperCase() &&
         str !== str.toLowerCase();
}

export class FetchError extends Error {
  constructor(msg: string, public response: Response, public status: number) {
    super(msg);
  }
}

export const ERR_EMAIL_EXISTS = 'EMAIL_EXISTS';
export const ERR_WRONG_EMAIL_OR_PASSWORD = 'WRONG_EMAIL_OR_PASSWORD';
export const API_ERRORS = [ERR_EMAIL_EXISTS, ERR_WRONG_EMAIL_OR_PASSWORD] as const;

export const ERROR_MESSAGES: Record<typeof API_ERRORS[number], string> = {
  EMAIL_EXISTS: 'Email already exists',
  WRONG_EMAIL_OR_PASSWORD: 'Wrong email or password',
}