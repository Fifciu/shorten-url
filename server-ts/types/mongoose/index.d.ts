declare module 'mongoose' {
  namespace Schema {
    namespace Types {
      interface ObjectId {
        toHexString(): string;
      }
    }
  }
}
