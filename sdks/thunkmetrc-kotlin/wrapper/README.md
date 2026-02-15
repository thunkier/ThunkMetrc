# ThunkMetrc Kotlin Wrapper SDK

Type-safe, rate-limited, coroutine-friendly wrapper for the Metrc API.

## 📦 Installation
Add dependencies to your `build.gradle.kts`:

```kotlin
implementation("io.github.thunkier:thunkmetrc-kotlin-wrapper:0.3.1")
implementation("io.github.thunkier:thunkmetrc-kotlin-client:0.3.1")
implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:1.7.3")
```

Or `pom.xml`:
```xml
<dependency>
    <groupId>io.github.thunkier</groupId>
    <artifactId>thunkmetrc-kotlin-wrapper</artifactId>
    <version>0.3.1</version>
</dependency>
```

## 🚀 Getting Started

### 1. Initialize
Use `MetrcFactory` to handle shared rate limiting effectively.

```kotlin
import io.github.thunkier.thunkmetrc.wrapper.MetrcFactory
import io.github.thunkier.thunkmetrc.wrapper.MetrcWrapper
import kotlinx.coroutines.runBlocking

fun main() = runBlocking {
    // Shared rate limiter (150 requests/sec integrator limit)
    val factory = MetrcFactory(maxConcurrentRequests = 150)

    val metrc = factory.create(
        baseUrl = "https://sandbox-api-or.metrc.com",
        vendorKey = "vendor_key",
        userKey = "user_key"
    )
    
    // Use metrc...
}
```

### 2. Make Requests
All service methods are `suspend` functions.

```kotlin
try {
    val facilities = metrc.facilities.getFacilities()
    facilities?.forEach { facility ->
        println("Facility: ${facility.License.Number}")
    }
} catch (e: Exception) {
    println("Error: ${e.message}")
}
```

### 3. Pagination (Flows)
Use the `iterate...` methods to get a `Flow` of items, automatically handling page traversal.

```kotlin
import kotlinx.coroutines.flow.collect

metrc.packages.iterateGetActivePackages(
    licenseNumber = "123-ABC-LICENSE",
    lastModifiedStart = null,
    lastModifiedEnd = null
).collect { pkg ->
    println("Package: ${pkg.Label}")
}
```

## 🛡️ Rate Limiting
The SDK uses `MetrcRateLimiter` to enforce:
- **Integrator Limits**: Default 150/sec.
- **Facility Limits**: Default 50/sec per facility.
- **Backoff**: Exponential backoff on 429 Too Many Requests.

Configuration is handled via the `MetrcFactory` constructor.
