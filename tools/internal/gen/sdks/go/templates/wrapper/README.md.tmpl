# ThunkMetrc Go Wrapper SDK

Type-safe, rate-limited, idiomatic Go wrapper for the Metrc API.

## 📦 Installation

```bash
go get github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper
go get github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client
```

## 🚀 Getting Started

### 1. Initialize
Use `MetrcFactory` to handle shared rate limiting effectively.

```go
package main

import (
    "context"
    "fmt"
    "time"
    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper"
)

func main() {
    // 1. Create Factory (Shared RateLimiter & HttpClient)
    // Default config uses 150 req/sec integration limit
    f := wrapper.NewFactory(
        wrapper.DefaultConfig(),
        client.WithTimeout(60 * time.Second),
    )

    // 2. Get Wrapper for specific license
    // Shares the factory's resources but enforces 50 req/sec facility limit
    w := f.GetWrapper("https://sandbox-api-or.metrc.com", "vendor_key", "user_key")
    
    // ...
}
```

### 2. Make Requests
All methods take `context.Context` and typed arguments.

```go
license := "C12-0000001-LIC"
facilities, err := w.Facilities.GetFacilities(context.Background())
if err != nil {
    panic(err)
}

for _, f := range facilities {
    fmt.Printf("Facility: %s\n", *f.License.Number)
}
```

### 3. Pagination (Iterators)
Use the `Iterate...` methods which return a next-item closure.

```go
// Create the iterator (returns a function)
iter := w.Packages.IterateGetActivePackages(context.Background(), license, nil, nil)

for {
    // Call the iterator to get the next item (lazily fetches pages)
    pkg, err := iter()
    if err != nil {
        panic(err)
    }
    if pkg == nil {
        break // End of iteration
    }
    
    fmt.Printf("Package: %s\n", *pkg.Label)
}
```

## 🛡️ Rate Limiting
The SDK uses `MetrcRateLimiter` to enforce:
- **Integrator Limits**: Default 150/sec.
- **Facility Limits**: Default 50/sec per facility.
- **Backoff**: Exponential backoff on 429/500 errors.
- **Retries**: Automatic retry logic.

Configuration is handled via `wrapper.Config` passed to `NewFactory`.

