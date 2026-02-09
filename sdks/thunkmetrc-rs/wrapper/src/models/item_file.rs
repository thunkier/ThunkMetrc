use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ItemFile {
    #[serde(rename = "ContentType", skip_serializing_if = "Option::is_none")]
    pub content_type: Option<String>,
    #[serde(rename = "FileContents", skip_serializing_if = "Option::is_none")]
    pub file_contents: Option<String>,
    #[serde(rename = "FileDownloadName", skip_serializing_if = "Option::is_none")]
    pub file_download_name: Option<String>,
}
