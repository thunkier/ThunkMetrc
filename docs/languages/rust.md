# Rust SDK Guide

## Crates

```bash
cargo add thunkmetrc-client
cargo add thunkmetrc-wrapper
# Optional high-level helpers
cargo add thunkmetrc
```

## Recommended Setup (Wrapper)

```rust
use thunkmetrc_client::MetrcClient;
use thunkmetrc_wrapper::{MetrcWrapper, RateLimiterConfig};

let client = MetrcClient::new(
    "https://sandbox-api-or.metrc.com",
    "vendor_key",
    "user_key",
);

let limiter = RateLimiterConfig {
    enabled: true,
    ..Default::default()
};

let wrapper = MetrcWrapper::new_with_config(client, limiter);
```

Fetch facilities:

```rust
if let Some(facilities) = wrapper.facilities.get_facilities(None).await? {
    println!("Facilities: {}", facilities.len());
}
```

## Pagination Pattern

```rust
let license = "C12-0000001-LIC";
let mut all = Vec::new();
let mut page = 1;

loop {
    let response = wrapper
        .packages
        .get_active_packages(
            None, // last_modified_end
            None, // last_modified_start
            Some(license.to_string()),
            Some(page.to_string()),
            Some("20".to_string()),
            None,
        )
        .await?;

    let Some(response) = response else { break; };
    let data = response.data.unwrap_or_default();

    if data.is_empty() {
        break;
    }

    all.extend(data);
    page += 1;
}
```

## High-Level Helpers (`thunkmetrc`)

```rust
use chrono::{Duration, Utc};
use thunkmetrc::ThunkMetrc;

let thunk = ThunkMetrc::new(&wrapper);
let packages = thunk
    .active_packages_inventory_sync(
        "C12-0000001-LIC",
        Some(Utc::now() - Duration::hours(1)),
        5,
    )
    .await?;
```

## Build Locally

```bash
task build:rust
task test:rust
```
