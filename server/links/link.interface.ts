export interface Link {
  _id: string;
  name: string;
  shortenUrl: string;
  originalUrl: string;
  createdAt: Date;
  expireAt: Date;
};
