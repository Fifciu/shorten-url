import { Controller } from "../interfaces/controller.interface";
import express, { Request, Response, NextFunction } from 'express';
import { CreateLinkDto } from "./link.dto";
import { linkService } from "./link.service";
import { UrlWithThatAliasAlreadyExistsException } from "../exceptions/UrlWithThatAliasAlreadyExistsException";
import { ExpiryAtExceedsLimitException } from "../exceptions/ExpiryAtExceedsLimitException";
import { authMiddleware } from "../middlewares/auth.middleware";
import { validatePayload } from "../middlewares/validatePayload.middleware";
import { LinkNotFoundException } from '../exceptions/LinkNotFoundException';

export class LinkController implements Controller {
  public path = '/links';
  public router = express.Router();
  private linkService = linkService;

  constructor() {
    this.initializeRoutes();
  }

  private initializeRoutes() {
    this.router.get('/alias/:alias/originalUrl', this.getOriginalUrlByAlias.bind(this));
    this.router.all('/*', authMiddleware)
      .post('/', validatePayload(CreateLinkDto), this.createLink.bind(this))
    // this.router.delete('/:id', this.deleteLink.bind(this));
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

  private async getOriginalUrlByAlias(request: Request, response: Response, next: NextFunction) {
    const alias = request.params.alias;
    const link = await this.linkService.getLinkByAlias(alias);
    if (link) {
      response.send({
        orginalUrl: link.originalUrl
      });
    } else {
      next(new LinkNotFoundException(alias));
    }
  }
}