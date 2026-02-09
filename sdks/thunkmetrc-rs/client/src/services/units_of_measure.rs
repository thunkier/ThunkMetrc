use crate::client::MetrcClient;
use serde_json::Value;
use std::error::Error;

pub struct UnitsOfMeasureClient<'a> {
    pub(crate) client: &'a MetrcClient,
}

impl<'a> UnitsOfMeasureClient<'a> {
    /// GET GetActiveUnitsOfMeasure
    /// Retrieves all active units of measure.
    /// Parameters:
    pub async fn get_active_units_of_measure(&self, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let path = format!("/unitsofmeasure/v2/active");

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// GET GetInactiveUnitsOfMeasure
    /// Retrieves all inactive units of measure.
    /// Parameters:
    pub async fn get_inactive_units_of_measure(&self, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let path = format!("/unitsofmeasure/v2/inactive");

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
}

