import { App } from './app';
import 'dotenv/config';
import { validateEnv } from './utils/validateEnv';
import { AuthenticationController } from './authentication/authentication.controller';

validateEnv();

const port = process.env.PORT || 8080;

new App([
  new AuthenticationController(),
], +port).listen();
