import { Controller } from "../interfaces/controller.interface";
import express from 'express';
import { linkService } from '../links/link.service';
import { LinkNotFoundException } from '../exceptions/LinkNotFoundException';

export class RedirectController implements Controller {
  public path = '/c';
  public router = express.Router();
  private linkService = linkService;

  constructor() {
    this.initializeRoutes();
  }

  private initializeRoutes() {
    this.router.get('/:alias', this.redirect.bind(this));
  }

  private async redirect(request: express.Request, response: express.Response, next: express.NextFunction) {
    const alias = request.params.alias;
    const link = await this.linkService.getLinkByAlias(alias);
    if (link) {
      response.redirect(link.originalUrl);
    } else {
      response.redirect('/page-not-found');
    }
  }

}
