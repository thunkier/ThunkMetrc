# ThunkMetrc Python Wrapper

Type-safe, rate-limited, async wrapper for the Metrc API.

## 📦 Installation

```bash
pip install thunkmetrc-wrapper
pip install thunkmetrc-client
```

## 🚀 Getting Started

### 1. Initialize
Use `MetrcFactory` to handle shared rate limiting effectively.

```python
import asyncio
from thunkmetrc.wrapper import MetrcFactory

async def main():
    # 1. Create Factory
    # Shared RateLimiter (150 requests/sec integrator limit)
    factory = MetrcFactory(max_concurrent_requests=150)

    # 2. Create Wrapper for specific license
    metrc = factory.create(
        base_url="https://sandbox-api-or.metrc.com",
        vendor_key="vendor_key",
        user_key="user_key"
    )
    
    # ...
```

### 2. Make Requests
All service methods are `async`.

```python
    try:
        facilities = await metrc.facilities.get_facilities()
        for f in facilities:
            print(f"Facility: {f.License.Number}")
    except Exception as e:
        print(f"Error: {e}")
```

### 3. Pagination (Async Iterator)
Use the `iterate_` methods to stream pages comfortably.

```python
    async for pkg in metrc.packages.iterate_get_active_packages(
        license_number="123-ABC-LICENSE"
    ):
        print(f"Package: {pkg.Label}")

if __name__ == "__main__":
    asyncio.run(main())
```

## 🛡️ Rate Limiting
The SDK uses `MetrcRateLimiter` to enforce:
- **Integrator Limits**: Default 150/sec.
- **Facility Limits**: Default 50/sec per facility.
- **Backoff**: Exponential backoff on 429/500 errors.
- **Retries**: Automatic retry logic.

Configuration is handled via `MetrcFactory`.
