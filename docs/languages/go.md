# Go SDK Guide

## Modules

```bash
go get github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client
go get github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper
# Optional high-level helpers
go get github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/thunkmetrc
```

Note: the Go modules use `sdks/thunkmetrc-go/...` paths.

## Recommended Setup (Wrapper + Factory)

```go
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client"
    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper"
    wrapperservices "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/services"
)

func main() {
    cfg := wrapperservices.DefaultConfig()
    cfg.Enabled = true

    factory := wrapper.NewFactory(
        cfg,
        client.WithTimeout(45*time.Second),
    )

    w := factory.GetWrapper(
        "https://sandbox-api-or.metrc.com",
        "vendor_key",
        "user_key",
    )

    facilities, err := w.Facilities.GetFacilities(context.Background())
    if err != nil {
        panic(err)
    }

    fmt.Printf("Facilities: %d\n", len(facilities))
}
```

## Pagination Pattern

Wrapper services expose paginated calls directly. Iterate pages explicitly:

```go
import (
    "context"
    "strconv"

    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper"
    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper/models"
)

func loadAllActivePackages(ctx context.Context, w *wrapper.MetrcWrapper, license string) ([]*models.PackagesPackage, error) {
    all := make([]*models.PackagesPackage, 0)

    for page := 1; ; page++ {
        resp, err := w.Packages.GetActivePackages(
            ctx,
            "", // lastModifiedEnd
            "", // lastModifiedStart
            license,
            strconv.Itoa(page),
            "20",
        )
        if err != nil {
            return nil, err
        }
        if len(resp.Data) == 0 {
            break
        }
        all = append(all, resp.Data...)
    }

    return all, nil
}
```

## High-Level Helpers (`thunkmetrc`)

The high-level package provides generated sync/iteration utilities in `extensions`.

```go
import (
    "context"

    "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/thunkmetrc/extensions"
)

packages, err := extensions.SyncGetActivePackages(
    context.Background(),
    w.Packages,
    nil, // lastKnownSync
    5,   // bufferMinutes
    "C12-0000001-LIC",
    "20",
)
```

## Build Locally

```bash
task build:go
task test:go
```
