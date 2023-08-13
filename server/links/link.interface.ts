export interface Link {
  _id: string;
  name: string;
  alias: string;
  originalUrl: string;
  createdAt: Date;
  expiryAt: Date;
};
