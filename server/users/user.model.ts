import * as mongoose from 'mongoose';
import { UserDocument } from './user.interface';

export const userSchema = new mongoose.Schema<UserDocument>({
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

export const userModel = mongoose.model<UserDocument & mongoose.Document>('User', userSchema);

