use reqwest::Client;
use serde_json::Value;
use std::error::Error;

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

    async fn send(&self, method: reqwest::Method, path: &str, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let url = format!("{}{}", self.base_url, path);
        let mut req = self.client.request(method, &url);
        req = req.basic_auth(&self.vendor_key, Some(&self.user_key));

        if let Some(b) = body {
            req = req.json(b);
        }

        let resp = req.send().await?;
        let status = resp.status();
        if !status.is_success() {
            return Err(format!("API Error: {}", status).into());
        }
        if status == reqwest::StatusCode::NO_CONTENT {
            return Ok(None);
        }
        let json: Value = resp.json().await?;
        Ok(Some(json))
    }

    /// GET GetAll V1
    /// Permissions Required:
    ///   - Manage Employees
    ///
    pub async fn employees_get_all_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/employees/v1");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetAll V2
    /// Retrieves a list of employees for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Employees
    ///   - View Employees
    ///
    pub async fn employees_get_all_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/employees/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetPermissions V2
    /// Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Employees
    ///
    pub async fn employees_get_permissions_v2(&self, employee_license_number: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/employees/v2/permissions");
        let mut query_params = Vec::new();
        if let Some(p) = employee_license_number {
            query_params.push(format!("employeeLicenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Create V1
    /// NOTE: To include a photo with an item, first use POST /items/v1/photo to POST the photo, and then use the returned ID in the request body in this endpoint.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_create_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v1/create");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Create V2
    /// Creates one or more new products for the specified Facility. NOTE: To include a photo with an item, first use POST /items/v2/photo to POST the photo, and then use the returned Id in the request body in this endpoint.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_create_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateBrand V2
    /// Creates one or more new item brands for the specified Facility identified by the License Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_create_brand_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v2/brand");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateFile V2
    /// Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_create_file_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v2/file");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePhoto V1
    /// This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_create_photo_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v1/photo");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePhoto V2
    /// This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_create_photo_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v2/photo");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateUpdate V1
    /// Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_create_update_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v1/update");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V1
    /// Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_delete_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v1/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V2
    /// Archives the specified Product by Id for the given Facility License Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_delete_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE DeleteBrand V2
    /// Archives the specified Item Brand by Id for the given Facility License Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_delete_brand_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v2/brand/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v1/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V2
    /// Retrieves detailed information about a specific Item by Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_active_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v1/active");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V2
    /// Returns a list of active items for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v2/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetBrands V1
    /// Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_brands_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v1/brands");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetBrands V2
    /// Retrieves a list of active item brands for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_brands_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v2/brands");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetCategories V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn items_get_categories_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v1/categories");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetCategories V2
    /// Retrieves a list of item categories.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn items_get_categories_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v2/categories");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetFile V2
    /// Retrieves a file by its Id for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_file_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v2/file/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V1
    /// Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_inactive_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v1/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V2
    /// Retrieves a list of inactive items for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_inactive_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v2/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetPhoto V1
    /// Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_photo_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v1/photo/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetPhoto V2
    /// Retrieves an image by its Id for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_photo_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v2/photo/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT Update V2
    /// Updates one or more existing products for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_update_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateBrand V2
    /// Updates one or more existing item brands for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_update_brand_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/items/v2/brand");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Create V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_create_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/create");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Create V2
    /// Creates new packages for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_create_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAdjust V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_create_adjust_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/adjust");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAdjust V2
    /// Records a list of adjustments for packages at a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_create_adjust_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/adjust");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateChangeItem V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_create_change_item_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/change/item");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateChangeLocation V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_create_change_location_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/change/locations");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateFinish V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_create_finish_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/finish");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePlantings V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_create_plantings_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/create/plantings");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePlantings V2
    /// Creates new plantings from packages for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_create_plantings_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/plantings");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateRemediate V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_create_remediate_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/remediate");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateTesting V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_create_testing_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/create/testing");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateTesting V2
    /// Creates new packages for testing for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_create_testing_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/testing");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateUnfinish V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_create_unfinish_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/unfinish");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V2
    /// Discontinues a Package at a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_delete_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V2
    /// Retrieves a Package by its Id.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_active_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V2
    /// Retrieves a list of active packages for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetAdjustReasons V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn packages_get_adjust_reasons_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/adjust/reasons");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetAdjustReasons V2
    /// Retrieves a list of adjustment reasons for packages at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn packages_get_adjust_reasons_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/adjust/reasons");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetByLabel V1
    /// Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_by_label_v1(&self, label: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/{}", urlencoding::encode(label).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetByLabel V2
    /// Retrieves a Package by its label.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_by_label_v2(&self, label: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/{}", urlencoding::encode(label).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V1
    /// Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V2
    /// Retrieves a list of inactive packages for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetIntransit V2
    /// Retrieves a list of packages in transit for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_intransit_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/intransit");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetLabsamples V2
    /// Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_labsamples_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/labsamples");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetOnhold V1
    /// Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_onhold_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/onhold");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetOnhold V2
    /// Retrieves a list of packages on hold for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_onhold_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/onhold");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetSourceHarvest V2
    /// Retrieves the source harvests for a Package by its Id.
    /// 
    ///   Permissions Required:
    ///   - View Package Source Harvests
    ///
    pub async fn packages_get_source_harvest_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/{}/source/harvests", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTransferred V2
    /// Retrieves a list of transferred packages for a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_transferred_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/transferred");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTypes V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn packages_get_types_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/types");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTypes V2
    /// Retrieves a list of available Package types.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn packages_get_types_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/types");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateAdjust V2
    /// Set the final quantity for a Package.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_adjust_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/adjust");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateChangeNote V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///   - Manage Package Notes
    ///
    pub async fn packages_update_change_note_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v1/change/note");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDecontaminate V2
    /// Updates the Product decontaminate information for a list of packages at a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_decontaminate_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/decontaminate");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDonationFlag V2
    /// Flags one or more packages for donation at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_donation_flag_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/donation/flag");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDonationUnflag V2
    /// Removes the donation flag from one or more packages at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_donation_unflag_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/donation/unflag");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateExternalid V2
    /// Updates the external identifiers for one or more packages at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Package Inventory
    ///   - External Id Enabled
    ///
    pub async fn packages_update_externalid_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/externalid");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateFinish V2
    /// Updates a list of packages as finished for a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_finish_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/finish");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateItem V2
    /// Updates the associated Item for one or more packages at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_update_item_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/item");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateLabTestRequired V2
    /// Updates the list of required lab test batches for one or more packages at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_update_lab_test_required_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/labtests/required");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateLocation V2
    /// Updates the Location and Sublocation for one or more packages at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_update_location_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/location");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateNote V2
    /// Updates notes associated with one or more packages for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///   - Manage Package Notes
    ///
    pub async fn packages_update_note_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/note");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateRemediate V2
    /// Updates a list of Product remediations for packages at a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_remediate_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/remediate");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateTradesampleFlag V2
    /// Flags or unflags one or more packages at the specified Facility as trade samples.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_tradesample_flag_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/tradesample/flag");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateTradesampleUnflag V2
    /// Removes the trade sample flag from one or more packages at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_tradesample_unflag_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/tradesample/unflag");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateUnfinish V2
    /// Updates a list of packages as unfinished for a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_unfinish_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/unfinish");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateUsebydate V2
    /// Updates the use-by date for one or more packages at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_update_usebydate_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/packages/v2/usebydate");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Create V1
    /// Permissions Required:
    ///   - ManagePatientsCheckIns
    ///
    pub async fn patient_check_ins_create_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patient-checkins/v1");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Create V2
    /// Records patient check-ins for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - ManagePatientsCheckIns
    ///
    pub async fn patient_check_ins_create_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patient-checkins/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V1
    /// Permissions Required:
    ///   - ManagePatientsCheckIns
    ///
    pub async fn patient_check_ins_delete_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patient-checkins/v1/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V2
    /// Archives a Patient Check-In, identified by its Id, for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - ManagePatientsCheckIns
    ///
    pub async fn patient_check_ins_delete_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patient-checkins/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetAll V1
    /// Permissions Required:
    ///   - ManagePatientsCheckIns
    ///
    pub async fn patient_check_ins_get_all_v1(&self, checkin_date_end: Option<String>, checkin_date_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patient-checkins/v1");
        let mut query_params = Vec::new();
        if let Some(p) = checkin_date_end {
            query_params.push(format!("checkinDateEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = checkin_date_start {
            query_params.push(format!("checkinDateStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetAll V2
    /// Retrieves a list of patient check-ins for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - ManagePatientsCheckIns
    ///
    pub async fn patient_check_ins_get_all_v2(&self, checkin_date_end: Option<String>, checkin_date_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patient-checkins/v2");
        let mut query_params = Vec::new();
        if let Some(p) = checkin_date_end {
            query_params.push(format!("checkinDateEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = checkin_date_start {
            query_params.push(format!("checkinDateStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetLocations V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn patient_check_ins_get_locations_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patient-checkins/v1/locations");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetLocations V2
    /// Retrieves a list of Patient Check-In locations.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn patient_check_ins_get_locations_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patient-checkins/v2/locations");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT Update V1
    /// Permissions Required:
    ///   - ManagePatientsCheckIns
    ///
    pub async fn patient_check_ins_update_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patient-checkins/v1");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT Update V2
    /// Updates patient check-ins for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - ManagePatientsCheckIns
    ///
    pub async fn patient_check_ins_update_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patient-checkins/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAdditives V1
    /// Permissions Required:
    ///   - Manage Plants Additives
    ///
    pub async fn plant_batches_create_additives_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v1/additives");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAdditives V2
    /// Records Additive usage details for plant batches at a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Plants Additives
    ///
    pub async fn plant_batches_create_additives_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/additives");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAdditivesUsingtemplate V2
    /// Records Additive usage for plant batches at a Facility using predefined additive templates.
    /// 
    ///   Permissions Required:
    ///   - Manage Plants Additives
    ///
    pub async fn plant_batches_create_additives_usingtemplate_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/additives/usingtemplate");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAdjust V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///
    pub async fn plant_batches_create_adjust_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v1/adjust");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAdjust V2
    /// Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///
    pub async fn plant_batches_create_adjust_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/adjust");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateChangegrowthphase V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plant_batches_create_changegrowthphase_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v1/changegrowthphase");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateGrowthphase V2
    /// Updates the growth phase of plants at a specified Facility based on tracking information.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plant_batches_create_growthphase_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/growthphase");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePackage V2
    /// Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn plant_batches_create_package_v2(&self, is_from_mother_plant: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/packages");
        let mut query_params = Vec::new();
        if let Some(p) = is_from_mother_plant {
            query_params.push(format!("isFromMotherPlant={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePackageFrommotherplant V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn plant_batches_create_package_frommotherplant_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v1/create/packages/frommotherplant");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePackageFrommotherplant V2
    /// Creates packages from mother plants at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn plant_batches_create_package_frommotherplant_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/packages/frommotherplant");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePlantings V2
    /// Creates new plantings for a Facility by generating plant batches based on provided planting details.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///
    pub async fn plant_batches_create_plantings_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/plantings");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateSplit V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///
    pub async fn plant_batches_create_split_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v1/split");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateSplit V2
    /// Splits an existing Plant Batch into multiple groups at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///
    pub async fn plant_batches_create_split_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/split");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateWaste V1
    /// Permissions Required:
    ///   - Manage Plants Waste
    ///
    pub async fn plant_batches_create_waste_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v1/waste");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateWaste V2
    /// Records waste information for plant batches based on the submitted data for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Plants Waste
    ///
    pub async fn plant_batches_create_waste_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/waste");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Createpackages V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn plant_batches_createpackages_v1(&self, is_from_mother_plant: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v1/createpackages");
        let mut query_params = Vec::new();
        if let Some(p) = is_from_mother_plant {
            query_params.push(format!("isFromMotherPlant={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Createplantings V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///
    pub async fn plant_batches_createplantings_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v1/createplantings");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Destroy Immature Plants
    ///
    pub async fn plant_batches_delete_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v1");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V2
    /// Completes the destruction of plant batches based on the provided input data.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Destroy Immature Plants
    ///
    pub async fn plant_batches_delete_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///
    pub async fn plant_batches_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v1/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V2
    /// Retrieves a Plant Batch by Id.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///
    pub async fn plant_batches_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///
    pub async fn plant_batches_get_active_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v1/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V2
    /// Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///
    pub async fn plant_batches_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///
    pub async fn plant_batches_get_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v1/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V2
    /// Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///
    pub async fn plant_batches_get_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTypes V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn plant_batches_get_types_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v1/types");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTypes V2
    /// Retrieves a list of plant batch types.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn plant_batches_get_types_v2(&self, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/types");
        let mut query_params = Vec::new();
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetWaste V2
    /// Retrieves waste details associated with plant batches at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Plants Waste
    ///
    pub async fn plant_batches_get_waste_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/waste");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetWasteReasons V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn plant_batches_get_waste_reasons_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v1/waste/reasons");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetWasteReasons V2
    /// Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn plant_batches_get_waste_reasons_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/waste/reasons");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateLocation V2
    /// Moves one or more plant batches to new locations with in a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants
    ///
    pub async fn plant_batches_update_location_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/location");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateMoveplantbatches V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///
    pub async fn plant_batches_update_moveplantbatches_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v1/moveplantbatches");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateName V2
    /// Renames plant batches at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plant_batches_update_name_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/name");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateStrain V2
    /// Changes the strain of plant batches at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plant_batches_update_strain_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/strain");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateTag V2
    /// Replaces tags for plant batches at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plant_batches_update_tag_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plantbatches/v2/tag");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAdjust V1
    /// Permissions Required:
    ///   - ManageProcessingJobs
    ///
    pub async fn processing_jobs_create_adjust_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v1/adjust");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAdjust V2
    /// Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_create_adjust_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v2/adjust");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateJobtypes V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_create_jobtypes_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v1/jobtypes");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateJobtypes V2
    /// Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_create_jobtypes_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v2/jobtypes");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateStart V1
    /// Permissions Required:
    ///   - ManageProcessingJobs
    ///
    pub async fn processing_jobs_create_start_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v1/start");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateStart V2
    /// Initiates new processing jobs at a Facility, including job details and associated packages.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_create_start_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v2/start");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Createpackages V1
    /// Permissions Required:
    ///   - ManageProcessingJobs
    ///
    pub async fn processing_jobs_createpackages_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v1/createpackages");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Createpackages V2
    /// Creates packages from processing jobs at a Facility, including optional location and note assignments.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_createpackages_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v2/createpackages");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_delete_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v1/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V2
    /// Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_delete_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE DeleteJobtypes V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_delete_jobtypes_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v1/jobtypes/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE DeleteJobtypes V2
    /// Archives a Processing Job Type at a Facility, making it inactive for future use.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_delete_jobtypes_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v2/jobtypes/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v1/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V2
    /// Retrieves a ProcessingJob by Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_active_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v1/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V2
    /// Retrieves active processing jobs at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v2/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v1/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V2
    /// Retrieves inactive processing jobs at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v2/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetJobtypesActive V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_jobtypes_active_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v1/jobtypes/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetJobtypesActive V2
    /// Retrieves a list of all active processing job types defined within a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_jobtypes_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v2/jobtypes/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetJobtypesAttributes V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_jobtypes_attributes_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v1/jobtypes/attributes");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetJobtypesAttributes V2
    /// Retrieves all processing job attributes available for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_jobtypes_attributes_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v2/jobtypes/attributes");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetJobtypesCategories V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_jobtypes_categories_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v1/jobtypes/categories");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetJobtypesCategories V2
    /// Retrieves all processing job categories available for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_jobtypes_categories_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v2/jobtypes/categories");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetJobtypesInactive V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_jobtypes_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v1/jobtypes/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetJobtypesInactive V2
    /// Retrieves a list of all inactive processing job types defined within a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_jobtypes_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v2/jobtypes/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateFinish V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_update_finish_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v1/finish");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateFinish V2
    /// Completes processing jobs at a Facility by recording final notes and waste measurements.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_update_finish_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v2/finish");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateJobtypes V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_update_jobtypes_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v1/jobtypes");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateJobtypes V2
    /// Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_update_jobtypes_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v2/jobtypes");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateUnfinish V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_update_unfinish_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v1/unfinish");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateUnfinish V2
    /// Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_update_unfinish_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/processing/v2/unfinish");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Create V2
    /// Creates new sublocation records for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn sublocations_create_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sublocations/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V2
    /// Archives an existing Sublocation record for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn sublocations_delete_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sublocations/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V2
    /// Retrieves a Sublocation by its Id, with an optional license number.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn sublocations_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sublocations/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V2
    /// Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn sublocations_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sublocations/v2/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V2
    /// Retrieves a list of inactive sublocations for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn sublocations_get_inactive_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sublocations/v2/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT Update V2
    /// Updates existing sublocation records for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn sublocations_update_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sublocations/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateDelivery V1
    /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_create_delivery_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateDelivery V2
    /// Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
    ///
    pub async fn sales_create_delivery_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateDeliveryRetailer V1
    /// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/retailer");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateDeliveryRetailer V2
    /// Records retailer delivery data for a given License Number, including delivery destinations. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    ///   - Industry/Facility Type/Retailer Delivery
    ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/retailer");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateDeliveryRetailerDepart V1
    /// Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_depart_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/retailer/depart");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateDeliveryRetailerDepart V2
    /// Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    ///   - Industry/Facility Type/Retailer Delivery
    ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_depart_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/retailer/depart");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateDeliveryRetailerEnd V1
    /// Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_end_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/retailer/end");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateDeliveryRetailerEnd V2
    /// Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    ///   - Industry/Facility Type/Retailer Delivery
    ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_end_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/retailer/end");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateDeliveryRetailerRestock V1
    /// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_restock_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/retailer/restock");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateDeliveryRetailerRestock V2
    /// Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    ///   - Industry/Facility Type/Retailer Delivery
    ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_restock_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/retailer/restock");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateDeliveryRetailerSale V1
    /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_sale_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/retailer/sale");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateDeliveryRetailerSale V2
    /// Records sales deliveries originating from a retailer delivery for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    ///   - Industry/Facility Type/Retailer Delivery
    ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
    ///
    pub async fn sales_create_delivery_retailer_sale_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/retailer/sale");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateReceipt V1
    /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_create_receipt_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/receipts");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateReceipt V2
    /// Records a list of sales deliveries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales (Write)
    ///   - Industry/Facility Type/Consumer Sales or Industry/Facility Type/Patient Sales or Industry/Facility Type/External Patient Sales or Industry/Facility Type/Caregiver Sales
    ///   - Industry/Facility Type/Advanced Sales
    ///   - WebApi Sales Read Write State (All or WriteOnly)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
    ///
    pub async fn sales_create_receipt_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/receipts");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateTransactionByDate V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_create_transaction_by_date_v1(&self, date: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/transactions/{}", urlencoding::encode(date).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE DeleteDelivery V1
    /// Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_delete_delivery_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE DeleteDelivery V2
    /// Voids a sales delivery for a Facility using the provided License Number and delivery Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales Delivery
    ///
    pub async fn sales_delete_delivery_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE DeleteDeliveryRetailer V1
    /// Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_delete_delivery_retailer_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/retailer/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE DeleteDeliveryRetailer V2
    /// Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    ///   - Industry/Facility Type/Retailer Delivery
    ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_delete_delivery_retailer_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/retailer/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE DeleteReceipt V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_delete_receipt_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/receipts/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE DeleteReceipt V2
    /// Archives a sales receipt for a Facility using the provided License Number and receipt Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales
    ///
    pub async fn sales_delete_receipt_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/receipts/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetCounties V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn sales_get_counties_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/counties");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetCounties V2
    /// Returns a list of counties available for sales deliveries.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn sales_get_counties_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/counties");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetCustomertypes V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn sales_get_customertypes_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/customertypes");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetCustomertypes V2
    /// Returns a list of customer types.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn sales_get_customertypes_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/customertypes");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveriesActive V1
    /// Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_get_deliveries_active_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, sales_date_end: Option<String>, sales_date_start: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = sales_date_end {
            query_params.push(format!("salesDateEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = sales_date_start {
            query_params.push(format!("salesDateStart={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveriesActive V2
    /// Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
    /// 
    ///   Permissions Required:
    ///   - View Sales Delivery
    ///   - Manage Sales Delivery
    ///
    pub async fn sales_get_deliveries_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, sales_date_end: Option<String>, sales_date_start: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if let Some(p) = sales_date_end {
            query_params.push(format!("salesDateEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = sales_date_start {
            query_params.push(format!("salesDateStart={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveriesInactive V1
    /// Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_get_deliveries_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, sales_date_end: Option<String>, sales_date_start: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = sales_date_end {
            query_params.push(format!("salesDateEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = sales_date_start {
            query_params.push(format!("salesDateStart={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveriesInactive V2
    /// Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
    /// 
    ///   Permissions Required:
    ///   - View Sales Delivery
    ///   - Manage Sales Delivery
    ///
    pub async fn sales_get_deliveries_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, sales_date_end: Option<String>, sales_date_start: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if let Some(p) = sales_date_end {
            query_params.push(format!("salesDateEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = sales_date_start {
            query_params.push(format!("salesDateStart={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveriesRetailerActive V1
    /// Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_get_deliveries_retailer_active_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/retailer/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveriesRetailerActive V2
    /// Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
    /// 
    ///   Permissions Required:
    ///   - View Retailer Delivery
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_get_deliveries_retailer_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/retailer/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveriesRetailerInactive V1
    /// Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_get_deliveries_retailer_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/retailer/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveriesRetailerInactive V2
    /// Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
    /// 
    ///   Permissions Required:
    ///   - View Retailer Delivery
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_get_deliveries_retailer_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/retailer/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveriesReturnreasons V1
    /// Permissions Required:
    ///   -
    ///
    pub async fn sales_get_deliveries_returnreasons_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/returnreasons");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveriesReturnreasons V2
    /// Returns a list of return reasons for sales deliveries based on the provided License Number.
    /// 
    ///   Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_get_deliveries_returnreasons_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/returnreasons");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDelivery V1
    /// Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_get_delivery_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDelivery V2
    /// Retrieves a sales delivery record by its Id, with an optional License Number.
    /// 
    ///   Permissions Required:
    ///   - View Sales Delivery
    ///   - Manage Sales Delivery
    ///
    pub async fn sales_get_delivery_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveryRetailer V1
    /// Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_get_delivery_retailer_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/retailer/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveryRetailer V2
    /// Retrieves a retailer delivery record by its ID, with an optional License Number.
    /// 
    ///   Permissions Required:
    ///   - View Retailer Delivery
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_get_delivery_retailer_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/retailer/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetPatientRegistrationsLocations V1
    /// Permissions Required:
    ///   -
    ///
    pub async fn sales_get_patient_registrations_locations_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/patientregistration/locations");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetPatientRegistrationsLocations V2
    /// Returns a list of valid Patient registration locations for sales.
    /// 
    ///   Permissions Required:
    ///   -
    ///
    pub async fn sales_get_patient_registrations_locations_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/patientregistration/locations");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetPaymenttypes V1
    /// Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_get_paymenttypes_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/paymenttypes");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetPaymenttypes V2
    /// Returns a list of available payment types for the specified License Number.
    /// 
    ///   Permissions Required:
    ///   - View Sales Delivery
    ///   - Manage Sales Delivery
    ///
    pub async fn sales_get_paymenttypes_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/paymenttypes");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetReceipt V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_get_receipt_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/receipts/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetReceipt V2
    /// Retrieves a sales receipt by its Id, with an optional License Number.
    /// 
    ///   Permissions Required:
    ///   - View Sales
    ///   - Manage Sales
    ///
    pub async fn sales_get_receipt_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/receipts/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetReceiptsActive V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_get_receipts_active_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, sales_date_end: Option<String>, sales_date_start: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/receipts/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = sales_date_end {
            query_params.push(format!("salesDateEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = sales_date_start {
            query_params.push(format!("salesDateStart={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetReceiptsActive V2
    /// Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
    /// 
    ///   Permissions Required:
    ///   - View Sales
    ///   - Manage Sales
    ///
    pub async fn sales_get_receipts_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, sales_date_end: Option<String>, sales_date_start: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/receipts/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if let Some(p) = sales_date_end {
            query_params.push(format!("salesDateEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = sales_date_start {
            query_params.push(format!("salesDateStart={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetReceiptsExternalByExternalNumber V2
    /// Retrieves a Sales Receipt by its external number, with an optional License Number.
    /// 
    ///   Permissions Required:
    ///   - View Sales
    ///   - Manage Sales
    ///
    pub async fn sales_get_receipts_external_by_external_number_v2(&self, external_number: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/receipts/external/{}", urlencoding::encode(external_number).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetReceiptsInactive V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_get_receipts_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, sales_date_end: Option<String>, sales_date_start: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/receipts/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = sales_date_end {
            query_params.push(format!("salesDateEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = sales_date_start {
            query_params.push(format!("salesDateStart={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetReceiptsInactive V2
    /// Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
    /// 
    ///   Permissions Required:
    ///   - View Sales
    ///   - Manage Sales
    ///
    pub async fn sales_get_receipts_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, sales_date_end: Option<String>, sales_date_start: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/receipts/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if let Some(p) = sales_date_end {
            query_params.push(format!("salesDateEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = sales_date_start {
            query_params.push(format!("salesDateStart={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTransactions V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_get_transactions_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/transactions");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTransactionsBySalesDateStartAndSalesDateEnd V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_get_transactions_by_sales_date_start_and_sales_date_end_v1(&self, sales_date_start: &str, sales_date_end: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/transactions/{}/{}", urlencoding::encode(sales_date_start).as_ref(), urlencoding::encode(sales_date_end).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDelivery V1
    /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_update_delivery_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDelivery V2
    /// Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales Delivery
    ///
    pub async fn sales_update_delivery_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDeliveryComplete V1
    /// Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_update_delivery_complete_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/complete");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDeliveryComplete V2
    /// Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales Delivery
    ///
    pub async fn sales_update_delivery_complete_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/complete");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDeliveryHub V1
    /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_update_delivery_hub_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/hub");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDeliveryHub V2
    /// Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales Delivery, Manage Sales Delivery Hub
    ///
    pub async fn sales_update_delivery_hub_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/hub");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDeliveryHubAccept V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_update_delivery_hub_accept_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/hub/accept");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDeliveryHubAccept V2
    /// Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales Delivery Hub
    ///
    pub async fn sales_update_delivery_hub_accept_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/hub/accept");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDeliveryHubDepart V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_update_delivery_hub_depart_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/hub/depart");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDeliveryHubDepart V2
    /// Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales Delivery Hub
    ///
    pub async fn sales_update_delivery_hub_depart_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/hub/depart");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDeliveryHubVerifyID V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_update_delivery_hub_verify_id_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/hub/verifyID");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDeliveryHubVerifyID V2
    /// Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales Delivery Hub
    ///
    pub async fn sales_update_delivery_hub_verify_id_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/hub/verifyID");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDeliveryRetailer V1
    /// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_update_delivery_retailer_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/deliveries/retailer");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDeliveryRetailer V2
    /// Updates retailer delivery records for a given License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    ///   - Industry/Facility Type/Retailer Delivery
    ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_update_delivery_retailer_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/deliveries/retailer");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateReceipt V1
    /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_update_receipt_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/receipts");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateReceipt V2
    /// Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales
    ///
    pub async fn sales_update_receipt_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/receipts");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateReceiptFinalize V2
    /// Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales
    ///
    pub async fn sales_update_receipt_finalize_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/receipts/finalize");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateReceiptUnfinalize V2
    /// Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales
    ///
    pub async fn sales_update_receipt_unfinalize_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v2/receipts/unfinalize");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateTransactionByDate V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_update_transaction_by_date_v1(&self, date: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sales/v1/transactions/{}", urlencoding::encode(date).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Create V1
    /// Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_create_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/strains/v1/create");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Create V2
    /// Creates new strain records for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_create_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/strains/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateUpdate V1
    /// Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_create_update_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/strains/v1/update");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V1
    /// Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_delete_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/strains/v1/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V2
    /// Archives an existing strain record for a Facility
    /// 
    ///   Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_delete_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/strains/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/strains/v1/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V2
    /// Retrieves a Strain record by its Id, with an optional license number.
    /// 
    ///   Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/strains/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_get_active_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/strains/v1/active");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V2
    /// Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
    /// 
    ///   Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/strains/v2/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V2
    /// Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
    /// 
    ///   Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_get_inactive_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/strains/v2/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT Update V2
    /// Updates existing strain records for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_update_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/strains/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetPackageAvailable V2
    /// Returns a list of available package tags. NOTE: This is a premium endpoint.
    /// 
    ///   Permissions Required:
    ///   - Manage Tags
    ///
    pub async fn tags_get_package_available_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/tags/v2/package/available");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetPlantAvailable V2
    /// Returns a list of available plant tags. NOTE: This is a premium endpoint.
    /// 
    ///   Permissions Required:
    ///   - Manage Tags
    ///
    pub async fn tags_get_plant_available_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/tags/v2/plant/available");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetStaged V2
    /// Returns a list of staged tags. NOTE: This is a premium endpoint.
    /// 
    ///   Permissions Required:
    ///   - Manage Tags
    ///   - RetailId.AllowPackageStaging Key Value enabled
    ///
    pub async fn tags_get_staged_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/tags/v2/staged");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Create V2
    /// Creates new additive templates for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Additives
    ///
    pub async fn additives_templates_create_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/additivestemplates/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V2
    /// Retrieves an Additive Template by its Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Additives
    ///
    pub async fn additives_templates_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/additivestemplates/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V2
    /// Retrieves a list of active additive templates for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Additives
    ///
    pub async fn additives_templates_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/additivestemplates/v2/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V2
    /// Retrieves a list of inactive additive templates for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Additives
    ///
    pub async fn additives_templates_get_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/additivestemplates/v2/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT Update V2
    /// Updates existing additive templates for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Additives
    ///
    pub async fn additives_templates_update_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/additivestemplates/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateFinish V1
    /// Permissions Required:
    ///   - View Harvests
    ///   - Finish/Discontinue Harvests
    ///
    pub async fn harvests_create_finish_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v1/finish");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePackage V1
    /// Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn harvests_create_package_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v1/create/packages");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePackage V2
    /// Creates packages from harvested products for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn harvests_create_package_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v2/packages");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePackageTesting V1
    /// Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn harvests_create_package_testing_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v1/create/packages/testing");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePackageTesting V2
    /// Creates packages for testing from harvested products for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn harvests_create_package_testing_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v2/packages/testing");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateRemoveWaste V1
    /// Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///
    pub async fn harvests_create_remove_waste_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v1/removewaste");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateUnfinish V1
    /// Permissions Required:
    ///   - View Harvests
    ///   - Finish/Discontinue Harvests
    ///
    pub async fn harvests_create_unfinish_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v1/unfinish");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateWaste V2
    /// Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///
    pub async fn harvests_create_waste_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v2/waste");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE DeleteWaste V2
    /// Discontinues a specific harvest waste record by Id for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Discontinue Harvest Waste
    ///
    pub async fn harvests_delete_waste_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v2/waste/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v1/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V2
    /// Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_active_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v1/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V2
    /// Retrieves a list of active harvests for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v2/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V1
    /// Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v1/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V2
    /// Retrieves a list of inactive harvests for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v2/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetOnhold V1
    /// Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_onhold_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v1/onhold");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetOnhold V2
    /// Retrieves a list of harvests on hold for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_onhold_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v2/onhold");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetWaste V2
    /// Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_waste_v2(&self, harvest_id: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v2/waste");
        let mut query_params = Vec::new();
        if let Some(p) = harvest_id {
            query_params.push(format!("harvestId={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetWasteTypes V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn harvests_get_waste_types_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v1/waste/types");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetWasteTypes V2
    /// Retrieves a list of Waste types for harvests.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn harvests_get_waste_types_v2(&self, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v2/waste/types");
        let mut query_params = Vec::new();
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateFinish V2
    /// Marks one or more harvests as finished for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Finish/Discontinue Harvests
    ///
    pub async fn harvests_update_finish_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v2/finish");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateLocation V2
    /// Updates the Location of Harvest for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///
    pub async fn harvests_update_location_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v2/location");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateMove V1
    /// Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///
    pub async fn harvests_update_move_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v1/move");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateRename V1
    /// Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///
    pub async fn harvests_update_rename_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v1/rename");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateRename V2
    /// Renames one or more harvests for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///
    pub async fn harvests_update_rename_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v2/rename");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateRestoreHarvestedPlants V2
    /// Restores previously harvested plants to their original state for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Finish/Discontinue Harvests
    ///
    pub async fn harvests_update_restore_harvested_plants_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v2/restore/harvestedplants");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateUnfinish V2
    /// Reopens one or more previously finished harvests for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Finish/Discontinue Harvests
    ///
    pub async fn harvests_update_unfinish_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/harvests/v2/unfinish");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateRecord V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_create_record_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/labtests/v1/record");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateRecord V2
    /// Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_create_record_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/labtests/v2/record");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetBatches V2
    /// Retrieves a list of Lab Test batches.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn lab_tests_get_batches_v2(&self, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/labtests/v2/batches");
        let mut query_params = Vec::new();
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetLabtestdocument V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_get_labtestdocument_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/labtests/v1/labtestdocument/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetLabtestdocument V2
    /// Retrieves a specific Lab Test result document by its Id for a given Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_get_labtestdocument_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/labtests/v2/labtestdocument/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetResults V1
    /// Permissions Required:
    ///   - View Packages
    ///
    pub async fn lab_tests_get_results_v1(&self, license_number: Option<String>, package_id: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/labtests/v1/results");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = package_id {
            query_params.push(format!("packageId={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetResults V2
    /// Retrieves Lab Test results for a specified Package.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_get_results_v2(&self, license_number: Option<String>, package_id: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/labtests/v2/results");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = package_id {
            query_params.push(format!("packageId={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetStates V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn lab_tests_get_states_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/labtests/v1/states");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetStates V2
    /// Returns a list of all lab testing states.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn lab_tests_get_states_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/labtests/v2/states");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTypes V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn lab_tests_get_types_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/labtests/v1/types");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTypes V2
    /// Returns a list of Lab Test types.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn lab_tests_get_types_v2(&self, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/labtests/v2/types");
        let mut query_params = Vec::new();
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateLabtestdocument V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_update_labtestdocument_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/labtests/v1/labtestdocument");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateLabtestdocument V2
    /// Updates one or more documents for previously submitted lab tests.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_update_labtestdocument_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/labtests/v2/labtestdocument");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateResultRelease V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_update_result_release_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/labtests/v1/results/release");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateResultRelease V2
    /// Releases Lab Test results for one or more packages.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_update_result_release_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/labtests/v2/results/release");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Create V1
    /// Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_create_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/locations/v1/create");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Create V2
    /// Creates new locations for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_create_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/locations/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateUpdate V1
    /// Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_create_update_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/locations/v1/update");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V1
    /// Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_delete_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/locations/v1/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V2
    /// Archives a specified Location, identified by its Id, for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_delete_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/locations/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/locations/v1/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V2
    /// Retrieves a Location by its Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/locations/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_get_active_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/locations/v1/active");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V2
    /// Retrieves a list of active locations for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/locations/v2/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V2
    /// Retrieves a list of inactive locations for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_get_inactive_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/locations/v2/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTypes V1
    /// Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_get_types_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/locations/v1/types");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTypes V2
    /// Retrieves a list of active location types for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_get_types_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/locations/v2/types");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT Update V2
    /// Updates existing locations for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_update_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/locations/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetStatusesByPatientLicenseNumber V1
    /// Data returned by this endpoint is cached for up to one minute.
    /// 
    ///   Permissions Required:
    ///   - Lookup Patients
    ///
    pub async fn patients_status_get_statuses_by_patient_license_number_v1(&self, patient_license_number: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patients/v1/statuses/{}", urlencoding::encode(patient_license_number).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetStatusesByPatientLicenseNumber V2
    /// Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
    /// 
    ///   Permissions Required:
    ///   - Lookup Patients
    ///
    pub async fn patients_status_get_statuses_by_patient_license_number_v2(&self, patient_license_number: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patients/v2/statuses/{}", urlencoding::encode(patient_license_number).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAdditives V1
    /// Permissions Required:
    ///   - Manage Plants Additives
    ///
    pub async fn plants_create_additives_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/additives");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAdditives V2
    /// Records additive usage details applied to specific plants at a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Plants Additives
    ///
    pub async fn plants_create_additives_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/additives");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAdditivesBylocation V1
    /// Permissions Required:
    ///   - Manage Plants
    ///   - Manage Plants Additives
    ///
    pub async fn plants_create_additives_bylocation_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/additives/bylocation");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAdditivesBylocation V2
    /// Records additive usage for plants based on their location within a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Plants
    ///   - Manage Plants Additives
    ///
    pub async fn plants_create_additives_bylocation_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/additives/bylocation");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAdditivesBylocationUsingtemplate V2
    /// Records additive usage for plants by location using a predefined additive template at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Plants Additives
    ///
    pub async fn plants_create_additives_bylocation_usingtemplate_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/additives/bylocation/usingtemplate");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAdditivesUsingtemplate V2
    /// Records additive usage for plants using predefined additive templates at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Plants Additives
    ///
    pub async fn plants_create_additives_usingtemplate_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/additives/usingtemplate");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateChangegrowthphases V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_create_changegrowthphases_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/changegrowthphases");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateHarvestplants V1
    /// NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manicure/Harvest Veg/Flower Plants
    ///
    pub async fn plants_create_harvestplants_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/harvestplants");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateManicure V2
    /// Creates harvest product records from plant batches at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manicure/Harvest Veg/Flower Plants
    ///
    pub async fn plants_create_manicure_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/manicure");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateManicureplants V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manicure/Harvest Veg/Flower Plants
    ///
    pub async fn plants_create_manicureplants_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/manicureplants");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateMoveplants V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_create_moveplants_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/moveplants");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePlantbatchPackage V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn plants_create_plantbatch_package_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/create/plantbatch/packages");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePlantbatchPackage V2
    /// Creates packages from plant batches at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn plants_create_plantbatch_package_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/plantbatch/packages");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePlantings V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_create_plantings_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/create/plantings");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePlantings V2
    /// Creates new plant batches at a specified Facility from existing plant data.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_create_plantings_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/plantings");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateWaste V1
    /// Permissions Required:
    ///   - Manage Plants Waste
    ///
    pub async fn plants_create_waste_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/waste");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateWaste V2
    /// Records waste events for plants at a Facility, including method, reason, and location details.
    /// 
    ///   Permissions Required:
    ///   - Manage Plants Waste
    ///
    pub async fn plants_create_waste_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/waste");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Destroy Veg/Flower Plants
    ///
    pub async fn plants_delete_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V2
    /// Removes plants from a Facility’s inventory while recording the reason for their disposal.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Destroy Veg/Flower Plants
    ///
    pub async fn plants_delete_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V2
    /// Retrieves a Plant by Id.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetAdditives V1
    /// Permissions Required:
    ///   - View/Manage Plants Additives
    ///
    pub async fn plants_get_additives_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/additives");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetAdditives V2
    /// Retrieves additive records applied to plants at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View/Manage Plants Additives
    ///
    pub async fn plants_get_additives_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/additives");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetAdditivesTypes V1
    /// Permissions Required:
    ///   -
    ///
    pub async fn plants_get_additives_types_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/additives/types");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetAdditivesTypes V2
    /// Retrieves a list of all plant additive types defined within a Facility.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn plants_get_additives_types_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/additives/types");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetByLabel V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_by_label_v1(&self, label: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/{}", urlencoding::encode(label).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetByLabel V2
    /// Retrieves a Plant by label.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_by_label_v2(&self, label: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/{}", urlencoding::encode(label).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetFlowering V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_flowering_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/flowering");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetFlowering V2
    /// Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_flowering_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/flowering");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetGrowthPhases V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn plants_get_growth_phases_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/growthphases");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetGrowthPhases V2
    /// Retrieves the list of growth phases supported by a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn plants_get_growth_phases_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/growthphases");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V2
    /// Retrieves inactive plants at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetMother V2
    /// Retrieves mother-phase plants at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Mother Plants
    ///
    pub async fn plants_get_mother_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/mother");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetMotherInactive V2
    /// Retrieves inactive mother-phase plants at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Mother Plants
    ///
    pub async fn plants_get_mother_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/mother/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetMotherOnhold V2
    /// Retrieves mother-phase plants currently marked as on hold at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Mother Plants
    ///
    pub async fn plants_get_mother_onhold_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/mother/onhold");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetOnhold V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_onhold_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/onhold");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetOnhold V2
    /// Retrieves plants that are currently on hold at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_onhold_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/onhold");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetVegetative V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_vegetative_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/vegetative");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetVegetative V2
    /// Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_vegetative_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/vegetative");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetWaste V2
    /// Retrieves a list of recorded plant waste events for a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - View Plants Waste
    ///
    pub async fn plants_get_waste_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/waste");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetWasteMethodsAll V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn plants_get_waste_methods_all_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/waste/methods/all");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetWasteMethodsAll V2
    /// Retrieves a list of all available plant waste methods for use within a Facility.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn plants_get_waste_methods_all_v2(&self, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/waste/methods/all");
        let mut query_params = Vec::new();
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetWastePackage V2
    /// Retrieves a list of package records linked to the specified plantWasteId for a given facility.
    /// 
    ///   Permissions Required:
    ///   - View Plants Waste
    ///
    pub async fn plants_get_waste_package_v2(&self, id: &str, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/waste/{}/package", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetWastePlant V2
    /// Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
    /// 
    ///   Permissions Required:
    ///   - View Plants Waste
    ///
    pub async fn plants_get_waste_plant_v2(&self, id: &str, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/waste/{}/plant", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetWasteReasons V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn plants_get_waste_reasons_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v1/waste/reasons");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetWasteReasons V2
    /// Retriveves available reasons for recording mature plant waste at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn plants_get_waste_reasons_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/waste/reasons");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateAdjust V2
    /// Adjusts the recorded count of plants at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_update_adjust_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/adjust");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateGrowthphase V2
    /// Changes the growth phases of plants within a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_update_growthphase_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/growthphase");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateHarvest V2
    /// Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manicure/Harvest Veg/Flower Plants
    ///
    pub async fn plants_update_harvest_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/harvest");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateLocation V2
    /// Moves plant batches to new locations within a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_update_location_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/location");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateMerge V2
    /// Merges multiple plant groups into a single group within a Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manicure/Harvest Veg/Flower Plants
    ///
    pub async fn plants_update_merge_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/merge");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateSplit V2
    /// Splits an existing plant group into multiple groups within a Facility.
    /// 
    ///   Permissions Required:
    ///   - View Plant
    ///
    pub async fn plants_update_split_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/split");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateStrain V2
    /// Updates the strain information for plants within a Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_update_strain_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/strain");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateTag V2
    /// Replaces existing plant tags with new tags for plants within a Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_update_tag_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/plants/v2/tag");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAssociate V2
    /// Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly)
    ///   - Industry/View Packages
    ///   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
    ///
    pub async fn retail_id_create_associate_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/retailid/v2/associate");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateGenerate V2
    /// Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly)
    ///   - Industry/View Packages
    ///   - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
    ///
    pub async fn retail_id_create_generate_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/retailid/v2/generate");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateMerge V2
    /// Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly)
    ///   - Key Value Settings/Retail ID Merge Packages Enabled
    ///
    pub async fn retail_id_create_merge_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/retailid/v2/merge");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreatePackageInfo V2
    /// Retrieves Package information for given list of Package labels.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly)
    ///   - Industry/View Packages
    ///   - Admin/Employees/Packages Page/Product Labels(Manage)
    ///
    pub async fn retail_id_create_package_info_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/retailid/v2/packages/info");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetReceiveByLabel V2
    /// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Manage RetailId
    ///   - WebApi Retail ID Read Write State (All or ReadOnly)
    ///   - Industry/View Packages
    ///   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
    ///
    pub async fn retail_id_get_receive_by_label_v2(&self, label: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/retailid/v2/receive/{}", urlencoding::encode(label).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetReceiveQrByShortCode V2
    /// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Manage RetailId
    ///   - WebApi Retail ID Read Write State (All or ReadOnly)
    ///   - Industry/View Packages
    ///   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
    ///
    pub async fn retail_id_get_receive_qr_by_short_code_v2(&self, short_code: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/retailid/v2/receive/qr/{}", urlencoding::encode(short_code).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateIntegratorSetup V2
    /// This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn sandbox_create_integrator_setup_v2(&self, user_key: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/sandbox/v2/integrator/setup");
        let mut query_params = Vec::new();
        if let Some(p) = user_key {
            query_params.push(format!("userKey={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetAll V1
    /// This endpoint provides a list of facilities for which the authenticated user has access.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn facilities_get_all_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/facilities/v1");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetAll V2
    /// This endpoint provides a list of facilities for which the authenticated user has access.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn facilities_get_all_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/facilities/v2");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST Create V2
    /// Adds new patients to a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_create_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patients/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateAdd V1
    /// Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_create_add_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patients/v1/add");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateUpdate V1
    /// Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_create_update_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patients/v1/update");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V1
    /// Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_delete_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patients/v1/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE Delete V2
    /// Removes a Patient, identified by an Id, from a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_delete_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patients/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patients/v1/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET Get V2
    /// Retrieves a Patient by Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patients/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_get_active_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patients/v1/active");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V2
    /// Retrieves a list of active patients for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_get_active_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patients/v2/active");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT Update V2
    /// Updates Patient information for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_update_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/patients/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateExternalIncoming V1
    /// Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_create_external_incoming_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/external/incoming");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateExternalIncoming V2
    /// Creates external incoming shipment plans for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///
    pub async fn transfers_create_external_incoming_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/external/incoming");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateTemplates V1
    /// Permissions Required:
    ///   - Transfer Templates
    ///
    pub async fn transfers_create_templates_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/templates");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateTemplatesOutgoing V2
    /// Creates new transfer templates for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfer Templates
    ///
    pub async fn transfers_create_templates_outgoing_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/templates/outgoing");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE DeleteExternalIncoming V1
    /// Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_delete_external_incoming_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/external/incoming/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE DeleteExternalIncoming V2
    /// Voids an external incoming shipment plan for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///
    pub async fn transfers_delete_external_incoming_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/external/incoming/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE DeleteTemplates V1
    /// Permissions Required:
    ///   - Transfer Templates
    ///
    pub async fn transfers_delete_templates_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/templates/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE DeleteTemplatesOutgoing V2
    /// Archives a transfer template for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfer Templates
    ///
    pub async fn transfers_delete_templates_outgoing_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/templates/outgoing/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveriesPackagesStates V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn transfers_get_deliveries_packages_states_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/deliveries/packages/states");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveriesPackagesStates V2
    /// Returns a list of available shipment Package states.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn transfers_get_deliveries_packages_states_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/deliveries/packages/states");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDelivery V1
    /// Please note: that the {id} parameter above represents a Shipment Plan ID.
    /// 
    ///   Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_delivery_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/{}/deliveries", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDelivery V2
    /// Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_delivery_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/{}/deliveries", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveryPackage V1
    /// Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_delivery_package_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/deliveries/{}/packages", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveryPackage V2
    /// Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_delivery_package_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/deliveries/{}/packages", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveryPackageRequiredlabtestbatches V1
    /// Please note: The {id} parameter above represents a Transfer Delivery Package ID, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_delivery_package_requiredlabtestbatches_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/deliveries/package/{}/requiredlabtestbatches", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveryPackageRequiredlabtestbatches V2
    /// Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_delivery_package_requiredlabtestbatches_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/deliveries/package/{}/requiredlabtestbatches", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveryPackageWholesale V1
    /// Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_delivery_package_wholesale_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/deliveries/{}/packages/wholesale", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveryPackageWholesale V2
    /// Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_delivery_package_wholesale_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/deliveries/{}/packages/wholesale", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveryTransporters V1
    /// Please note: that the {id} parameter above represents a Shipment Delivery ID.
    /// 
    ///   Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_delivery_transporters_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/deliveries/{}/transporters", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveryTransporters V2
    /// Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_delivery_transporters_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/deliveries/{}/transporters", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveryTransportersDetails V1
    /// Please note: The {id} parameter above represents a Shipment Delivery ID.
    /// 
    ///   Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_delivery_transporters_details_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/deliveries/{}/transporters/details", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDeliveryTransportersDetails V2
    /// Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_delivery_transporters_details_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/deliveries/{}/transporters/details", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetHub V2
    /// Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_hub_v2(&self, estimated_arrival_end: Option<String>, estimated_arrival_start: Option<String>, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/hub");
        let mut query_params = Vec::new();
        if let Some(p) = estimated_arrival_end {
            query_params.push(format!("estimatedArrivalEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = estimated_arrival_start {
            query_params.push(format!("estimatedArrivalStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetIncoming V1
    /// Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_incoming_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/incoming");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetIncoming V2
    /// Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_incoming_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/incoming");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetOutgoing V1
    /// Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_outgoing_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/outgoing");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetOutgoing V2
    /// Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_outgoing_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/outgoing");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetRejected V1
    /// Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_rejected_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/rejected");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetRejected V2
    /// Retrieves a list of shipments with rejected packages for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_rejected_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/rejected");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTemplates V1
    /// Permissions Required:
    ///   - Transfer Templates
    ///
    pub async fn transfers_get_templates_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/templates");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTemplatesDelivery V1
    /// Permissions Required:
    ///   - Transfer Templates
    ///
    pub async fn transfers_get_templates_delivery_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/templates/{}/deliveries", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTemplatesDeliveryPackage V1
    /// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_templates_delivery_package_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/templates/deliveries/{}/packages", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTemplatesDeliveryTransporters V1
    /// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Transfer Templates
    ///
    pub async fn transfers_get_templates_delivery_transporters_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/templates/deliveries/{}/transporters", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTemplatesDeliveryTransportersDetails V1
    /// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Transfer Templates
    ///
    pub async fn transfers_get_templates_delivery_transporters_details_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/templates/deliveries/{}/transporters/details", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTemplatesOutgoing V2
    /// Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfer Templates
    ///   - View Transfer Templates
    ///
    pub async fn transfers_get_templates_outgoing_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/templates/outgoing");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTemplatesOutgoingDelivery V2
    /// Retrieves a list of deliveries associated with a specific transfer template.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfer Templates
    ///   - View Transfer Templates
    ///
    pub async fn transfers_get_templates_outgoing_delivery_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/templates/outgoing/{}/deliveries", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTemplatesOutgoingDeliveryPackage V2
    /// Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfer Templates
    ///   - View Transfer Templates
    ///
    pub async fn transfers_get_templates_outgoing_delivery_package_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/templates/outgoing/deliveries/{}/packages", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTemplatesOutgoingDeliveryTransporters V2
    /// Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfer Templates
    ///   - View Transfer Templates
    ///
    pub async fn transfers_get_templates_outgoing_delivery_transporters_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/templates/outgoing/deliveries/{}/transporters", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTemplatesOutgoingDeliveryTransportersDetails V2
    /// Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfer Templates
    ///   - View Transfer Templates
    ///
    pub async fn transfers_get_templates_outgoing_delivery_transporters_details_v2(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/templates/outgoing/deliveries/{}/transporters/details", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTypes V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn transfers_get_types_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/types");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetTypes V2
    /// Retrieves a list of available transfer types for a Facility based on its license number.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn transfers_get_types_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/types");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateExternalIncoming V1
    /// Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_update_external_incoming_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/external/incoming");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateExternalIncoming V2
    /// Updates external incoming shipment plans for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///
    pub async fn transfers_update_external_incoming_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/external/incoming");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateTemplates V1
    /// Permissions Required:
    ///   - Transfer Templates
    ///
    pub async fn transfers_update_templates_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v1/templates");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateTemplatesOutgoing V2
    /// Updates existing transfer templates for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfer Templates
    ///
    pub async fn transfers_update_templates_outgoing_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transfers/v2/templates/outgoing");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateDriver V2
    /// Creates new driver records for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transporters
    ///
    pub async fn transporters_create_driver_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transporters/v2/drivers");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// POST CreateVehicle V2
    /// Creates new vehicle records for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transporters
    ///
    pub async fn transporters_create_vehicle_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transporters/v2/vehicles");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE DeleteDriver V2
    /// Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Transporters
    ///
    pub async fn transporters_delete_driver_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transporters/v2/drivers/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// DELETE DeleteVehicle V2
    /// Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Transporters
    ///
    pub async fn transporters_delete_vehicle_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transporters/v2/vehicles/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDriver V2
    /// Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
    /// 
    ///   Permissions Required:
    ///   - Transporters
    ///
    pub async fn transporters_get_driver_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transporters/v2/drivers/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetDrivers V2
    /// Retrieves a list of drivers for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Transporters
    ///
    pub async fn transporters_get_drivers_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transporters/v2/drivers");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetVehicle V2
    /// Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
    /// 
    ///   Permissions Required:
    ///   - Transporters
    ///
    pub async fn transporters_get_vehicle_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transporters/v2/vehicles/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetVehicles V2
    /// Retrieves a list of vehicles for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Transporters
    ///
    pub async fn transporters_get_vehicles_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transporters/v2/vehicles");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateDriver V2
    /// Updates existing driver records for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transporters
    ///
    pub async fn transporters_update_driver_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transporters/v2/drivers");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// PUT UpdateVehicle V2
    /// Updates existing vehicle records for a facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transporters
    ///
    pub async fn transporters_update_vehicle_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/transporters/v2/vehicles");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn units_of_measure_get_active_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/unitsofmeasure/v1/active");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetActive V2
    /// Retrieves all active units of measure.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn units_of_measure_get_active_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/unitsofmeasure/v2/active");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetInactive V2
    /// Retrieves all inactive units of measure.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn units_of_measure_get_inactive_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/unitsofmeasure/v2/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetAll V2
    /// Retrieves all available waste methods.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn waste_methods_get_all_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/wastemethods/v2");
        let mut query_params = Vec::new();
        if let Some(p) = no {
            query_params.push(format!("No={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetByCaregiverLicenseNumber V1
    /// Data returned by this endpoint is cached for up to one minute.
    /// 
    ///   Permissions Required:
    ///   - Lookup Caregivers
    ///
    pub async fn caregivers_status_get_by_caregiver_license_number_v1(&self, caregiver_license_number: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/caregivers/v1/status/{}", urlencoding::encode(caregiver_license_number).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

    /// GET GetByCaregiverLicenseNumber V2
    /// Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
    /// 
    ///   Permissions Required:
    ///   - Lookup Caregivers
    ///
    pub async fn caregivers_status_get_by_caregiver_license_number_v2(&self, caregiver_license_number: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let mut path = format!("/caregivers/v2/status/{}", urlencoding::encode(caregiver_license_number).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }
        self.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }

}
