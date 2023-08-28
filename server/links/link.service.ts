import { userModel } from '../users/user.model';
import { CreateLinkDto, UpdateLinkDto } from './link.dto';
import { linkModel } from './link.model';

class LinkService {
  private linkModel = linkModel;
  private userModel = userModel;

  public async aliasExists(alias: string) {
    const aliasExists = await this.linkModel.findOne({ alias });
    return Boolean(aliasExists);
  }

  public async getLinkByAlias(alias: string) {
    const link = await this.linkModel.findOne({ alias });
    return link;
  }

  public expiryAtExceedsLimit(expireAt: number, limit: number = 1000 * 60 * 60 * 24 * 365 * 5) {
    const now = Date.now();
    return expireAt - now > limit;
  }

  /**
   * Creates a link in database and attaches its ID to user's links.
   * @param createLinkDto 
   * @param userId 
   * @returns 
   */
  public async createLink(createLinkDto: CreateLinkDto, userId: string) {
    const link = new this.linkModel(createLinkDto);
    const user = await this.userModel.findById(userId);
    user!.links = [...(user?.links || []), link._id];
    await user!.save();
    const savedLink = await link.save();
    return savedLink;
  }

  public buildAlias(len = 10) {
    const chars = 'abcdefghijklmnopqrstuvwxyzABCDEFGHJKLMNOPQRSTUVWXYZ'.split('');
    let alias = '';
    for (let i = 0; i < len; i++) {
      alias += chars[Math.floor(Math.random() * chars.length)];
    }
    return alias;
  }

  public async linkExists(linkId: string) {
    return await this.linkModel.exists({ _id: linkId });
  }

  public async deleteLink(linkId: string, userId: string) {
    await this.linkModel.deleteOne({ _id: linkId });
    const updatedUser = await this.userModel.findByIdAndUpdate(userId, { $pull: { links: linkId } });
    return updatedUser;
  }

  public async patchLink(linkId: string, updateLinkDto: UpdateLinkDto) {
    const updatedLink = await this.linkModel.findByIdAndUpdate({ _id: linkId }, updateLinkDto, { new: true });
    return updatedLink;
  }
}

export const linkService = new LinkService();
