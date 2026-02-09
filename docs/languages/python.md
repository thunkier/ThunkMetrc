# Python SDK Guide

## Packages

```bash
pip install thunkmetrc-client
pip install thunkmetrc-wrapper
# Optional high-level helpers
pip install thunkmetrc
```

Python 3.10+ is recommended.

## Recommended Setup (Wrapper)

```python
import asyncio
from thunkmetrc.client import MetrcClient
from thunkmetrc.wrapper import MetrcWrapper

async def main() -> None:
    client = MetrcClient(
        base_url="https://sandbox-api-or.metrc.com",
        vendor_key="vendor_key",
        user_key="user_key",
    )
    wrapper = MetrcWrapper(client)

    facilities = await wrapper.facilities.get_facilities()
    print(f"Facilities: {len(facilities)}")

asyncio.run(main())
```

## Rate Limiter Configuration

```python
from thunkmetrc.client import MetrcClient
from thunkmetrc.wrapper import MetrcWrapper
from thunkmetrc.wrapper.ratelimiter import RateLimiterConfig

client = MetrcClient("https://sandbox-api-or.metrc.com", "vendor_key", "user_key")
limiter = RateLimiterConfig(enabled=True, max_get_per_second_integrator=150)
wrapper = MetrcWrapper(client, limiter)
```

## Pagination Pattern

```python
async def load_all_active_packages(wrapper: MetrcWrapper, license_number: str) -> list:
    all_packages = []
    page = 1

    while True:
        response = await wrapper.packages.get_active_packages(
            license_number=license_number,
            page_number=str(page),
            page_size="20",
        )

        data = (response or {}).get("Data", [])
        if not data:
            break

        all_packages.extend(data)
        page += 1

    return all_packages
```

## High-Level Helpers (`thunkmetrc`)

```python
from datetime import datetime, timedelta, timezone
from thunkmetrc import SynchronizationTarget, active_packages_inventory_sync

targets = [SynchronizationTarget("C12-0000001-LIC", wrapper)]
results = await active_packages_inventory_sync(
    targets,
    last_known_sync=datetime.now(timezone.utc) - timedelta(hours=1),
    buffer_minutes=5,
)

packages = results.get("C12-0000001-LIC", [])
```

## Build Locally

```bash
task build:python
task test:python
```
