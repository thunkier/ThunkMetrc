# Configuration

`config.json` is the shared configuration file used by tooling and probe.

## Location

The loader searches these paths in order:

1. `config.json`
2. `../config.json`
3. `../../config.json`

## Structure

```json
{
  "Automation": {
    "SanitizeRules": [
      {
        "Pattern": "*items/v*/create",
        "Fields": ["BadField"]
      }
    ],
    "BlockedEndpoints": ["/integrator/setup"]
  },
  "PermissionOverrides": {
    "Docs Permission Name": "Actual Permission Name"
  },
  "Scraper": {
    "SingularMap": {
      "deliveries": "Delivery"
    },
    "CasingOverrides": {
      "unitsofmeasure": "UOM"
    },
    "States": ["ca", "or"],
    "Concurrency": 5,
    "RequestTimeoutSeconds": 30,
    "MaxRetries": 2,
    "RetryBackoffMs": 500,
    "UserAgent": "ThunkMetrc-SpecGen/1.0 (+https://github.com/thunkier/thunkmetrc)"
  }
}
```

## Automation

### SanitizeRules
Use `SanitizeRules` to remove request fields that consistently fail in specific endpoints.

- `Pattern`: glob-style URL matcher.
- `Fields`: JSON keys to recursively remove.

### BlockedEndpoints
List endpoint names or fragments the probe should skip.

## PermissionOverrides
Map documentation permission labels to actual API permission names when Metrc docs differ from runtime behavior.

## Scraper

`Scraper` drives `tools gen specs` behavior.

- `SingularMap`: naming singularization overrides.
- `CasingOverrides`: token casing overrides in generated operation names.
- `States`: optional list of states/regions to scrape.
- `Concurrency`: number of concurrent state workers.
- `RequestTimeoutSeconds`: HTTP timeout per request.
- `MaxRetries`: retry attempts for transient errors.
- `RetryBackoffMs`: exponential backoff base delay.
- `UserAgent`: explicit scraper User-Agent.

CLI flags for `tools gen specs` override these values when provided.

## Environment Variables

- `METRC_BASE_URL`
- `METRC_VENDOR_API_KEY`
- `METRC_USER_API_KEY`
- `METRC_FACILITY_LICENSE_NUMBER`
- `METRC_RESPONSES_DIR`
- `METRC_USE_MOCK_PERMISSIONS`

## Rate Limiting Defaults (Wrappers)

- `MaxGetPerSecondPerFacility`: `50.0`
- `MaxGetPerSecondIntegrator`: `150.0`
- `MaxConcurrentGetPerFacility`: `10`
- `MaxConcurrentGetIntegrator`: `30`
