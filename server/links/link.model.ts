import * as mongoose from 'mongoose';
import { Link } from './link.interface';

const linkSchema = new mongoose.Schema({
  name: String,
  shortenUrl: String,
  originalUrl: String,
  createdAt: Date,
  expireAt: Date
});

export const linkModel = mongoose.model<Link & mongoose.Document>('Link', linkSchema);
