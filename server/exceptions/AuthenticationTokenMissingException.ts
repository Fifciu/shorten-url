import { HttpException } from './HttpException';

export class AuthenticationTokenMissingException extends HttpException {
  constructor() {
    super(400, 'Authentication token missing');
  }
}
