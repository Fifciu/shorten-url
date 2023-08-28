import { App } from './app';
import 'dotenv/config';
import { validateEnv } from './utils/validateEnv';
import { AuthenticationController } from './authentication/authentication.controller';
import { LinkController } from './links/link.controller';
import { RedirectController } from './redirect/redirect.controller';
import { UserController } from './users/user.controller';

validateEnv();

const port = process.env.PORT || 8080;

new App([
  new AuthenticationController(),
  new LinkController(),
  new UserController()
], [
  new RedirectController()
], +port).listen();
