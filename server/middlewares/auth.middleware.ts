import { Request, Response, NextFunction } from 'express';
import { AuthenticationTokenMissingException } from '../exceptions/AuthenticationTokenMissingException';
import { WrongAuthenticationTokenException } from '../exceptions/WrongAuthenticationTokenException';
import * as jwt from 'jsonwebtoken';
import { DataStoredInToken } from '../interfaces/dataStoredInToken.interface';
import { userModel } from '../users/user.model';

export async function authMiddleware(request: Request, _: Response, next: NextFunction) {
  const cookies = request.cookies;
  if (cookies?.Authorization) {
    const secret = process.env.JWT_SECRET!;
    try {
      const verificationResponse = jwt.verify(cookies.Authorization, secret) as DataStoredInToken;
      const id = verificationResponse._id;
      const user = await userModel.findById(id);
      if (user) {
        request.user = user;
        next();
      } else {
        next(new WrongAuthenticationTokenException());
      }
    } catch (err) {
      next(new WrongAuthenticationTokenException());
    }
  } else {
    next(new AuthenticationTokenMissingException());
  }
}
