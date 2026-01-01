# Contributing to Zama Golang Relayer SDK

Thank you for your interest in contributing to the Zama Golang Relayer SDK!

## Important: AI Contribution Policy

**This project does NOT accept AI-assisted code contributions.**

Before contributing, please read our [AI Policy](./.claude/README.md).

### Why?

This SDK handles cryptographic operations for Fully Homomorphic Encryption (FHE). AI-generated code may introduce:

1. Subtle cryptographic bugs that are difficult to detect
2. Security vulnerabilities in encryption flows
3. Incorrect implementations of TFHE primitives
4. Memory safety issues

### Certification Requirement

All pull requests must include this certification:

```
I certify that:
- This contribution was written entirely by me without AI assistance
- I have read and understood the Zama FHEVM protocol documentation
- I have tested this code thoroughly
- I understand the cryptographic implications of my changes

Signed: [Your Name]
Date: [YYYY-MM-DD]
```

## Getting Started

### Prerequisites

- Go 1.23 or later
- Make
- golangci-lint
- pre-commit

### Setup

```bash
# Clone the repository
git clone https://github.com/PrivaraXYZ/zama-golang-relayer-sdk.git
cd zama-golang-relayer-sdk

# Install dependencies
go mod download

# Install pre-commit hooks
make setup-hooks

# Run tests
make test

# Run linter
make lint
```

## Development Workflow

### 1. Create a Branch

```bash
git checkout -b feature/your-feature-name
```

### 2. Make Changes

- Follow the code style in [CODE_STYLE.md](./CODE_STYLE.md) (if exists)
- Keep changes focused and atomic
- Write tests for new functionality
- Update documentation as needed

### 3. Test Your Changes

```bash
# Run all tests
make test

# Run tests with coverage
make coverage

# Run linter
make lint-strict

# Run pre-commit checks
make pre-commit
```

### 4. Commit Your Changes

Follow conventional commit format:

```
feat: add new encryption type support
fix: correct handle generation for addresses
docs: update API documentation
test: add tests for edge cases
chore: update dependencies
```

### 5. Submit a Pull Request

- Fill out the PR template completely
- Include the AI certification statement
- Link any related issues
- Request review from maintainers

## Code Guidelines

### Style

- Use `gofmt` and `goimports` (enforced by pre-commit)
- Follow Go idioms and conventions
- Keep functions focused and small
- Use meaningful variable names

### Documentation

- Document all exported types and functions
- Keep comments minimal but meaningful
- Update README for user-facing changes

### Testing

- Write table-driven tests
- Test edge cases and error conditions
- Aim for >80% coverage
- Include benchmarks for performance-critical code

### Security

- Never commit secrets or credentials
- Follow secure coding practices
- Report security issues privately to security@privara.xyz

## Review Process

1. Automated checks must pass (CI, linting, tests)
2. At least one maintainer approval required
3. Cryptographic changes require additional expert review
4. Changes may require live testing on testnet

## Questions?

- Open a [GitHub Discussion](https://github.com/PrivaraXYZ/zama-golang-relayer-sdk/discussions) for questions
- Open a [GitHub Issue](https://github.com/PrivaraXYZ/zama-golang-relayer-sdk/issues) for bugs or features
- Email security@privara.xyz for security concerns

## License

By contributing, you agree that your contributions will be licensed under the BSD-3-Clause License.
