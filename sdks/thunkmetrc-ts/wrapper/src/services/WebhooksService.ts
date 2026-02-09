
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    UpdateWebhooksRequest,
    WriteResponse,
} from '../models';

export class WebhooksService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * 
   * DELETE DeleteWebhooksBySubscriptionId
   */
  public async deleteWebhooksBySubscriptionId(subscriptionId: string, body?: any): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deleteWebhooksBySubscriptionId(subscriptionId, body,
        )
    );
  }

  /**
   * 
   * GET GetWebhooks
   */
  public async getWebhooks(body?: any): Promise<any> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getWebhooks(body,
        )
    );
  }

  /**
   * 
   * PUT UpdateDisableBySubscriptionId
   */
  public async updateWebhooksDisableBySubscriptionId(subscriptionId: string, body?: any): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateWebhooksDisableBySubscriptionId(subscriptionId, body,
        )
    );
  }

  /**
   * 
   * PUT UpdateEnableBySubscriptionId
   */
  public async updateWebhooksEnableBySubscriptionId(subscriptionId: string, body?: any): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateWebhooksEnableBySubscriptionId(subscriptionId, body,
        )
    );
  }

  /**
   * 
   * PUT UpdateWebhooks
   */
  public async updateWebhooks(body: UpdateWebhooksRequest): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateWebhooks(body,
        )
    );
  }
}

