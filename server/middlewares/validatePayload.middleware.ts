import { plainToInstance } from "class-transformer";
import { ValidationError, validate } from "class-validator";
import { HttpException } from '../exceptions/HttpException';
import { Request, Response, NextFunction } from 'express';

export function validatePayload(type: any, skipMissingProperties = false) {
  return (req: Request, res: Response, next: NextFunction) => {
    validate(plainToInstance(type, req.body), { skipMissingProperties, whitelist: true, forbidNonWhitelisted: true })
      .then((errors: ValidationError[]) => {
        if (errors.length > 0) {
          const message = errors.map(
            (error: ValidationError) => error.constraints 
              ? Object.values(error.constraints).join(', ') 
              : 'Unknown error. Property "constraints" is missing'
          );
          next(new HttpException(400, message.join(', ')));
        } else {
          next();
        }
      })
  }
}