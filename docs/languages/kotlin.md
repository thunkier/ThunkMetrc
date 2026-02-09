# Kotlin SDK Guide

## Dependencies

Maven:

```xml
<dependency>
  <groupId>io.github.thunkier</groupId>
  <artifactId>thunkmetrc-kotlin-client</artifactId>
  <version>0.2.2</version>
</dependency>
<dependency>
  <groupId>io.github.thunkier</groupId>
  <artifactId>thunkmetrc-kotlin-wrapper</artifactId>
  <version>0.2.2</version>
</dependency>
<!-- Optional high-level helpers -->
<dependency>
  <groupId>io.github.thunkier</groupId>
  <artifactId>thunkmetrc-kotlin</artifactId>
  <version>0.2.2</version>
</dependency>
```

Gradle Kotlin DSL:

```kotlin
implementation("io.github.thunkier:thunkmetrc-kotlin-client:0.2.2")
implementation("io.github.thunkier:thunkmetrc-kotlin-wrapper:0.2.2")
implementation("io.github.thunkier:thunkmetrc-kotlin:0.2.2") // optional
```

## Recommended Setup (Wrapper + Factory)

```kotlin
import io.github.thunkier.thunkmetrc.wrapper.MetrcFactory
import kotlinx.coroutines.runBlocking

fun main() = runBlocking {
    val factory = MetrcFactory(maxConcurrentRequests = 150)
    val wrapper = factory.create(
        baseUrl = "https://sandbox-api-or.metrc.com",
        vendorKey = "vendor_key",
        userKey = "user_key"
    )

    val facilities = wrapper.facilities.getFacilities().orEmpty()
    println("Facilities: ${facilities.size}")
}
```

## Pagination Pattern

```kotlin
import io.github.thunkier.thunkmetrc.wrapper.models.PackagesPackage

suspend fun loadAllActivePackages(wrapper: io.github.thunkier.thunkmetrc.wrapper.MetrcWrapper): List<PackagesPackage> {
    val all = mutableListOf<PackagesPackage>()
    var page = 1

    while (true) {
        val response = wrapper.packages.getActivePackages(
            lastModifiedEnd = null,
            lastModifiedStart = null,
            licenseNumber = "C12-0000001-LIC",
            pageNumber = page.toString(),
            pageSize = "20"
        )

        val data = response?.Data.orEmpty()
        if (data.isEmpty()) break

        all.addAll(data)
        page++
    }

    return all
}
```

## High-Level Helpers (`thunkmetrc-kotlin`)

```kotlin
import io.github.thunkier.thunkmetrc.ThunkMetrc
import java.time.OffsetDateTime

val thunk = ThunkMetrc(wrapper)
val packages = thunk.activePackagesInventorySync(
    licenseNumber = "C12-0000001-LIC",
    lastModifiedStart = OffsetDateTime.now().minusHours(1),
    bufferMinutes = 5
)
```

## Rate Limiter Configuration

```kotlin
import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcWrapper
import io.github.thunkier.thunkmetrc.wrapper.RateLimiterConfig

val client = MetrcClient("https://sandbox-api-or.metrc.com", "vendor_key", "user_key")
val config = RateLimiterConfig(enabled = true, maxGetPerSecondIntegrator = 150.0)
val configuredWrapper = MetrcWrapper(client, config)
```

## Build Locally

```bash
task build:kotlin
task test:kotlin
```
