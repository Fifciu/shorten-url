import * as mongoose from 'mongoose';
import { Link } from './link.interface';

const linkSchema = new mongoose.Schema<Link>({
  name: String,
  alias: String,
  originalUrl: String,
  createdAt: Date,
  expiryAt: Number
});

export const linkModel = mongoose.model<Link & mongoose.Document>('Link', linkSchema);
