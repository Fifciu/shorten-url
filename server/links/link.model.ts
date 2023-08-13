import * as mongoose from 'mongoose';
import { Link } from './link.interface';

const linkSchema = new mongoose.Schema({
  name: String,
  alias: String,
  originalUrl: String,
  createdAt: Date,
  expiryAt: Date
});

export const linkModel = mongoose.model<Link & mongoose.Document>('Link', linkSchema);
