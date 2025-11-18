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
