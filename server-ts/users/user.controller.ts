import express, { Request, Response, NextFunction } from 'express';
import { Controller } from "../interfaces/controller.interface";
import { authMiddleware } from "../middlewares/auth.middleware";
import { linkService } from '../links/link.service';

export class UserController implements Controller {
  public path = '/users'
  public router = express.Router();
  private link = linkService;
  constructor() {
    this.initializeRoutes();
  }

  private initializeRoutes() {
    this.router.all('/*', authMiddleware)
      .post('/me', this.getMe.bind(this))
      .post('/me/links', this.getMyLinks.bind(this))
  }

  private async getMe(request: Request, response: Response, next: NextFunction) {
    const user = request.user;
    (user!.links as any) = undefined;
    response.send(request.user);
  }

  private async getMyLinks(request: Request, response: Response, next: NextFunction) {
    const linksIds = request.user!.links;
    const links = await this.link.getLinksByIds(linksIds);
    response.send(links);
  }
};
