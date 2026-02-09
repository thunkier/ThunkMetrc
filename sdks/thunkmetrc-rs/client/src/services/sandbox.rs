use crate::client::MetrcClient;
use serde_json::Value;
use std::error::Error;

pub struct SandboxClient<'a> {
    pub(crate) client: &'a MetrcClient,
}

impl<'a> SandboxClient<'a> {
    /// POST CreateIntegratorSetup
    /// This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - user_key (Option<String>): Filter by userKey
    pub async fn create_sandbox_integrator_setup(&self, user_key: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/sandbox/v2/integrator/setup");
        let mut query_params = Vec::new();
        if let Some(p) = user_key {
            query_params.push(format!("userKey={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
}

