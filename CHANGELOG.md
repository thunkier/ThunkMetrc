# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

> **Note**: Versions prior to 0.2.2 were developed in a private repository as an internal tool.
> The project was open-sourced starting with v0.2.2. Earlier version history is reconstructed
> from internal release notes; corresponding Git history is not available.

## [Unreleased]

## [0.2.2] - 2025-12-13
### Added
- Full Maven Central publishing configuration for Java/Kotlin SDKs
- README documentation for rate limiter default values
- Python thunkmetrc README.md
- This CHANGELOG (retroactively documenting prior releases)

### Fixed
- TypeScript package scope changed from `@thunkmetrc` to `@thunkier`
- Rust wrapper/thunkmetrc crate metadata (description, license)
- Kotlin groupId standardized to `io.github.thunkier`

### Changed
- Maven publishing uses Central Portal plugin 0.9.0
- Open-sourced project and published to GitHub

## [0.2.1] - 2025-12-12
### Added
- Documentation comments injected into all SDK generators from Bruno specs
- Java/Kotlin thunkmetrc high-level packages

### Fixed
- URL encoding for path parameters across all language generators

## [0.2.0] - 2025-12-11
### Added
- Initial SDK generation for C#, Go, TypeScript, Python, Java, Kotlin, Rust
- Wrapper packages with built-in rate limiting (token bucket + semaphore)
- `ActivePackagesInventorySync` utility function in high-level packages

### Changed
- Unified code generation from Bruno specification files

## [0.1.2] - 2025-11-28
### Added
- Rate limiter configuration options (enabled, max requests, concurrency limits)
- Query parameter support in all generators

### Fixed
- JSON serialization edge cases in Python and TypeScript clients

## [0.1.1] - 2025-11-15
### Added
- Kotlin SDK generator
- Rust SDK generator

### Fixed
- Authentication header formatting in Java client

## [0.1.0] - 2025-11-01
### Added
- Bruno specification parser (`automation/pkg/bruno`)
- SDK generator framework (`automation/cmd/sdkgen`)
- C#, Go, TypeScript, Python, Java client generators
- Basic wrapper generators with method delegation

## [0.0.3] - 2025-10-15
### Added
- Taskfile for build automation
- Multi-language output directory structure (`sdks/`)

### Changed
- Refactored generator to support grouped API endpoints

## [0.0.2] - 2025-10-01
### Added
- Initial Bruno collection structure (`specs/metrc-bruno`)
- Prototype TypeScript client generator

### Fixed
- Request body handling for POST/PUT methods

## [0.0.1] - 2025-09-15
### Added
- Project scaffolding
- Initial Go module setup for automation tooling
- Proof-of-concept code generation from Bruno files

[Unreleased]: https://github.com/thunkier/ThunkMetrc/compare/v0.2.2...HEAD
[0.2.2]: https://github.com/thunkier/ThunkMetrc/releases/tag/v0.2.2
