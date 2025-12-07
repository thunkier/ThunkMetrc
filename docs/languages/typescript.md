# TypeScript / Node.js Integration Guide

## 📦 Installation

Compatible with Node.js 18+ and modern browsers (bundlers required).

### Wrapper (Recommended)
```bash
npm install @thunkier/thunkmetrc-wrapper
```

### High-Level Utilities (Optional)
```bash
npm install @thunkier/thunkmetrc
```

---

## 🚀 Getting Started

### 1. Initialize
We export everything from the wrapper package for convenience.

```typescript
import { MetrcClient, MetrcWrapper } from '@thunkier/thunkmetrc-wrapper';

const client = new MetrcClient(
    "https://sandbox-api-or.metrc.com", 
    "your_vendor_key", 
    "user_api_key"
);

// Basic wrapper
const wrapper = new MetrcWrapper(client);
```

### 2. Make Requests
Parameters are named arguments (object destructuring) or standard arguments depending on the generated signature.

```typescript
try {
    const deliveries = await wrapper.transfersGetIncomingV2(
        "2023-01-01T00:00:00Z", // lastModifiedEnd
        "2023-01-02T00:00:00Z", // lastModifiedStart
        "C12-0000001-LIC"       // licenseNumber
    );
    
    console.log(deliveries);
} catch (error) {
    console.error("Metrc Error:", error);
}
```

---

## 🛡️ Rate Limiting

The Rate Limiter is **disabled by default**. To enable it, access the `rateLimiter` property on the wrapper.

```typescript
// Enable with defaults
wrapper.rateLimiter.config.enabled = true;

// Or customize
wrapper.rateLimiter.config = {
    enabled: true,
    maxGetPerSecondPerFacility: 50,
    maxConcurrentGetPerFacility: 10,
    // ...other options
};
```

---

## 🛠️ High-Level Features (`@thunkier/thunkmetrc`)

### Inventory Sync
Efficiently fetch all active packages.

```typescript
import { ThunkMetrc } from '@thunkier/thunkmetrc';

const thunk = new ThunkMetrc(wrapper);

const packages = await thunk.activePackagesInventorySync(
    "C12-0000001-LIC",
    new Date("2023-01-01"), // lastKnownSync
    15                      // bufferMinutes
);
```
