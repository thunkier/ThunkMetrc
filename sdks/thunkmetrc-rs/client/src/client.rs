use reqwest::Client;
use serde_json::Value;
use std::error::Error;
use log::{debug, warn};
use serde::Deserialize;
use crate::services::*;

#[derive(Debug)]
pub struct ApiError {
    pub status: reqwest::StatusCode,
    pub headers: reqwest::header::HeaderMap,
    pub message: String,
}

impl std::fmt::Display for ApiError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "API Error: {} - {}", self.status, self.message)
    }
}

impl Error for ApiError {}

#[derive(Clone)]
pub struct MetrcClient {
    client: Client,
    base_url: String,
    vendor_key: String,
    user_key: String,
}

impl MetrcClient {
    pub fn new(base_url: &str, vendor_key: &str, user_key: &str) -> Self {
        MetrcClient {
            client: Client::new(),
            base_url: base_url.trim_end_matches('/').to_string(),
            vendor_key: vendor_key.to_string(),
            user_key: user_key.to_string(),
        }
    }
    pub fn additives_templates(&self) -> AdditivesTemplatesClient {
        AdditivesTemplatesClient { client: self }
    }
    pub fn caregivers_status(&self) -> CaregiversStatusClient {
        CaregiversStatusClient { client: self }
    }
    pub fn employees(&self) -> EmployeesClient {
        EmployeesClient { client: self }
    }
    pub fn facilities(&self) -> FacilitiesClient {
        FacilitiesClient { client: self }
    }
    pub fn harvests(&self) -> HarvestsClient {
        HarvestsClient { client: self }
    }
    pub fn items(&self) -> ItemsClient {
        ItemsClient { client: self }
    }
    pub fn lab_tests(&self) -> LabTestsClient {
        LabTestsClient { client: self }
    }
    pub fn locations(&self) -> LocationsClient {
        LocationsClient { client: self }
    }
    pub fn packages(&self) -> PackagesClient {
        PackagesClient { client: self }
    }
    pub fn patient_check_ins(&self) -> PatientCheckInsClient {
        PatientCheckInsClient { client: self }
    }
    pub fn patients(&self) -> PatientsClient {
        PatientsClient { client: self }
    }
    pub fn patients_status(&self) -> PatientsStatusClient {
        PatientsStatusClient { client: self }
    }
    pub fn plant_batches(&self) -> PlantBatchesClient {
        PlantBatchesClient { client: self }
    }
    pub fn plants(&self) -> PlantsClient {
        PlantsClient { client: self }
    }
    pub fn processing_job(&self) -> ProcessingJobClient {
        ProcessingJobClient { client: self }
    }
    pub fn retail_id(&self) -> RetailIdClient {
        RetailIdClient { client: self }
    }
    pub fn sales(&self) -> SalesClient {
        SalesClient { client: self }
    }
    pub fn sandbox(&self) -> SandboxClient {
        SandboxClient { client: self }
    }
    pub fn strains(&self) -> StrainsClient {
        StrainsClient { client: self }
    }
    pub fn sublocations(&self) -> SublocationsClient {
        SublocationsClient { client: self }
    }
    pub fn tags(&self) -> TagsClient {
        TagsClient { client: self }
    }
    pub fn transfers(&self) -> TransfersClient {
        TransfersClient { client: self }
    }
    pub fn transporters(&self) -> TransportersClient {
        TransportersClient { client: self }
    }
    pub fn units_of_measure(&self) -> UnitsOfMeasureClient {
        UnitsOfMeasureClient { client: self }
    }
    pub fn waste_methods(&self) -> WasteMethodsClient {
        WasteMethodsClient { client: self }
    }
    pub fn webhooks(&self) -> WebhooksClient {
        WebhooksClient { client: self }
    }

    pub(crate) async fn send(&self, method: reqwest::Method, path: &str, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let url = format!("{}{}", self.base_url, path);
        debug!("Sending Request: {} {}", method, url);
        
        let mut req = self.client.request(method, &url);
        req = req.basic_auth(&self.vendor_key, Some(&self.user_key));

        if let Some(b) = body {
            req = req.json(b);
        }

        let resp = req.send().await?;
        let status = resp.status();
        let headers = resp.headers().clone();

        if !status.is_success() {
            let code = status.as_u16();
            warn!("API Error Response: {}", code);

            let text = resp.text().await.unwrap_or_default();
            let mut message = format!("Unexpected code {}", status);
            #[derive(Deserialize)]
            struct ApiErrorBody {
                #[serde(rename = "Message")]
                message: Option<String>,
                #[serde(rename = "ValidationErrors")]
                _validation_errors: Option<Vec<String>>,
            }

            if let Ok(body) = serde_json::from_str::<ApiErrorBody>(&text) {
                if let Some(m) = body.message {
                    message = m;
                }
            }

            return Err(Box::new(ApiError {
                status,
                headers,
                message,
            }));
        }
        if status == reqwest::StatusCode::NO_CONTENT {
            return Ok(None);
        }
        let json: Value = resp.json().await?;
        Ok(Some(json))
    }
}
