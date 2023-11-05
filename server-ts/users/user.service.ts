import { CreateUserDto } from "./user.dto";
import { userModel } from "./user.model";
import bcrypt from 'bcrypt';

export class UserService {
  private userModel = userModel;

  /**
   * Creates new user with hashed password.
   * @param userData {CreateUserDto}
   * @returns Created user
   */
  public async createWithHashedPassword(userData: CreateUserDto) {
    const hashedPassword = await bcrypt.hash(userData.password, Number(process.env.BCRYPT_SALT_OR_ROUNDS!));
    const user = await this.userModel.create({
      ...userData,
      password: hashedPassword
    })
    return user;
  }

  /**
   * Checks if user with given email exists.
   * @param email 
   * @returns 
   */
  public async userWithEmailExists(email: string) {
    const user = await this.userModel.findOne({ email });
    return Boolean(user);
  }

  /**
   * Returns user with given email.
   * @param email 
   * @returns 
   */
  public async getUserByEmail(email: string) {
    const user = await this.userModel.findOne({ email });
    return user;
  }
}

export const userService = new UserService();