use crate::client::MetrcClient;
use serde_json::Value;
use std::error::Error;

pub struct FacilitiesClient<'a> {
    pub(crate) client: &'a MetrcClient,
}

impl<'a> FacilitiesClient<'a> {
    /// GET GetFacilities
    /// This endpoint provides a list of facilities for which the authenticated user has access.
    /// Parameters:
    pub async fn get_facilities(&self, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let path = format!("/facilities/v2");

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
}

