use crate::client::MetrcClient;
use serde_json::Value;
use std::error::Error;

pub struct WebhooksClient<'a> {
    pub(crate) client: &'a MetrcClient,
}

impl<'a> WebhooksClient<'a> {
    /// DELETE DeleteWebhooksBySubscriptionId
    /// 
    /// Parameters:
    /// - subscription_id (str): Path parameter subscriptionId
    pub async fn delete_webhooks_by_subscription_id(&self, subscription_id: &str, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let path = format!("/webhooks/v2/{}", urlencoding::encode(subscription_id).as_ref());

        self.client.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// GET GetWebhooks
    /// 
    /// Parameters:
    pub async fn get_webhooks(&self, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let path = format!("/webhooks/v2");

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// PUT UpdateDisableBySubscriptionId
    /// 
    /// Parameters:
    /// - subscription_id (str): Path parameter subscriptionId
    /// - body (Option<&Value>): Request body
    pub async fn update_webhooks_disable_by_subscription_id(&self, subscription_id: &str, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let path = format!("/webhooks/v2/disable/{}", urlencoding::encode(subscription_id).as_ref());

        self.client.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// PUT UpdateEnableBySubscriptionId
    /// 
    /// Parameters:
    /// - subscription_id (str): Path parameter subscriptionId
    /// - body (Option<&Value>): Request body
    pub async fn update_webhooks_enable_by_subscription_id(&self, subscription_id: &str, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let path = format!("/webhooks/v2/enable/{}", urlencoding::encode(subscription_id).as_ref());

        self.client.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// PUT UpdateWebhooks
    /// 
    /// Parameters:
    /// - body (Option<&Value>): Request body
    pub async fn update_webhooks(&self, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let path = format!("/webhooks/v2");

        self.client.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
}

