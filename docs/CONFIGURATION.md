# Configuration

All automation configuration is consolidated in `automation/config.json`.

## `automation/config.json` Structure

```json
{
  "Automation": {
    "SanitizeRules": [
      {
        "Pattern": "*items/v*/create", // Wildcard matching on URL
        "Fields": ["BadField"]         // Fields to remove from request
      }
    ],
    "BlockedEndpoints": ["/integrator/setup"]
  },
  "PermissionOverrides": {
    "Docs Permission Name": "Actual Permission Name"
  }
}
```

## SanitizeRules
Use `SanitizeRules` to remove specific fields from the request body if they cause issues.
*   **Pattern**: A glob pattern matching the request URL (e.g. `*`, `*v1*`).
*   **Fields**: A list of JSON keys to recursively remove from the body.

## PermissionOverrides
Sometimes, the Metrc documentation lists a permission that differs from what the implementation actually requires. Use this map to translate them.
*   **Key**: The permission name as seen in the HTML documentation.
*   **Value**: The actual permission string required by the API token.

## Environment Variables

The SDKs and Automation tools rely on the following environment variables (also supported via `.env` files):

*   **`METRC_BASE_URL`**: Base URL for the Metrc API (e.g., `https://sandbox-api-or.metrc.com`).
*   **`METRC_VENDOR_API_KEY`**: vendor API key.
*   **`METRC_USER_API_KEY`**: user API key.
*   **`METRC_LICENSE_NUMBER`**: Facility license number for tests.

## Rate Limiting Configuration

The SDK Wrappers include a configurable Rate Limiter.

### Default Limits
*   `MaxGetPerSecondPerFacility`: 50.0
*   `MaxGetPerSecondIntegrator`: 150.0
*   `MaxConcurrentGetPerFacility`: 10
*   `MaxConcurrentGetIntegrator`: 30
*   **Enabled**: `false` (default)
