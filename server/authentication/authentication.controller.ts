import { Controller } from "../interfaces/controller.interface";
import express, { Request, Response, NextFunction } from 'express';
import { validatePayload } from "../middlewares/validatePayload.middleware";
import { CreateUserDto } from "../users/user.dto";
import { userService } from '../users/user.service';
import { UserWithThatEmailAlreadyExistsException } from "../exceptions/UserWithThatEmailAlreadyExistsException";
import { DataStoredInToken } from '../interfaces/dataStoredInToken.interface';
import { User } from "../users/user.interface";
import * as jwt from 'jsonwebtoken';
import { LogInDto } from "./logIn.dto";
import { WrongCredentialsException } from '../exceptions/WrongCredentialsException';
import bcrypt from 'bcrypt';

interface TokenData {
  token: string;
  expiresIn: number;
}

export class AuthenticationController implements Controller {
  public path = '/authentication';
  public router = express.Router();
  private userService = userService;

  constructor() {
    this.initializeRoutes();
  }

  private initializeRoutes() {
    this.router.post('/register', validatePayload(CreateUserDto), this.registration.bind(this));
    this.router.post('/login', validatePayload(LogInDto), this.login.bind(this));
    this.router.post('/logout', this.logout.bind(this));
  }

  private createToken(user: User): TokenData {
    const expiresIn = Number(process.env.JWT_TTL!);
    const secret = process.env.JWT_SECRET!;
    const dataStoredInToken: DataStoredInToken = {
      _id: user._id
    };

    return {
      expiresIn,
      token: jwt.sign(dataStoredInToken, secret, { expiresIn })
    }
  }

  private createCookie({ token, expiresIn }: TokenData) {
    return `Authorization=${token}; HttpOnly; Max-Age=${expiresIn}`;
  }

  private async registration(request: Request, response: Response, next: NextFunction) {
    const userData: CreateUserDto = request.body;
    if (await this.userService.userWithEmailExists(userData.email)) {
      next(new UserWithThatEmailAlreadyExistsException(userData.email));
    } else {
      const user = await this.userService.createWithHashedPassword(userData);
      (user.password as any) = undefined; // TODO
      const tokenData = this.createToken(user);
      response.setHeader('Set-Cookie', [
        this.createCookie(tokenData)
      ]);
      response.send(user);
    }
  }

  private async login(request: Request, response: Response, next: NextFunction) {
    const logInData: LogInDto = request.body;
    const user = await this.userService.getUserByEmail(logInData.email);
    if (user) {
      const isPasswordMatching = await bcrypt.compare(logInData.password, user.password);
      if (isPasswordMatching) {
        (user.password as any) = undefined; // TODO
        const tokenData = this.createToken(user);
        response.setHeader('Set-Cookie', [
          this.createCookie(tokenData)
        ]);
        response.send(user);
      } else {
        next(new WrongCredentialsException());
      }
    } else {
      next(new WrongCredentialsException());
    }
  }

  private async logout(_: Request, response: Response) {
    response.setHeader('Set-Cookie', ['Authorization=;Max-age=0']);
    response.sendStatus(200);
  }
};
