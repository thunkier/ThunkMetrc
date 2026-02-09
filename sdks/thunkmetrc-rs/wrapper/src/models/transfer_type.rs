use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct TransferType {
    #[serde(rename = "BypassApproval", skip_serializing_if = "Option::is_none")]
    pub bypass_approval: Option<bool>,
    #[serde(rename = "ExternalIncomingCanRecordExternalIdentifier", skip_serializing_if = "Option::is_none")]
    pub external_incoming_can_record_external_identifier: Option<bool>,
    #[serde(rename = "ExternalIncomingExternalIdentifierRequired", skip_serializing_if = "Option::is_none")]
    pub external_incoming_external_identifier_required: Option<bool>,
    #[serde(rename = "ExternalOutgoingCanRecordExternalIdentifier", skip_serializing_if = "Option::is_none")]
    pub external_outgoing_can_record_external_identifier: Option<bool>,
    #[serde(rename = "ExternalOutgoingExternalIdentifierRequired", skip_serializing_if = "Option::is_none")]
    pub external_outgoing_external_identifier_required: Option<bool>,
    #[serde(rename = "ForExternalIncomingShipments", skip_serializing_if = "Option::is_none")]
    pub for_external_incoming_shipments: Option<bool>,
    #[serde(rename = "ForExternalOutgoingShipments", skip_serializing_if = "Option::is_none")]
    pub for_external_outgoing_shipments: Option<bool>,
    #[serde(rename = "ForLicensedShipments", skip_serializing_if = "Option::is_none")]
    pub for_licensed_shipments: Option<bool>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "RequiresDestinationGrossWeight", skip_serializing_if = "Option::is_none")]
    pub requires_destination_gross_weight: Option<bool>,
    #[serde(rename = "RequiresInvoiceNumber", skip_serializing_if = "Option::is_none")]
    pub requires_invoice_number: Option<bool>,
    #[serde(rename = "RequiresPDFDocument", skip_serializing_if = "Option::is_none")]
    pub requires_pdf_document: Option<bool>,
    #[serde(rename = "RequiresPackagesGrossWeight", skip_serializing_if = "Option::is_none")]
    pub requires_packages_gross_weight: Option<bool>,
    #[serde(rename = "TransactionType", skip_serializing_if = "Option::is_none")]
    pub transaction_type: Option<String>,
}
