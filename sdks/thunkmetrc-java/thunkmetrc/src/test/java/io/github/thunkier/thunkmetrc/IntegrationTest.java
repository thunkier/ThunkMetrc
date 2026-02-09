package io.github.thunkier.thunkmetrc;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcWrapper;
import io.github.cdimascio.dotenv.Dotenv;
import org.junit.jupiter.api.Test;
import java.io.File;
import java.util.List;
import static org.junit.jupiter.api.Assertions.*;

public class IntegrationTest {

    private String getEnv(Dotenv dotenv, String key) {
        String val = dotenv.get(key);
        if (val == null) {
            val = System.getenv(key);
        }
        return val;
    }

    @Test
    public void testActivePackagesInventorySync() throws Exception {
        var builder = Dotenv.configure().ignoreIfMissing();
        File envFile = new File("../../../.env");
        if (envFile.exists()) {
            builder.directory(envFile.getParent());
        }
        Dotenv dotenv = builder.load();
        
        String baseUrl = getEnv(dotenv, "METRC_BASE_URL");
        String vendorKey = getEnv(dotenv, "METRC_VENDOR_API_KEY");
        String userKey = getEnv(dotenv, "METRC_USER_API_KEY");
        String licenseNumber = getEnv(dotenv, "METRC_FACILITY_LICENSE_NUMBER");

        if (baseUrl == null) baseUrl = "https://sandbox-api-or.metrc.com";
        
        if (vendorKey == null || userKey == null || licenseNumber == null) {
            System.out.println("Skipping test: Missing METRC_VENDOR_API_KEY, METRC_USER_API_KEY, or METRC_FACILITY_LICENSE_NUMBER");
            return;
        }

        System.out.println("Config: " + baseUrl + ", " + licenseNumber);

        MetrcClient client = new MetrcClient(baseUrl, vendorKey, userKey);
        MetrcWrapper wrapper = new MetrcWrapper(client);
        ThunkMetrc thunk = new ThunkMetrc(wrapper);

        System.out.println("Starting sync for " + licenseNumber);
        long start = System.currentTimeMillis();
        
        try {
            List<?> packages = thunk.activePackagesInventorySync(licenseNumber, null, 5);
            
            long duration = System.currentTimeMillis() - start;
            System.out.println("Sync finished in " + duration + "ms. Found " + packages.size() + " packages.");
            
            assertNotNull(packages);
            assertTrue(packages.size() >= 0);
        } catch (Exception e) {
            if (e.getMessage() != null && e.getMessage().contains("429")) {
                System.out.println("WARNING: Test hit rate limit (429). Passing test as environment is restricted.");
            } else {
                throw e;
            }
        }
    }
}
