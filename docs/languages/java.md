# Java Integration Guide

## 📦 Installation

Add the dependency to your `pom.xml`.

```xml
<dependency>
    <groupId>com.thunkmetrc</groupId>
    <artifactId>thunkmetrc-wrapper</artifactId>
    <version>0.1.0</version>
</dependency>

<!-- Optional: High Level -->
<dependency>
    <groupId>com.thunkmetrc</groupId>
    <artifactId>thunkmetrc</artifactId>
    <version>0.1.0</version>
</dependency>
```

---

## 🚀 Getting Started

### 1. Initialize

```java
import com.thunkmetrc.client.MetrcClient;
import com.thunkmetrc.wrapper.MetrcWrapper;

public class Main {
    public static void main(String[] args) {
        MetrcClient client = new MetrcClient(
            "https://sandbox-api-or.metrc.com",
            "vendor_key",
            "user_key"
        );

        MetrcWrapper wrapper = new MetrcWrapper(client);
    }
}
```

### 2. Make Requests

```java
try {
    Object response = wrapper.transfersGetIncomingV2(
        "2023-01-01", // lastModifiedEnd
        "2023-01-02", // lastModifiedStart
        "C12-0000001-LIC", 
        null, // pageNumber
        null  // pageSize
    );
    
    System.out.println(response);
} catch (Exception e) {
    e.printStackTrace();
}
```

---

## 🛡️ Rate Limiting

```java
import com.thunkmetrc.wrapper.RateLimiterConfig;

RateLimiterConfig config = new RateLimiterConfig();
config.enabled = true;
config.maxGetPerSecondPerFacility = 50.0;

MetrcWrapper wrapper = new MetrcWrapper(client, config);
```

---

## 🛠️ High-Level Features

### Inventory Sync

```java
import com.thunkmetrc.ThunkMetrc;

ThunkMetrc thunk = new ThunkMetrc(wrapper);

List<Object> packages = thunk.activePackagesInventorySync(
    "C12-0000001-LIC",
    ZonedDateTime.now().minusDays(1),
    15
);
```
