use thunkmetrc::ThunkMetrc;
use thunkmetrc_wrapper::{MetrcWrapper, MetrcClient};
use dotenv::dotenv;
use std::env;

#[tokio::test]
async fn test_active_inventory_sync() {
    // Load .env from root
    let _ = dotenv(); 
    // Manual fallback if dotenv doesn't find it (sometimes flaky in workspaces)
    if env::var("METRC_VENDOR_API_KEY").is_err() {
        // Try to verify if we can find it relative to crate
        let path = std::path::Path::new("../../../.env");
        if path.exists() {
            dotenv::from_path(path).ok();
        }
    }

    let vendor_key = env::var("METRC_VENDOR_API_KEY").expect("METRC_VENDOR_API_KEY not set");
    let user_key = env::var("METRC_USER_API_KEY").expect("METRC_USER_API_KEY not set");
    let license_number = env::var("METRC_LICENSE_NUMBER").expect("METRC_LICENSE_NUMBER not set");
    let base_url = env::var("METRC_BASE_URL").unwrap_or("https://sandbox-api-or.metrc.com".to_string());

    let client = MetrcClient::new(&base_url, &vendor_key, &user_key);
    // Wrapper uses default config (no rate limit by default)
    let wrapper = MetrcWrapper::new(client);
    let thunk = ThunkMetrc::new(&wrapper);

    println!("Starting sync for {}", license_number);
    let start = std::time::Instant::now();

    let packages = thunk.active_packages_inventory_sync(
        &license_number,
        None,
        5 // buffer minutes
    ).await.expect("Sync failed");

    println!("Sync finished in {:?}. Found {} packages.", start.elapsed(), packages.len());
    
    // Assert result is reasonable (not crashing)
    // assert!(packages.len() >= 0); // Removed useless comparison
    assert!(true);
}
