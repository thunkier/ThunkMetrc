from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class WebhooksService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def delete_webhooks_by_subscription_id(self, subscription_id: str, body: Any = None) -> Any:
        """
        
        DELETE /webhooks/v2/{subscriptionId}
        Parameters:
        - subscription_id (str): Path parameter subscriptionId
        """
        async def op():
            return await self.client.webhooks.delete_webhooks_by_subscription_id(subscription_id, body=body,
            )
        return await self.limiter.execute(None, False, op)

    async def get_webhooks(self, body: Any = None) -> Any:
        """
        
        GET /webhooks/v2
        Parameters:
        """
        async def op():
            return await self.client.webhooks.get_webhooks(body=body,
            )
        return await self.limiter.execute(None, True, op)

    async def update_webhooks_disable_by_subscription_id(self, subscription_id: str, body: Any = None) -> Any:
        """
        
        PUT /webhooks/v2/disable/{subscriptionId}
        Parameters:
        - subscription_id (str): Path parameter subscriptionId
        - body (Any): Request body
        """
        async def op():
            return await self.client.webhooks.update_webhooks_disable_by_subscription_id(subscription_id, body=body,
            )
        return await self.limiter.execute(None, False, op)

    async def update_webhooks_enable_by_subscription_id(self, subscription_id: str, body: Any = None) -> Any:
        """
        
        PUT /webhooks/v2/enable/{subscriptionId}
        Parameters:
        - subscription_id (str): Path parameter subscriptionId
        - body (Any): Request body
        """
        async def op():
            return await self.client.webhooks.update_webhooks_enable_by_subscription_id(subscription_id, body=body,
            )
        return await self.limiter.execute(None, False, op)

    async def update_webhooks(self, body: Optional['WebhooksUpdateWebhooksRequest'] = None) -> 'WriteResponse':
        """
        
        PUT /webhooks/v2
        Parameters:
        - body ('WebhooksUpdateWebhooksRequest'): Request body
        """
        async def op():
            return await self.client.webhooks.update_webhooks(body=body,
            )
        return await self.limiter.execute(None, False, op)

