use crate::client::MetrcClient;
use serde_json::Value;
use std::error::Error;

pub struct TagsClient<'a> {
    pub(crate) client: &'a MetrcClient,
}

impl<'a> TagsClient<'a> {
    /// GET GetAvailablePackage
    /// Returns a list of available package tags. NOTE: This is a premium endpoint.
    /// Permissions Required:
    /// - Manage Tags
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_tags_available_package(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/tags/v2/package/available");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// GET GetAvailablePlant
    /// Returns a list of available plant tags. NOTE: This is a premium endpoint.
    /// Permissions Required:
    /// - Manage Tags
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_tags_available_plant(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/tags/v2/plant/available");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// GET GetStagedTags
    /// Returns a list of staged tags. NOTE: This is a premium endpoint.
    /// Permissions Required:
    /// - Manage Tags
    /// - RetailId.AllowPackageStaging Key Value enabled
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_staged_tags(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/tags/v2/staged");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
}

