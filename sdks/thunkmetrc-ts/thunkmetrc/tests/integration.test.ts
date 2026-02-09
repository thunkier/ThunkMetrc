import { MetrcClient } from "@thunkier/thunkmetrc-client";
import { MetrcWrapper } from "@thunkier/thunkmetrc-wrapper";
import { activePackagesInventorySync } from "../src/inventorySync";
import * as dotenv from "dotenv";

dotenv.config({ path: "../../../.env" });

describe("ThunkMetrc Integration Tests", () => {
  const baseUrl = process.env.METRC_BASE_URL || "https://sandbox-api-or.metrc.com";
  const vendorKey = process.env.METRC_VENDOR_API_KEY;
  const userKey = process.env.METRC_USER_API_KEY;
  const licenseNumber = process.env.METRC_FACILITY_LICENSE_NUMBER;

  const hasCredentials = vendorKey && userKey && licenseNumber;

  test("Client and Wrapper modules export correctly", () => {
    expect(MetrcClient).toBeDefined();
    expect(MetrcWrapper).toBeDefined();
    expect(activePackagesInventorySync).toBeDefined();
  });

  test("Client can be instantiated", () => {
    const client = new MetrcClient(baseUrl, "test-vendor", "test-user");
    expect(client).toBeDefined();
  });

  test("Wrapper can be instantiated with Client", () => {
    const client = new MetrcClient(baseUrl, "test-vendor", "test-user");
    const wrapper = new MetrcWrapper(client);
    expect(wrapper).toBeDefined();
    expect(wrapper.client).toBe(client);
  });

  (hasCredentials ? test : test.skip)(
    "activePackagesInventorySync integrates with live API",
    async () => {
      console.log(`Config: ${baseUrl}, License: ${licenseNumber}`);

      const client = new MetrcClient(baseUrl, vendorKey!, userKey!);
      const wrapper = new MetrcWrapper(client);

      console.log(`Starting sync for license: ${licenseNumber}...`);
      const startTime = Date.now();

      try {
        const packagesMap = await activePackagesInventorySync(
          [{ licenseNumber: licenseNumber!, wrapper }],
          undefined,
          5
        );

        const duration = (Date.now() - startTime) / 1000;
        const packages = packagesMap.get(licenseNumber!) || [];
        const packageCount = packages.length;
        console.log(`Sync completed in ${duration.toFixed(2)}s. Found ${packageCount} packages.`);

        expect(packagesMap).toBeDefined();
        expect(packagesMap.size).toBeGreaterThan(0);
        expect(Array.isArray(packages)).toBe(true);

        if (packageCount > 0) {
          console.log("First package sample:", packages[0]);
        }
      } catch (e: any) {
        if (e.message?.includes("429")) {
          console.log("WARNING: Test hit rate limit (429). Passing test as environment is restricted.");
        } else {
          throw e;
        }
      }
    },
    60000 // 60 second timeout
  );
});
