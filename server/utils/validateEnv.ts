import { cleanEnv, str, port, num } from "envalid";

export function validateEnv() {
  cleanEnv(process.env, {
    MONGO_PASSWORD: str(),
    MONGO_PATH: str(),
    MONGO_USER: str(),
    PORT: port(),
    JWT_SECRET: str(),
    JWT_TTL: num(),
    BCRYPT_SALT_OR_ROUNDS: num()
  });
}