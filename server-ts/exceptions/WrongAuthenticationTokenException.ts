import { HttpException } from './HttpException';

export class WrongAuthenticationTokenException extends HttpException {
  constructor() {
    super(403, 'Wrong authentication token');
  }
}
