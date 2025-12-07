# Kotlin Integration Guide

## 📦 Installation

Add to `pom.xml` (Maven) or `build.gradle.kts`.

```kotlin
implementation("com.thunkmetrc:thunkmetrc-kotlin-wrapper:0.1.0")
implementation("com.thunkmetrc:thunkmetrc-kotlin:0.1.0") // High-level
```

---

## 🚀 Getting Started

### 1. Initialize

```kotlin
import com.thunkmetrc.client.MetrcClient
import com.thunkmetrc.wrapper.MetrcWrapper

val client = MetrcClient(
    "https://sandbox-api-or.metrc.com",
    "vendor_key",
    "user_key"
)

val wrapper = MetrcWrapper(client)
```

### 2. Make Requests
The wrapper is `suspend` friendly.

```kotlin
import kotlinx.coroutines.runBlocking

fun main() = runBlocking {
    try {
        val result = wrapper.transfersGetIncomingV2(
            licenseNumber = "C12-0000001-LIC"
        )
        println(result)
    } catch (e: Exception) {
        println("Error: ${e.message}")
    }
}
```

---

## 🛡️ Rate Limiting

```kotlin
import com.thunkmetrc.wrapper.RateLimiterConfig

val config = RateLimiterConfig(enabled = true)
val wrapper = MetrcWrapper(client, config)
```

---

## 🛠️ High-Level Features

### Inventory Sync

```kotlin
import com.thunkmetrc.ThunkMetrc
import java.time.OffsetDateTime

val thunk = ThunkMetrc(wrapper)

val packages = thunk.activePackagesInventorySync(
    licenseNumber = "C12-0000001-LIC",
    lastKnownSync = OffsetDateTime.now().minusDays(1),
    bufferMinutes = 20
)
```
