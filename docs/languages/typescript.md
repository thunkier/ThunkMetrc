# TypeScript SDK Guide

## Packages

```bash
npm install @thunkier/thunkmetrc-client
npm install @thunkier/thunkmetrc-wrapper
# Optional high-level helpers
npm install @thunkier/thunkmetrc
```

Requires modern Node.js (18+ recommended).

## Recommended Setup (Wrapper)

```ts
import { MetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcWrapper } from '@thunkier/thunkmetrc-wrapper';

const client = new MetrcClient({
  baseUrl: 'https://sandbox-api-or.metrc.com',
  vendorKey: process.env.METRC_VENDOR_KEY!,
  userKey: process.env.METRC_USER_KEY!,
});

const wrapper = new MetrcWrapper(client);
const facilities = await wrapper.facilities.getFacilities();
console.log(`Facilities: ${facilities.length}`);
```

## Raw Client

```ts
import { MetrcClient } from '@thunkier/thunkmetrc-client';

const client = new MetrcClient({
  baseUrl: 'https://sandbox-api-or.metrc.com',
  vendorKey: process.env.METRC_VENDOR_KEY!,
  userKey: process.env.METRC_USER_KEY!,
});

const facilities = await client.getFacilities();
```

## Pagination Pattern

Use explicit page loops for paginated endpoints:

```ts
const allPackages: any[] = [];
let page = 1;

for (;;) {
  const response = await wrapper.packages.getActivePackages(
    undefined,
    undefined,
    undefined,
    'C12-0000001-LIC',
    page,
    '20'
  );

  const data = response?.Data ?? [];
  if (data.length === 0) {
    break;
  }

  allPackages.push(...data);
  page += 1;
}
```

## High-Level Helpers (`@thunkier/thunkmetrc`)

`@thunkier/thunkmetrc` provides helper workflows such as `activePackagesInventorySync`.

```ts
import { activePackagesInventorySync } from '@thunkier/thunkmetrc';

const result = await activePackagesInventorySync(
  [{ licenseNumber: 'C12-0000001-LIC', wrapper }],
  new Date(Date.now() - 60 * 60 * 1000),
  5
);

const packages = result.get('C12-0000001-LIC') ?? [];
```

## Build Locally

```bash
task build:typescript
task test:typescript
```
