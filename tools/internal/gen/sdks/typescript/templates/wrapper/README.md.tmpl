# ThunkMetrc TypeScript Wrapper

Type-safe, rate-limited, async wrapper for the Metrc API.

## 📦 Installation

```bash
npm install @thunkier/thunkmetrc-wrapper
npm install @thunkier/thunkmetrc-client
```

## 🚀 Getting Started

### 1. Initialize
Use `MetrcFactory` to handle shared rate limiting effectively.

```typescript
import { MetrcFactory } from '@thunkier/thunkmetrc-wrapper';

// 1. Create Factory
// Shared RateLimiter (150 requests/sec integrator limit)
const factory = new MetrcFactory(150);

// 2. Create Wrapper for specific license
const metrc = factory.create(
    "https://sandbox-api-or.metrc.com",
    "vendor_key",
    "user_key"
);
```

### 2. Make Requests
All service methods are `async`.

```typescript
try {
    const facilities = await metrc.facilities.getFacilities();
    facilities.forEach(f => {
        console.log(`Facility: ${f.License.Number}`);
    });
} catch (e) {
    console.error(e);
}
```

### 3. Pagination (Async Iterator)
Use the `iterate...` methods to stream pages comfortably.

```typescript
const iterator = metrc.packages.iterateGetActivePackages(
    "123-ABC-LICENSE"
);

// Use for-await-of loop
for await (const pkg of iterator) {
    console.log(`Package: ${pkg.Label}`);
}
```

## 🛡️ Rate Limiting
The SDK uses `MetrcRateLimiter` to enforce:
- **Integrator Limits**: Default 150/sec.
- **Facility Limits**: Default 50/sec per facility.
- **Backoff**: Exponential backoff on 429/500 errors.
- **Retries**: Automatic retry logic.

Configuration is handled via `MetrcFactory`.
