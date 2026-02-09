
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    UnitOfMeasure,
} from '../models';

export class UnitsOfMeasureService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Retrieves all active units of measure.
   * GET GetActiveUnitsOfMeasure
   */
  public async getActiveUnitsOfMeasure(body?: any): Promise<PaginatedResponse<UnitOfMeasure>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getActiveUnitsOfMeasure(body,
        )
    );
  }

  /**
   * Retrieves all inactive units of measure.
   * GET GetInactiveUnitsOfMeasure
   */
  public async getInactiveUnitsOfMeasure(body?: any): Promise<PaginatedResponse<UnitOfMeasure>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getInactiveUnitsOfMeasure(body,
        )
    );
  }
}

