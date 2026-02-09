use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(rename_all = "PascalCase")]
pub struct PaginatedResponse<T> {
    pub total: Option<i64>,
    pub total_pages: Option<i64>,
    pub page_size: Option<i64>,
    pub current: Option<i64>,
    pub next: Option<i64>,
    pub previous: Option<i64>,
    pub data: Option<Vec<T>>,
}
