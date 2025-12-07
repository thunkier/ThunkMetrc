# ThunkMetrc: Universal Metrc API Client

**ThunkMetrc** makes integrating with the **Metrc Track & Trace** system effortless. We provide a unified, auto-generated suite of **API Clients** and **Type-Safe Wrappers** for every major programming language.

Stop manually typing HTTP requests or debugging JSON serialization. Our libraries give you:
*   **Two Layers of Control**:
    *   **Core Client**: A lightweight, generic HTTP client for raw API access.
    *   **Wrapper**: A fully type-safe layer with generated models and helper methods.
*   **Multi-Language Support**: Native packages for **C#, Go, TypeScript, Python, Java, Kotlin, and Rust**.
*   **Consistency**: A single API surface across all languages, generated from a unified Bruno specification.
*   **Rate Limiting**: Built-in, configurable rate limiter to gracefully handle Metrc's API limits (429s) and concurrent request throttling.

> **Recommendation**: use the **Wrapper** packages for most applications to get strong typing and model validation. Use the **Client** packages if you need raw access or custom serialization.

Whether you're building a POS system, an ERP, or a compliance tool, ThunkMetrc is the best way to get your code talking to Metrc.

## 🚀 Quick Start (Wrappers)

Install the wrapper for your language for the best developer experience:

*   **NuGet**: `dotnet add package ThunkMetrc.Wrapper`
*   **npm**: `npm install @thunkier/thunkmetrc-wrapper`
*   **PyPI**: `pip install thunkmetrc-wrapper`
*   **Go**: `go get github.com/thunkier/ThunkMetrc/sdks/go/wrapper`
*   **Cargo**: `cargo add thunkmetrc-wrapper`
*   **Java/Kotlin**: Add via Maven/Gradle (Artifact: `com.thunkmetrc:wrapper`)

### Authentication
All wrappers require an instance of the underlying client. Authentication is handled by the client.

### Rate Limiting
The wrappers include a built-in rate limiter (token bucket + semaphore) that handles **429 retries** and **concurrency limits**. It is **disabled by default**.

**Enable it in your code:**
*   **C#**: `var wrapper = new MetrcWrapper(client); // Configure via MetrcRateLimiter`
*   **TypeScript**: `wrapper.rateLimiter.config.enabled = true;`
*   **Go**: `wrapper.RateLimiter.Config.Enabled = true`
*   **Python**: `wrapper = MetrcWrapper(client, RateLimiterConfig(enabled=True))`

### Usage Examples

#### C# / .NET 8 (Wrapper)
```csharp
using ThunkMetrc.Client;
using ThunkMetrc.Wrapper;

// 1. Initialize the Core Client
var client = new MetrcClient("https://sandbox-api-or.metrc.com", "vendor_key", "user_key");

// 2. Wrap it for Type Safety
var wrapper = new MetrcWrapper(client);

// 3. Call the API with strong types (Models are generated!)
var delivery = await wrapper.SalesGetDeliveryV1("12345");
```

#### TypeScript (Wrapper)
```typescript
import { MetrcClient, MetrcWrapper } from '@thunkier/thunkmetrc-wrapper';

const client = new MetrcClient("https://sandbox-api-or.metrc.com", "vendor_key", "user_key");
const wrapper = new MetrcWrapper(client);

// Fully verified types for request bodies and response handling
const delivery = await wrapper.salesGetDeliveryV1("12345");
```

#### Python (Wrapper)
```python
from thunkmetrc.wrapper import MetrcClient, MetrcWrapper

client = MetrcClient("https://sandbox-api-or.metrc.com", "vendor_key", "user_key")
wrapper = MetrcWrapper(client)

# Methods use generated data classes
delivery = wrapper.sales_get_delivery_v1("12345")
```

#### Go (Wrapper)
```go
import "github.com/thunkier/ThunkMetrc/sdks/go/wrapper"

// New convenience constructor
w := wrapper.NewClient("https://sandbox-api-or.metrc.com", "vendor_key", "user_key")

// Use structs for requests
resp, err := w.SalesGetDeliveryV1("12345")
```

---

## Roadmap & ToDo

We are actively improving the SDKs. The next major milestones are:

- [ ] **Response Models**: Generate strong types within wrapper packages.
- [ ] **Implement Response Models**: Ensure all API responses are strongly typed in the wrappers.
- [ ] **High-Level Utilities**: Add more helper functions for common Metrc workflows (e.g. Multi-State and Multi-Facility requests).

---

## 📦 High-Level Package (`thunkmetrc`)

We provide a **High-Level Package** that aggregates the Wrapper and Client, and includes powerful utility functions for common workflows.

**Install:**
*   **NuGet**: `dotnet add package ThunkMetrc`
*   **npm**: `npm install @thunkier/thunkmetrc`
*   **PyPI**: `pip install thunkmetrc`
*   **Go**: `go get github.com/thunkier/ThunkMetrc/sdks/go/thunkmetrc`
*   **Java/Kotlin**: `com.thunkmetrc:thunkmetrc`

### Features
*   **`ActivePackagesInventorySync`**: Efficiently syncs all active packages for a facility, handling time windows, pagination (limit 20), and date formatting automatically.

---

## 🛠️ Prerequisites

To develop or run the automation tools, you need:

*   **[Go](https://go.dev/)** (1.23+)
*   **[Task](https://taskfile.dev/)** (Build tool)

---

## 📚 Documentation

### Language Guides
*   **[C# / .NET](docs/languages/csharp.md)**
*   **[TypeScript / Node.js](docs/languages/typescript.md)**
*   **[Go](docs/languages/go.md)**
*   **[Python](docs/languages/python.md)**
*   **[Java](docs/languages/java.md)**
*   **[Kotlin](docs/languages/kotlin.md)**
*   **[Rust](docs/languages/rust.md)**

### Core Documentation
*   **[Configuration](docs/CONFIGURATION.md)**: `config.json`, Sanitization Rules, and Permission Overrides.
*   **[Development & Architecture](docs/DEVELOPMENT.md)**: Repo structure, Generators, Debugging, and Taskfile commands.

---

## ⚖️ License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.

For third-party dependencies and attributions, see the [NOTICE](NOTICE) file.

