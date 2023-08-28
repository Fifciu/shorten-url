import express, { Request, Response, NextFunction } from 'express';
import { Controller } from "../interfaces/controller.interface";
import { userModel } from "./user.model";
import { linkModel } from "../links/link.model";
import { authMiddleware } from "../middlewares/auth.middleware";

export class UserController implements Controller {
  public path = '/users'
  public router = express.Router();
  private user = userModel;
  private link = linkModel;
  constructor() {
    this.initializeRoutes();
  }

  private initializeRoutes() {
    this.router.all('/*', authMiddleware)
      .post('/me', this.getMe.bind(this))
  }

  private async getMe(request: Request, response: Response, next: NextFunction) {
    response.send(request.user);
  }
};
