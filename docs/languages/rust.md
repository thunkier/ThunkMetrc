# Rust Integration Guide

## 📦 Installation

```bash
cargo add thunkmetrc-wrapper
cargo add thunkmetrc # High-level
```

---

## 🚀 Getting Started

### 1. Initialize

```rust
use thunkmetrc_client::MetrcClient;
use thunkmetrc_wrapper::MetrcWrapper;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let client = MetrcClient::new(
        "https://sandbox-api-or.metrc.com", 
        "vendor_key", 
        "user_key"
    );

    let wrapper = MetrcWrapper::new(client);

    // ...
    Ok(())
}
```

### 2. Make Requests
Use `Some(...)` for optional parameters or `None`.

```rust
let result = wrapper.transfers_get_incoming_v2(
    None, // last_modified_end
    None, // last_modified_start
    Some("C12-0000001-LIC".to_string()),
    None, // page
    None  // page_size
).await?;

println!("{:?}", result);
```

---

## 🛡️ Rate Limiting

Rate limiting config is passed via `RateLimiterConfig`.

```rust
use thunkmetrc_wrapper::{MetrcWrapper, RateLimiterConfig};

let config = RateLimiterConfig {
    enabled: true,
    max_get_per_second_per_facility: 50.0,
    ..Default::default()
};

let wrapper = MetrcWrapper::new_with_config(client, config);
```

---

## 🛠️ High-Level Features (`thunkmetrc`)

### Inventory Sync

```rust
use thunkmetrc::active_packages_inventory_sync;
use chrono::{Utc, Duration};

let packages = active_packages_inventory_sync(
    &wrapper,
    "C12-0000001-LIC",
    Utc::now() - Duration::hours(24),
    15
).await?;
```
