# C# / .NET Integration Guide

## 📦 Installation

We assume you are using `.NET 8` or later.

### Wrapper (Recommended)
This package includes the type-safe wrapper, models, and rate limiter.
```bash
dotnet add package ThunkMetrc.Wrapper
```

### High-Level Utilities (Optional)
This package includes `ActivePackagesInventorySync` and other helpers.
```bash
dotnet add package ThunkMetrc
```

---

## 🚀 Getting Started

### 1. Initialize the Client
The `MetrcClient` is the low-level HTTP client. Usually you instantiate this once per user-session (since it holds the User API Key).

```csharp
using ThunkMetrc.Client;

var client = new MetrcClient(
    baseUrl: "https://sandbox-api-or.metrc.com", 
    vendorKey: "your_vendor_key", 
    userKey: "user_api_key"
);
```

### 2. Create the Wrapper
Wrap the client to get strong typing and rate limiting.

```csharp
using ThunkMetrc.Wrapper;

var wrapper = new MetrcWrapper(client);
```

### 3. Make Requests
All methods are `async` and return strongly-typed models.

```csharp
try 
{
    var deliveries = await wrapper.TransfersGetIncomingV2(
        licenseNumber: "C12-0000001-LIC",
        lastModifiedStart: DateTime.UtcNow.AddDays(-1).ToString("o")
    );
    
    foreach (var delivery in deliveries)
    {
        Console.WriteLine($"Delivery: {delivery.Id}");
    }
}
catch (MetrcException ex)
{
    Console.WriteLine($"Metrc Error: {ex.Message}");
}
```

---

## 🛡️ Rate Limiting

The Rate Limiter is **disabled by default**. To enable it:

```csharp
var config = new RateLimiterConfig 
{
    Enabled = true,
    MaxGetPerSecondPerFacility = 50,
    MaxConcurrentGetPerFacility = 10
};

var wrapper = new MetrcWrapper(client, config);
```

### How it works
*   **Automatic Retries**: If Metrc returns a `429 Too Many Requests`, the wrapper will automatically wait (with exponential backoff) and retry.
*   **Concurrency Control**: It uses a Semaphore to ensure you don't exceed `MaxConcurrentGetPerFacility` active requests.

---

## 🛠️ High-Level Features (`ThunkMetrc`)

If you installed the `ThunkMetrc` package, you can use advanced workflows.

### Inventory Sync
Synchronize all active packages for a license, handling pagination automatically (pages of 20).

```csharp
using ThunkMetrc;

var thunk = new ThunkMetrc.ThunkMetrc(wrapper);

var packages = await thunk.ActivePackagesInventorySyncAsync(
    licenseNumber: "C12-0000001-LIC",
    lastKnownSync: DateTime.UtcNow.AddHours(-24),
    bufferMinutes: 20
);
```
