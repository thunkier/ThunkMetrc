use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct WebhooksUpdateWebhooksRequest {
    #[serde(rename = "error_response_json_template", skip_serializing_if = "Option::is_none")]
    pub error_response_json_template: Option<String>,
    #[serde(rename = "facility_license_numbers", skip_serializing_if = "Option::is_none")]
    pub facility_license_numbers: Option<String>,
    #[serde(rename = "object_type", skip_serializing_if = "Option::is_none")]
    pub object_type: Option<String>,
    #[serde(rename = "server_public_key_fingerprint", skip_serializing_if = "Option::is_none")]
    pub server_public_key_fingerprint: Option<String>,
    #[serde(rename = "status", skip_serializing_if = "Option::is_none")]
    pub status: Option<String>,
    #[serde(rename = "template", skip_serializing_if = "Option::is_none")]
    pub template: Option<String>,
    #[serde(rename = "tpi_api_key", skip_serializing_if = "Option::is_none")]
    pub tpi_api_key: Option<String>,
    #[serde(rename = "url", skip_serializing_if = "Option::is_none")]
    pub url: Option<String>,
    #[serde(rename = "user_api_key", skip_serializing_if = "Option::is_none")]
    pub user_api_key: Option<String>,
    #[serde(rename = "verb", skip_serializing_if = "Option::is_none")]
    pub verb: Option<String>,
}
