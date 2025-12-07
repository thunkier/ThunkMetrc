# Python Integration Guide

## 📦 Installation

```bash
# Wrapper
pip install thunkmetrc-wrapper

# High-Level Utilities
pip install thunkmetrc
```

---

## 🚀 Getting Started

### 1. Initialize

```python
from thunkmetrc.client import MetrcClient
from thunkmetrc.wrapper import MetrcWrapper

# 1. Initialize Client
client = MetrcClient(
    base_url="https://sandbox-api-or.metrc.com",
    vendor_key="your_vendor_key",
    user_key="user_api_key"
)

# 2. Wrap it
wrapper = MetrcWrapper(client)
```

### 2. Make Requests
Methods use snake_case and keyword arguments.

```python
try:
    deliveries = wrapper.transfers_get_incoming_v2(
        license_number="C12-0000001-LIC",
        last_modified_start="2023-01-01T12:00:00Z"
    )
    print(deliveries)

except Exception as e:
    print(f"Error: {e}")
```

---

## 🛡️ Rate Limiting

Pass a configuration object during wrapper initialization.

```python
from thunkmetrc.wrapper import RateLimiterConfig

config = RateLimiterConfig(
    enabled=True,
    max_get_per_second_per_facility=50.0
)

wrapper = MetrcWrapper(client, config)
```

---

## 🛠️ High-Level Features (`thunkmetrc`)

### Inventory Sync

```python
from thunkmetrc.inventory import active_packages_inventory_sync
import datetime

# The high-level functions often take the wrapper as an argument
packages = active_packages_inventory_sync(
    wrapper=wrapper,
    license_number="C12-0000001-LIC",
    last_known_sync=datetime.datetime.now() - datetime.timedelta(days=1),
    buffer_minutes=15
)
```
