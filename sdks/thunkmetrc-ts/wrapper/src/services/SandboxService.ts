
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
} from '../models';

export class SandboxService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
   * POST CreateIntegratorSetup
   */
  public async createSandboxIntegratorSetup(body?: any, userKey?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createSandboxIntegratorSetup(body,userKey,
        )
    );
  }
}

