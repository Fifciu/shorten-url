import { HttpException } from './HttpException';

export class ExpiryAtExceedsLimitException extends HttpException {
  constructor(expiryAt: Date) {
    super(400, `Specified expiry date "${expiryAt}" exceeds maximum limit of 5 years.`);
  }
}
