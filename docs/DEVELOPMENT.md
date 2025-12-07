# Development & Architecture

## 📂 Repository Structure

*   **`sdks/`**: **Public Clients**. Auto-generated SDKs for external distribution.
    *   `csharp/`, `go/`, `java/`, `kotlin/`, `python/`, `rust/`, `typescript/`.
    *   Each SDK typically contains:
        *   `client/`: Low-level HTTP client (generated).
        *   `wrapper/`: Type-safe wrapper with optional rate limiting (generated).
        *   `thunkmetrc/`: High-level utility package (hand-written).
*   **`specs/`**: **Source of Truth**.
    *   `metrc-bruno/`: Bruno API collection defining all endpoints.
*   **`automation/`**: **Tools & Internal Runner**.
    *   `cmd/specgen/`: **Doc Scraper**. Scrapes Metrc documentation -> Bruno Specs.
    *   `cmd/sdkgen/`: **Public SDK Generator**. Builds `sdks/` from `specs/`.
    *   `cmd/modelgen/`: **Internal Model Generator**. Generates Go structs for the runner.
    *   `cmd/clientgen/`: **Internal Client Generator**. Builds the runner's internal client.
    *   `cmd/bump/`: **Version Manager**. Bumps `VERSION` file and orchestrates updates.
    *   `internal/runner/`: Internal automation logic for sandbox simulation.

---

## 🏗️ Architecture & Generators

The project uses a data-driven approach where the **Bruno Specs** (`specs/metrc-bruno`) act as the single source of truth.

1.  **Ingestion (`gen:spec`)**: `cmd/specgen` crawls Metrc documentation to update the Bruno specifications.
2.  **Internal Client (`gen:internal-client`)**: `cmd/clientgen` builds a type-safe Go client for the automation runner.
3.  **Public SDKs (`gen:sdks`)**: `cmd/sdkgen` builds standalone libraries for distribution.

### Public SDK Conventions
*   **Naming**: PascalCase (C#/Go), camelCase (TS/Java/Kotlin), snake_case (Python/Rust).
*   **Singularization**: `GetDeliveries` -> `GetDelivery` for `GET /id` endpoints.
*   **Authentication**: Unified `(baseUrl, vendorKey, userKey)` constructor.
*   **Query Parameters**: Optional query parameters are handled consistently, sorted alphabetically and URL-encoded.

---

## 🔧 Development & Debugging

We use **Taskfile** to manage common operations.

### Common Commands

| Command | Description |
| :--- | :--- |
| `task setup-sandbox` | Configure a new Integrator Sandbox User (requires `.env` with `METRC_VENDOR_API_KEY`). |
| `task run` | Run the main automation runner (dry-run by default). |
| `task gen:all` | Run all generation steps (Spec -> Internal -> SDKs). |
| `task test` | Run integration tests for all SDKs (including high-level packages). |
| `task publish:all` | Build and publish all SDK packages (`client`, `wrapper`, `thunkmetrc`). |

### Verification Commands
Use these commands to verify the generated code:
*   **C#**: `dotnet build`
*   **Go**: `go build`
*   **Rust**: `cargo check`
*   **TypeScript**: `npm run build`
*   **Python**: `python -m py_compile`
*   **Java/Kotlin**: `mvn clean install` (Requires JDK 21+).

### Debugging the Runner

To debug specific endpoint groups in the internal runner:

1.  **Filter Groups**:
    In `automation/internal/runner/workflow.go`, modify `ExecuteGroupWorkflow`:
    ```go
    isTarget := strings.Contains(groupName, "Additives")
    ```
2.  **Permission Overrides**:
    See [Configuration](CONFIGURATION.md).
