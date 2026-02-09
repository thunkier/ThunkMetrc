
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    WasteMethod,
} from '../models';

export class WasteMethodsService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Retrieves all available waste methods.
   * GET GetWasteMethodsWasteMethods
   */
  public async getWasteMethodsWasteMethods(body?: any): Promise<WasteMethod[]> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getWasteMethodsWasteMethods(body,
        )
    );
  }
}

