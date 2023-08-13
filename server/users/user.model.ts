import * as mongoose from 'mongoose';
import { User } from './user.interface';

const userSchema = new mongoose.Schema({
  name: String,
  email: String,
  password: String,
  links: [
    {
      ref: 'Link',
      type: mongoose.Schema.Types.ObjectId
    }
  ]
});

export const userModel = mongoose.model<User & mongoose.Document>('User', userSchema);
