use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct RetailIdGenerate {
    #[serde(rename = "IssuanceId", skip_serializing_if = "Option::is_none")]
    pub issuance_id: Option<String>,
}
