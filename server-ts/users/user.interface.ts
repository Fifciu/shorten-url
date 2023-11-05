import { ObjectId } from 'mongoose';
import { Link } from '../links/link.interface';

export interface User {
  _id: string,
  name: string;
  email: string;
  password: string;
  links: Link[]
};

export type UserDocument = Omit<User, 'links'> & { links: ObjectId[] };
