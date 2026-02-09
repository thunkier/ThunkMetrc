# Java SDK Guide

## Dependencies

Maven:

```xml
<dependency>
  <groupId>io.github.thunkier</groupId>
  <artifactId>thunkmetrc-client</artifactId>
  <version>0.2.2</version>
</dependency>
<dependency>
  <groupId>io.github.thunkier</groupId>
  <artifactId>thunkmetrc-wrapper</artifactId>
  <version>0.2.2</version>
</dependency>
<!-- Optional high-level helpers -->
<dependency>
  <groupId>io.github.thunkier</groupId>
  <artifactId>thunkmetrc</artifactId>
  <version>0.2.2</version>
</dependency>
```

## Recommended Setup (Wrapper + Factory)

```java
import io.github.thunkier.thunkmetrc.wrapper.MetrcFactory;
import io.github.thunkier.thunkmetrc.wrapper.MetrcWrapper;

MetrcFactory factory = new MetrcFactory(150);
MetrcWrapper wrapper = factory.create(
    "https://sandbox-api-or.metrc.com",
    "vendor_key",
    "user_key"
);

var facilities = wrapper.facilities().getFacilities();
System.out.println("Facilities: " + facilities.size());
```

## Pagination Pattern

```java
import io.github.thunkier.thunkmetrc.wrapper.models.PackagesPackage;
import io.github.thunkier.thunkmetrc.wrapper.models.PaginatedResponse;
import java.util.ArrayList;
import java.util.List;

String license = "C12-0000001-LIC";
int page = 1;
List<PackagesPackage> all = new ArrayList<>();

while (true) {
    PaginatedResponse<PackagesPackage> response = wrapper.packages().getActivePackages(
        null, // lastModifiedEnd
        null, // lastModifiedStart
        license,
        Integer.toString(page),
        "20"
    );

    if (response == null || response.Data == null || response.Data.isEmpty()) {
        break;
    }

    all.addAll(response.Data);
    page++;
}
```

## High-Level Helpers (`thunkmetrc`)

```java
import io.github.thunkier.thunkmetrc.ThunkMetrc;
import io.github.thunkier.thunkmetrc.wrapper.models.PackagesPackage;

ThunkMetrc thunk = new ThunkMetrc(wrapper);
List<PackagesPackage> packages = thunk.activePackagesInventorySync(
    "C12-0000001-LIC",
    java.time.OffsetDateTime.now().minusHours(1),
    5
);
```

## Rate Limiter Configuration

```java
import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;
import io.github.thunkier.thunkmetrc.wrapper.MetrcWrapper;

MetrcRateLimiter.Config config = new MetrcRateLimiter.Config();
config.enabled = true;
config.maxGetPerSecondIntegrator = 150.0;

MetrcClient client = new MetrcClient("https://sandbox-api-or.metrc.com", "vendor_key", "user_key");
MetrcWrapper configuredWrapper = new MetrcWrapper(client, config);
```

## Build Locally

```bash
task build:java
task test:java
```
