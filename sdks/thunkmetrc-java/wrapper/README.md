# ThunkMetrc Java Wrapper SDK

Type-safe, rate-limited wrapper for the Metrc API.

## 📦 Installation
Add the dependency to your `pom.xml`:

```xml
<dependency>
    <groupId>io.github.thunkier</groupId>
    <artifactId>thunkmetrc-wrapper</artifactId>
    <version>0.3.1</version>
</dependency>
```

## 🚀 Getting Started

### 1. Initialize
Use the `MetrcFactory` to create a `MetrcWrapper`. This ensures efficient resource sharing (rate limits) across instances.

```java
import io.github.thunkier.thunkmetrc.wrapper.MetrcFactory;
import io.github.thunkier.thunkmetrc.wrapper.MetrcWrapper;

public class Main {
    public static void main(String[] args) {
        // Initialize factory with global concurrency limit (default 150)
        MetrcFactory factory = new MetrcFactory(150);

        // Create a wrapper instance
        MetrcWrapper metrc = factory.create(
            "https://sandbox-api-or.metrc.com",
            "vendor_key",
            "user_key"
        );
    }
}
```

### 2. Make Requests
Service methods are strongly typed.

```java
try {
    var facilities = metrc.facilities.getFacilities();
    for (var facility : facilities) {
        System.out.println(facility.License.Number);
    }
} catch (Exception e) {
    e.printStackTrace();
}
```

### 3. Pagination
For endpoints that support pagination, use the `iterate` helpers to lazily fetch all pages.

```java
// Iterate over all active packages
var iterator = metrc.packages.iterateGetActivePackages(
    "123-ABC-LICENSE",
    null, // lastModifiedStart
    null  // lastModifiedEnd
);

iterator.forEachRemaining(pkg -> {
    System.out.println("Package: " + pkg.Label);
});
```

## 🛡️ Rate Limiting
The SDK automatically handles Metrc's rate limits:
- **Integrator Limit**: 150 requests/second via `MetrcFactory`.
- **Facility Limit**: 50 requests/second/facility.
- **Concurrency**: Managed internally to prevent 429s.
- **Retries**: Built-in exponential backoff for 429/500 errors.

To customize configuration, modify `MetrcRateLimiter` internal config if needed (currently exposed via Factory constructor arg).
