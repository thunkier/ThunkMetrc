# Development & Architecture

## Repository Structure

- `sdks/`: Public SDK outputs, grouped by language family.
  - `sdks/thunkmetrc-cs`
  - `sdks/thunkmetrc-go`
  - `sdks/thunkmetrc-java`
  - `sdks/thunkmetrc-kotlin`
  - `sdks/thunkmetrc-py`
  - `sdks/thunkmetrc-rs`
  - `sdks/thunkmetrc-ts`
- `specs/`: Spec source and generated artifacts.
  - `specs/source/scraped/html`: HTML docs fetched from Metrc state sites.
  - `specs/source/scraped/postman`: Postman collections fetched per state.
  - `specs/generated/bruno`: Generated Bruno collection used by downstream generators.
  - `specs/generated/schemas`: Generated response/request schemas and manifest.
- `tools/`: Monorepo generator/analyzer CLI.
  - `gen specs`: Scrape docs and generate Bruno files.
  - `gen schemas`: Generate schema models from responses/docs.
  - `gen internal`: Generate probe internal client/models.
  - `gen sdks`: Generate public SDKs.
- `probe/`: Integration runner against mock server or real Metrc.
- `mockserver/`: Local mock API server.

## Generation Flow

1. `task gen:spec`
   - Reads scraped data from Metrc documentation.
   - Writes Bruno files to `specs/generated/bruno`.
2. `task gen:schema`
   - Builds schema files and `manifest.json` under `specs/generated/schemas/public`.
3. `task gen:internal`
   - Uses Bruno + internal schemas to generate probe client/models.
4. `task gen:sdks`
   - Uses Bruno + manifest + versions to generate all language SDKs.

## Versioning

- `versions.yaml` is the source of truth for `sdk_version` and dependency versions used by generators.

## Common Commands

- `task run`: Run probe workflow (dry-run by default).
- `task setup-sandbox`: Setup/validate integrator sandbox user.
- `task gen:all`: Run full generation pipeline (`spec -> internal -> sdks`).
- `task build:all`: Build all SDKs.
- `task test:all`: Run SDK test suites.
- `task publish:all`: Generate and publish all SDK packages.

## Specgen Controls

`tools gen specs` supports live-scrape controls:

- `-skip-scrape`: Use local scraped files only.
- `-states`: Comma-separated state codes (overrides config).
- `-timeout-seconds`: HTTP timeout.
- `-max-retries`: Request retries.
- `-retry-backoff-ms`: Retry backoff.
- `-concurrency`: Parallel state fetch workers.
- `-user-agent`: Custom User-Agent string.

These can also be configured in `config.json` under `Scraper`.
