# Terraform Provider Shared

[![shield_gh-workflow-test]][link_gh-workflow-test]
[![shield_license]][license_file]

---

Shared libraries, utilities, and common code for Sonatype Terraform Providers.

This repository contains reusable schema builders, validators, error handling, and common patterns used across multiple Terraform providers:
- [terraform-provider-sonatypeiq](https://github.com/sonatype-nexus-community/terraform-provider-sonatypeiq)
- [terraform-provider-sonatyperepo](https://github.com/sonatype-nexus-community/terraform-provider-sonatyperepo)

## Features

- **Schema Builders**: Utility functions for creating consistent Terraform schema attributes
- **Validators**: Pre-built validation rules for common patterns (IDs, emails, enums, etc.)
- **Error Handling**: Standardized error response handling across providers
- **Base Resources**: Common base resource implementations
- **Helper Functions**: Reusable utilities for API operations and data conversion

## Usage

Add this module as a dependency to your Terraform provider:

```go
require github.com/sonatype-nexus-community/terraform-provider-shared v0.1.0
```

## Examples

Comprehensive examples are available in the [`examples`](./examples) directory:

- [**Schema Builders**](./examples/schema_builders.md) - Creating consistent Terraform resource schema attributes
- [**Validators**](./examples/validators.md) - Enforcing constraints on attribute values (enums, patterns, ranges)
- [**Type Conversions**](./examples/conversion.md) - Converting between Go types and Terraform framework types
- [**Error Handling**](./examples/error_handling.md) - Standardized error messages and HTTP status code handling

## Development

### Building

```bash
go build -v .
```

### Testing

```bash
go test -v -cover -timeout=5m ./...
```

### Linting

```bash
go fmt ./...
go vet ./...
```

See [CONTRIBUTING.md](./CONTRIBUTING.md) for details.

## The Fine Print

This project is part of the [Sonatype Nexus Community](https://github.com/sonatype-nexus-community) organization, which is not officially supported by Sonatype. Please review the latest pull requests, issues, and commits to understand this project's readiness for contribution and use.

* File suggestions and requests on this repo through GitHub Issues, so that the community can pitch in
* Use or contribute to this project according to your organization's policies and your own risk tolerance
* Don't file Sonatype support tickets related to this project— it won't reach the right people that way

Last but not least of all - have fun!

[shield_gh-workflow-test]: https://img.shields.io/github/actions/workflow/status/sonatype-nexus-community/terraform-provider-shared/test.yml?branch=main&logo=GitHub&logoColor=white "build"
[shield_license]: https://img.shields.io/github/license/sonatype-nexus-community/terraform-provider-shared?logo=open%20source%20initiative&logoColor=white "license"

[link_gh-workflow-test]: https://github.com/sonatype-nexus-community/terraform-provider-shared/actions/workflows/test.yml?query=branch%3Amain
[license_file]: https://github.com/sonatype-nexus-community/terraform-provider-shared/blob/main/LICENSE
