use crate::client::MetrcClient;
use serde_json::Value;
use std::error::Error;

pub struct WasteMethodsClient<'a> {
    pub(crate) client: &'a MetrcClient,
}

impl<'a> WasteMethodsClient<'a> {
    /// GET GetWasteMethodsWasteMethods
    /// Retrieves all available waste methods.
    /// Parameters:
    pub async fn get_waste_methods_waste_methods(&self, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let path = format!("/wastemethods/v2");

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
}

