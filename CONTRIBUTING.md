# Contributing

This repository is part of the Sonatype Nexus Community and welcomes contributions from the community.

## Getting Started

1. **Fork** the repository on GitHub
2. **Clone** your fork locally
3. **Create** a feature branch (`git checkout -b feature/amazing-feature`)
4. **Commit** your changes (`git commit -S -s -m 'Add some amazing feature'`)
5. **Push** to the branch (`git push origin feature/amazing-feature`)
6. **Open** a Pull Request

## Development Setup

### Prerequisites

- Go 1.23.7 or later
- golangci-lint (optional, but recommended)

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
golangci-lint run ./...
```

## GitHub Actions Checks

All pull requests are automatically checked by GitHub Actions. The following checks must pass before merging:

### 1. Tests with Coverage Verification

- Runs all tests with coverage tracking
- **Requirement**: Test coverage must be at least 60%
- **Local equivalent**:
  ```bash
  go test -v -coverprofile=coverage.out -timeout=5m ./...
  go tool cover -func=coverage.out | grep total
  ```

### 2. Code Formatting (go fmt)

- Ensures all code is properly formatted according to Go standards
- **Local equivalent**:
  ```bash
  go fmt ./...
  ```
  Note: If `go fmt` makes changes, commit those changes before pushing.

### 3. Vet Analysis (go vet)

- Catches common coding mistakes and suspicious constructs
- **Local equivalent**:
  ```bash
  go vet ./...
  ```

### 4. Sonatype Lifecycle Scan

- Performs security and supply chain analysis
- Only runs on push to main and tags
- Verifies dependencies for known vulnerabilities

## Pre-Submission Checklist

Before submitting a pull request, ensure:

1. All tests pass locally:
   ```bash
   go test -v -timeout=5m ./...
   ```

2. Coverage is at least 60%:
   ```bash
   go test -coverprofile=coverage.out ./...
   go tool cover -func=coverage.out | grep total
   ```

3. Code is properly formatted:
   ```bash
   go fmt ./...
   ```

4. No vet issues:
   ```bash
   go vet ./...
   ```

5. Commit messages follow the format:
   - Use `-S` flag for signing commits: `git commit -S`
   - Use `-s` flag for DCO: `git commit -s`
   - Use `-m` with a clear message: `git commit -m 'Clear description of changes'`

## Guidelines

- Ensure all tests pass before submitting a PR
- Add tests for new functionality
- Follow the existing code style and conventions
- Update documentation as needed
- Keep commits focused and descriptive
- Link any related issues in your PR description

## Code Style

This project follows standard Go conventions:
- Run `go fmt` on all code
- Use `golangci-lint` for linting
- Keep functions small and focused
- Write clear, descriptive names

## Testing

All new code should include tests. Run the test suite with:

```bash
go test ./...
```

## Documentation

- Update README.md if adding new public functions
- Add inline documentation for exported functions
- Include examples where appropriate

## Questions or Issues?

- Open a GitHub Issue for bugs or feature requests
- Start a discussion for questions
- Be respectful and constructive in all communications

Thank you for contributing!
