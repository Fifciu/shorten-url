import { HttpException } from './HttpException';

export class NoPermissionsToTheLinkException extends HttpException {
  constructor(linkId: string) {
    super(403, `You don't have permissions to the link with id ${linkId}`);
  }
}
