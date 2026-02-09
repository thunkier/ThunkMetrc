
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    Facility,
} from '../models';

export class FacilitiesService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * This endpoint provides a list of facilities for which the authenticated user has access.
   * GET GetFacilities
   */
  public async getFacilities(body?: any): Promise<Facility[]> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getFacilities(body,
        )
    );
  }
}

