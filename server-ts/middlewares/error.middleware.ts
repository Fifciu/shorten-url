import { Request, Response, NextFunction } from 'express';
import { HttpException } from '../exceptions/HttpException';

export function errorMiddleware(error: HttpException, _: Request, response: Response, next: NextFunction) {
  const status = error.status || 500;
  const message = error.message || 'Something went wrong';
  response.status(status).send({
    message
  });
}
