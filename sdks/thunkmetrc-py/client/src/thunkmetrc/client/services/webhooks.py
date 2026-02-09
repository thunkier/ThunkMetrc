from typing import Any, Optional, Dict, List, Union
import urllib.parse

class WebhooksClient:
    def __init__(self, client):
        self.client = client
    def delete_webhooks_by_subscription_id(self, subscription_id: str) -> Any:
        """
        
        DELETE DeleteWebhooksBySubscriptionId
        Parameters:
        - subscription_id (str): Path parameter subscriptionId
        """
        query_params = {}
        path = f"/webhooks/v2/{urllib.parse.quote(subscription_id)}"
        return self.client._send("DELETE", path, body, query_params)
    def get_webhooks(self) -> Any:
        """
        
        GET GetWebhooks
        Parameters:
        """
        query_params = {}
        path = f"/webhooks/v2"
        return self.client._send("GET", path, body, query_params)
    def update_webhooks_disable_by_subscription_id(self, subscription_id: str, body: Any = None) -> Any:
        """
        
        PUT UpdateDisableBySubscriptionId
        Parameters:
        - subscription_id (str): Path parameter subscriptionId
        - body (Any): Request body
        """
        query_params = {}
        path = f"/webhooks/v2/disable/{urllib.parse.quote(subscription_id)}"
        return self.client._send("PUT", path, body, query_params)
    def update_webhooks_enable_by_subscription_id(self, subscription_id: str, body: Any = None) -> Any:
        """
        
        PUT UpdateEnableBySubscriptionId
        Parameters:
        - subscription_id (str): Path parameter subscriptionId
        - body (Any): Request body
        """
        query_params = {}
        path = f"/webhooks/v2/enable/{urllib.parse.quote(subscription_id)}"
        return self.client._send("PUT", path, body, query_params)
    def update_webhooks(self, body: Any = None) -> Any:
        """
        
        PUT UpdateWebhooks
        Parameters:
        - body (Any): Request body
        """
        query_params = {}
        path = f"/webhooks/v2"
        return self.client._send("PUT", path, body, query_params)

