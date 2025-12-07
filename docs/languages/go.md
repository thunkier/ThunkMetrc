# Go Integration Guide

## 📦 Installation

```bash
# Wrapper
go get github.com/thunkier/ThunkMetrc/sdks/go/wrapper

# High-Level Utilities
go get github.com/thunkier/ThunkMetrc/sdks/go/thunkmetrc
```

---

## 🚀 Getting Started

### 1. Initialize

```go
package main

import (
    "fmt"
    "github.com/thunkier/ThunkMetrc/sdks/go/client"
    "github.com/thunkier/ThunkMetrc/sdks/go/wrapper"
)

func main() {
    // 1. Create Core Client
    c := client.New(
        "https://sandbox-api-or.metrc.com", 
        "vendor_key", 
        "user_key",
    )

    // 2. Wrap it
    w := wrapper.New(c)

    // 3. Make Request
    // Note: Use pointers for optional string arguments
    license := "C12-0000001-LIC"
    resp, err := w.TransfersGetIncomingV2(nil, nil, &license, nil, nil)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Status: %d\n", resp.StatusCode)
}
```

---

## 🛡️ Rate Limiting

The rate limiter configuration is embedded in the wrapper struct.

```go
w := wrapper.New(c)

// Enable Rate Limiter
w.RateLimiter.Config.Enabled = true
w.RateLimiter.Config.MaxGetPerSecondPerFacility = 50.0
```

---

## 🛠️ High-Level Features (`thunkmetrc`)

### Inventory Sync

```go
import "github.com/thunkier/ThunkMetrc/sdks/go/thunkmetrc"

func main() {
    // ... setup wrapper w ...

    thunk := thunkmetrc.New(w)
    
    packages, err := thunk.ActivePackagesInventorySync(
        "C12-0000001-LIC",
        time.Now().Add(-24 * time.Hour), // lastKnownSync
        10,                              // bufferMinutes
    )
    
    if err != nil {
        panic(err)
    }
}
```
