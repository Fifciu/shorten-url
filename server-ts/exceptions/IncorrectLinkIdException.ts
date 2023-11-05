import { HttpException } from './HttpException';

export class IncorrectLinkIdException extends HttpException {
  constructor(linkId: string) {
    super(400, `Passed incorrect link id: "${linkId}"`);
  }
}
