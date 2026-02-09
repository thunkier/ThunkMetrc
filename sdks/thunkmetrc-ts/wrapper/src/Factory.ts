import { MetrcClient, type Logger } from '@thunkier/thunkmetrc-client';
import { MetrcWrapper } from './index';
import { MetrcRateLimiter, RateLimiterConfig } from './RateLimiter';

export interface MetrcFactoryConfig {
    maxConcurrentRequests?: number;
    logger?: Logger;
    timeout?: number;
}

export class MetrcFactory {
  private sharedLimiter: MetrcRateLimiter;

  constructor(private config: MetrcFactoryConfig = {}) {
    const limConfig = new RateLimiterConfig();
    if (config.maxConcurrentRequests) {
        limConfig.maxConcurrentGetIntegrator = config.maxConcurrentRequests;
    }
    this.sharedLimiter = new MetrcRateLimiter(limConfig);
  }

  public create(baseUrl: string, vendorKey: string, userKey: string): MetrcWrapper {
     const client = new MetrcClient({
        baseUrl,
        vendorKey,
        userKey,
        timeout: this.config.timeout,
        logger: this.config.logger
     });
     
     return new MetrcWrapper(client, this.sharedLimiter);
  }
}
