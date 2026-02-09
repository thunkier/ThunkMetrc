# ThunkMetrc Rust Wrapper SDK

Type-safe, rate-limited, async Rust wrapper for the Metrc API.

## 📦 Installation

Add to your `Cargo.toml`:

```toml
[dependencies]
thunkmetrc-wrapper = "0.3.0"
thunkmetrc-client = "0.3.0"
tokio = { version = "1.0", features = ["full"] }
futures = "0.3"
```

## 🚀 Getting Started

### 1. Initialize
Use `MetrcFactory` to handle shared rate limiting effectively.

```rust
use thunkmetrc_wrapper::MetrcFactory;
use thunkmetrc_wrapper::MetrcWrapper;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Shared rate limiter (150 requests/sec integrator limit)
    let factory = MetrcFactory::new(150);

    let wrapper = factory.create(
        "https://sandbox-api-or.metrc.com",
        "vendor_key",
        "user_key"
    );

    Ok(())
}
```

### 2. Make Requests
All service methods are `async` and return `Result<Option<T>, Error>`.

```rust
match wrapper.facilities.get_facilities().await {
    Ok(Some(facilities)) => {
        for facility in facilities {
            println!("Facility: {}", facility.license.number.unwrap_or_default());
        }
    }
    Ok(None) => println!("No content"),
    Err(e) => eprintln!("Error: {:?}", e),
}
```

### 3. Pagination (Streams)
Use the `iterate_` methods to get a `Stream` of items, automatically handling page traversal.

```rust
use futures::StreamExt;

let mut stream = wrapper.packages.iterate_get_active_packages(
    Some("123-ABC-LICENSE".to_string()),
    None, // last_modified_start
    None  // last_modified_end
);

while let Some(result) = stream.next().await {
    match result {
        Ok(pkg) => println!("Package: {:?}", pkg.label),
        Err(e) => eprintln!("Error fetching page: {:?}", e),
    }
}
```

## 🛡️ Rate Limiting
The SDK uses `MetrcRateLimiter` to enforce:
- **Integrator Limits**: Default 150/sec.
- **Facility Limits**: Default 50/sec per facility.
- **Backoff**: Exponential backoff on 429/500 errors.
- **Retries**: Automatic retry logic properly handling `Retry-After`.

Configuration is handled via the `MetrcFactory`.
