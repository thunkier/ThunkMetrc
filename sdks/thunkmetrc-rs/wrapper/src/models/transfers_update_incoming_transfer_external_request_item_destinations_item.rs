use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct TransfersUpdateIncomingTransferExternalRequestItemDestinationsItem {
    #[serde(rename = "estimated_arrival_date_time", skip_serializing_if = "Option::is_none")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "estimated_departure_date_time", skip_serializing_if = "Option::is_none")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "gross_unit_of_weight_id", skip_serializing_if = "Option::is_none")]
    pub gross_unit_of_weight_id: Option<i64>,
    #[serde(rename = "gross_weight", skip_serializing_if = "Option::is_none")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "invoice_number", skip_serializing_if = "Option::is_none")]
    pub invoice_number: Option<String>,
    #[serde(rename = "packages", skip_serializing_if = "Option::is_none")]
    pub packages: Option<Vec<TransfersUpdateIncomingTransferExternalRequestItemDestinationsItemPackagesItem>>,
    #[serde(rename = "planned_route", skip_serializing_if = "Option::is_none")]
    pub planned_route: Option<String>,
    #[serde(rename = "recipient_license_number", skip_serializing_if = "Option::is_none")]
    pub recipient_license_number: Option<String>,
    #[serde(rename = "transfer_destination_id", skip_serializing_if = "Option::is_none")]
    pub transfer_destination_id: Option<i64>,
    #[serde(rename = "transfer_type_name", skip_serializing_if = "Option::is_none")]
    pub transfer_type_name: Option<String>,
    #[serde(rename = "transporters", skip_serializing_if = "Option::is_none")]
    pub transporters: Option<Vec<TransfersUpdateIncomingTransferExternalRequestItemDestinationsItemTransportersItem>>,
}
