use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct LabTestsLabTestBatcheLabTestTypesItem {
    #[serde(rename = "AlwaysPasses", skip_serializing_if = "Option::is_none")]
    pub always_passes: Option<bool>,
    #[serde(rename = "DependencyMode", skip_serializing_if = "Option::is_none")]
    pub dependency_mode: Option<i64>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "InformationalOnly", skip_serializing_if = "Option::is_none")]
    pub informational_only: Option<bool>,
    #[serde(rename = "LabTestResultExpirationDays", skip_serializing_if = "Option::is_none")]
    pub lab_test_result_expiration_days: Option<i64>,
    #[serde(rename = "LabTestResultMaximum", skip_serializing_if = "Option::is_none")]
    pub lab_test_result_maximum: Option<String>,
    #[serde(rename = "LabTestResultMinimum", skip_serializing_if = "Option::is_none")]
    pub lab_test_result_minimum: Option<String>,
    #[serde(rename = "LabTestResultMode", skip_serializing_if = "Option::is_none")]
    pub lab_test_result_mode: Option<i64>,
    #[serde(rename = "MaxAllowedFailureCount", skip_serializing_if = "Option::is_none")]
    pub max_allowed_failure_count: Option<i64>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "RequiresTestResult", skip_serializing_if = "Option::is_none")]
    pub requires_test_result: Option<bool>,
    #[serde(rename = "ResearchAndDevelopment", skip_serializing_if = "Option::is_none")]
    pub research_and_development: Option<bool>,
}
