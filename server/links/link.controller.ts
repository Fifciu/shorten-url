import { Controller } from "../interfaces/controller.interface";
import express, { Request, Response, NextFunction } from 'express';
import { CreateLinkDto } from "./link.dto";
import { linkService } from "./link.service";
import { UrlWithThatAliasAlreadyExistsException } from "../exceptions/UrlWithThatAliasAlreadyExistsException";
import { ExpiryAtExceedsLimitException } from "../exceptions/ExpiryAtExceedsLimitException";
import { authMiddleware } from "../middlewares/auth.middleware";
import { validatePayload } from "../middlewares/validatePayload.middleware";
import { LinkNotFoundException } from '../exceptions/LinkNotFoundException';
import { NoPermissionsToTheLinkException } from "../exceptions/NoPermissionsToTheLinkException";
import { IncorrectLinkIdException } from "../exceptions/IncorrectLinkIdException";

export class LinkController implements Controller {
  public path = '/links';
  public router = express.Router();
  private linkService = linkService;

  constructor() {
    this.initializeRoutes();
  }

  private initializeRoutes() {
    this.router.all('/*', authMiddleware)
      .post('/', validatePayload(CreateLinkDto), this.createLink.bind(this))
      .delete('/:id', this.deleteLink.bind(this));
    // this.router.patch('/:id', this.patchLink.bind(this));
  }

  private async createLink(request: Request, response: Response, next: NextFunction) {
    const linkBody: CreateLinkDto = request.body;
    if (linkBody.alias && await this.linkService.aliasExists(linkBody.alias)) {
      next(new UrlWithThatAliasAlreadyExistsException(linkBody.alias));
    } else {
      if (linkBody.expiryAt && this.linkService.expiryAtExceedsLimit(linkBody.expiryAt)) {
        next(new ExpiryAtExceedsLimitException(linkBody.expiryAt));
      } else {
        if (!linkBody.alias) {
          do {
            linkBody.alias = this.linkService.buildAlias();
          } while(await this.linkService.aliasExists(linkBody.alias));
        }
        const link = await this.linkService.createLink(linkBody, request.user!._id);
        response.send(link); 
      }
    }
  }

  private async deleteLink(request: Request, response: Response, next: NextFunction) {
    const linkId = request.params.id;
    let linkExists = false;
    try {
      linkExists = Boolean(await this.linkService.linkExists(linkId));
    } catch (err) {
      console.log(err);
      next(new IncorrectLinkIdException(linkId));
      return;
    }
    if (!linkExists) {
      next(new LinkNotFoundException(linkId));
    } else {
      if (!request.user?.links.some(link => link.toHexString() === linkId)) {
        next(new NoPermissionsToTheLinkException(linkId));
      } else {
        const user = await this.linkService.deleteLink(linkId, request.user!._id);
        request.user = user!;
        response.sendStatus(204);
      }
    }
  }
}