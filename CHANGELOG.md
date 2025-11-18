# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2024-11-18

### Added

- Initial release of terraform-provider-shared
- Schema builders for common attribute patterns:
  - `StandardID()`, `Timestamp()` for computed attributes
  - `RequiredString()`, `OptionalString()` for basic strings
  - `SensitiveString()` for sensitive data
  - `StringEnum()` for validated enums
  - Boolean attributes with defaults
  - Integer/port attributes
  - Pattern builders for common resource types
- Validators for common patterns:
  - Status validators
  - Provider type validators
  - Authentication method validators
  - Environment validators
  - Severity/priority validators
- Standardized error handling:
  - `APIError()`, `NotFoundError()`, `ValidationError()`
  - HTTP status code helpers
  - Diagnostic error helpers
- Base resource implementation:
  - `BaseResource` for common resource functionality
  - Configuration management
  - Error handling utilities
- Utility functions:
  - Time/timestamp utilities with standard format (RFC850)
  - Type conversion functions (string, bool, int, etc.)
  - Safe value extraction functions
  - Slice-to-types conversions
- GitHub workflows:
  - Test workflow for multiple Go versions
  - Linting and formatting checks
  - Release workflow
- Project structure and documentation:
  - README with usage examples
  - Contributing guidelines
  - golangci-lint configuration
