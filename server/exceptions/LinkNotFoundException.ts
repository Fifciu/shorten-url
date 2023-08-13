import { HttpException } from './HttpException';

export class LinkNotFoundException extends HttpException {
  constructor(alias: string) {
    super(400, `URL shortened with the requested alias "${alias}" doesn't exist`);
  }
}
