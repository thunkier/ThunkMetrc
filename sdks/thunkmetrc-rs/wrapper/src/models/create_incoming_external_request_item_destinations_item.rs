use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateIncomingExternalRequestItemDestinationsItem {
    #[serde(rename = "EstimatedArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "GrossUnitOfWeightId", skip_serializing_if = "Option::is_none")]
    pub gross_unit_of_weight_id: Option<i64>,
    #[serde(rename = "GrossWeight", skip_serializing_if = "Option::is_none")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "InvoiceNumber", skip_serializing_if = "Option::is_none")]
    pub invoice_number: Option<String>,
    #[serde(rename = "Packages", skip_serializing_if = "Option::is_none")]
    pub packages: Option<Vec<serde_json::Value>>,
    #[serde(rename = "PlannedRoute", skip_serializing_if = "Option::is_none")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub recipient_license_number: Option<String>,
    #[serde(rename = "TransferTypeName", skip_serializing_if = "Option::is_none")]
    pub transfer_type_name: Option<String>,
    #[serde(rename = "Transporters", skip_serializing_if = "Option::is_none")]
    pub transporters: Option<Vec<serde_json::Value>>,
}
