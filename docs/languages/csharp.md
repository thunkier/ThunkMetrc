# C# SDK Guide

## Packages

```bash
dotnet add package ThunkMetrc.Client
dotnet add package ThunkMetrc.Wrapper
# Optional high-level helpers
dotnet add package ThunkMetrc
```

Targets `net8.0`.

## Recommended Setup (Wrapper)

```csharp
using ThunkMetrc.Wrapper;

var limiterConfig = new RateLimiterConfig
{
    Enabled = true,
    MaxGetPerSecondIntegrator = 150,
    MaxConcurrentGetIntegrator = 30
};

using var factory = new MetrcFactory(limiterConfig);

var wrapper = factory.Create(
    "https://sandbox-api-or.metrc.com",
    "vendor_key",
    "user_key"
);

var facilities = await wrapper.Facilities.GetFacilities();
Console.WriteLine($"Facilities: {facilities.Count}");
```

## Raw Client (When You Need Untyped Responses)

```csharp
using ThunkMetrc.Client;

var client = new MetrcClient(
    "https://sandbox-api-or.metrc.com",
    "vendor_key",
    "user_key"
);

object rawFacilities = await client.GetFacilities();
```

## Pagination Pattern

The wrapper exposes paginated endpoints directly. Iterate pages explicitly:

```csharp
using ThunkMetrc.Wrapper.Models;

var allPackages = new List<PackagesPackage>();
var page = 1;

while (true)
{
    var response = await wrapper.Packages.GetActivePackages(
        licenseNumber: "C12-0000001-LIC",
        pageNumber: page.ToString(),
        pageSize: "20"
    );

    if (response?.Data == null || response.Data.Count == 0)
    {
        break;
    }

    allPackages.AddRange(response.Data);
    page++;
}
```

## High-Level Helpers (`ThunkMetrc`)

The `ThunkMetrc` package adds extension helpers (sync windows, parallel helpers, iterators).

```csharp
using ThunkMetrc.Extensions;

var packages = await wrapper.Packages.SyncGetActivePackages(
    lastKnownSync: DateTimeOffset.UtcNow.AddHours(-2),
    bufferMinutes: 5,
    licenseNumber: "C12-0000001-LIC"
);
```

## Build Locally

```bash
task build:csharp
task test:csharp
```
