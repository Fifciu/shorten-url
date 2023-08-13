import express from 'express';
import cookieParser from 'cookie-parser';
import * as bodyParser from 'body-parser';
import { Controller } from './interfaces/controller.interface';
import mongoose from 'mongoose';
import { errorMiddleware } from './middlewares/error.middleware';

export class App {
  public app: express.Application;
  constructor(apiControllers: Controller[], viewControllers: Controller[], public port: number) {
    this.app = express();

    this.initializeMiddlewares();
    this.initializeApiControllers(apiControllers);
    this.initializeViewControllers(viewControllers);
    this.initializeErrorHandling();
    this.connectToDatabase();
  }

  private initializeMiddlewares () {
    this.app.use(bodyParser.json());
    this.app.use(cookieParser());
  }

  private initializeApiControllers (controllers: Controller[]) {
    const router = express.Router();
    for (let controller of controllers) {
      console.log(`Initializing API controller ${controller.path}/*`);
      router.use(controller.path, controller.router);
    }
    this.app.use('/api/v1', router);
  }

  private initializeViewControllers (controllers: Controller[]) {
    for (let controller of controllers) {
      console.log(`Initializing View controller ${controller.path}/*`);
      this.app.use(controller.path, controller.router);
    }
    this.app.use('/page-not-found', (request, response) => {
      response.status(404).send('Page not found');
    });
  }

  private initializeErrorHandling () {
    this.app.use(errorMiddleware);
  }

  private connectToDatabase () {
    const {
      MONGO_USER,
      MONGO_PASSWORD,
      MONGO_PATH
    } = process.env;

    mongoose.connect(`mongodb://${MONGO_USER}:${MONGO_PASSWORD}${MONGO_PATH}`);
  }

  public listen () {
    this.app.listen(this.port, () => {
      console.log(`Server is running on port ${this.port}`);
    });
  }
}
