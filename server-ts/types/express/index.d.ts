import { UserDocument } from '../../users/user.interface';

declare module '@types/express-serve-static-core' {
  export interface Request {
     user?: UserDocument
  }
}
