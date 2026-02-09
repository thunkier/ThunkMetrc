
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    Staged,
    Tag,
} from '../models';

export class TagsService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Returns a list of available package tags. NOTE: This is a premium endpoint.
   * Permissions Required:
   * - Manage Tags
   * GET GetAvailablePackage
   */
  public async getTagsAvailablePackage(body?: any, licenseNumber?: string): Promise<Tag[]> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTagsAvailablePackage(body,licenseNumber,
        )
    );
  }

  /**
   * Returns a list of available plant tags. NOTE: This is a premium endpoint.
   * Permissions Required:
   * - Manage Tags
   * GET GetAvailablePlant
   */
  public async getTagsAvailablePlant(body?: any, licenseNumber?: string): Promise<Tag[]> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTagsAvailablePlant(body,licenseNumber,
        )
    );
  }

  /**
   * Returns a list of staged tags. NOTE: This is a premium endpoint.
   * Permissions Required:
   * - Manage Tags
   * - RetailId.AllowPackageStaging Key Value enabled
   * GET GetStagedTags
   */
  public async getStagedTags(body?: any, licenseNumber?: string): Promise<Staged[]> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getStagedTags(body,licenseNumber,
        )
    );
  }
}

