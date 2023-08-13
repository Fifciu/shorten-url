import { HttpException } from './HttpException';

export class UrlWithThatAliasAlreadyExistsException extends HttpException {
  constructor(alias: string) {
    super(400, `URL shortened with the following alias "${alias}" already exists`);
  }
}
