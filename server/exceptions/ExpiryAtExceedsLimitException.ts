import { HttpException } from './HttpException';

export class ExpiryAtExceedsLimitException extends HttpException {
  constructor(expiryAt: number) {
    super(400, `Specified expiry date "${new Date(expiryAt)}" exceeds maximum limit of 5 years.`);
  }
}
