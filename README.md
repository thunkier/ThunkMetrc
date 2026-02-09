# ThunkMetrc

ThunkMetrc is a generator-driven SDK platform for the Metrc API. This repository contains the scraping/spec pipeline, schema generation, internal tooling models, and publishable SDKs for C#, Go, TypeScript, Python, Java, Kotlin, and Rust.

## Developer Navigation
- [Quick Start](#quick-start)
- [SDK Package Matrix](#sdk-package-matrix)
- [Generation Pipeline](#generation-pipeline)
- [Live Scrape Workflow](#live-scrape-workflow)
- [Publishing Workflow](#publishing-workflow)
- [Language Guides](#language-guides)
- [Core Docs](#core-docs)

## Quick Start

Recommended pattern:
1. Use the language-specific `wrapper` package for typed API calls.
2. Use the `client` package only when you want raw request/response control.
3. Use the high-level `thunkmetrc` package for convenience sync/helpers where available.

TypeScript wrapper example:

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

## SDK Package Matrix

| Language | Client | Wrapper (recommended) | High-level package |
| --- | --- | --- | --- |
| C# | `ThunkMetrc.Client` | `ThunkMetrc.Wrapper` | `ThunkMetrc` |
| Go | `github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client` | `github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper` | `github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/thunkmetrc` |
| TypeScript | `@thunkier/thunkmetrc-client` | `@thunkier/thunkmetrc-wrapper` | `@thunkier/thunkmetrc` |
| Python | `thunkmetrc-client` | `thunkmetrc-wrapper` | `thunkmetrc` |
| Java | `io.github.thunkier:thunkmetrc-client` | `io.github.thunkier:thunkmetrc-wrapper` | `io.github.thunkier:thunkmetrc` |
| Kotlin | `io.github.thunkier:thunkmetrc-kotlin-client` | `io.github.thunkier:thunkmetrc-kotlin-wrapper` | `io.github.thunkier:thunkmetrc-kotlin` |
| Rust | `thunkmetrc-client` | `thunkmetrc-wrapper` | `thunkmetrc` |

Note: Go module paths now live under `sdks/thunkmetrc-go/...`.

## Generation Pipeline

The end-to-end flow is:
1. Live scrape source docs into `specs/source/scraped/postman/`.
2. Build normalized Bruno specs in `specs/generated/bruno/`.
3. Generate schemas in `specs/generated/schemas/`.
4. Generate internal tooling assets.
5. Generate SDKs under `sdks/thunkmetrc-*`.

Most-used commands:

```bash
task gen:spec
task gen:schema
task gen:internal
task gen:sdks
task build:all
task test:all
```

## Live Scrape Workflow

`task gen:spec` proxies to `tools gen specs` and supports scraper controls.

Examples:

```bash
# Scrape selected states only
task gen:spec -- --states=ca,co,or

# Tune HTTP behavior
task gen:spec -- --states=ca --timeout-seconds=60 --max-retries=4 --retry-backoff-ms=800

# Regenerate from existing local scrape files
task gen:spec -- -skip-scrape
```

Available knobs:
- `-states`
- `-concurrency`
- `-timeout-seconds`
- `-max-retries`
- `-retry-backoff-ms`
- `-user-agent`
- `-skip-scrape`

## Publishing Workflow

Publishing is handled by `.github/workflows/publish-sdks.yml`.

Trigger modes:
- Tag push matching `v*`
- Manual `workflow_dispatch`

Each publish job regenerates codegen outputs before publishing language packages.

## Language Guides

- [C#](docs/languages/csharp.md)
- [Go](docs/languages/go.md)
- [TypeScript](docs/languages/typescript.md)
- [Python](docs/languages/python.md)
- [Java](docs/languages/java.md)
- [Kotlin](docs/languages/kotlin.md)
- [Rust](docs/languages/rust.md)

## Core Docs

- [Configuration](docs/CONFIGURATION.md)
- [Development](docs/DEVELOPMENT.md)
- [API Coverage](docs/API_COVERAGE.md)

## License

MIT. See [LICENSE](LICENSE).
