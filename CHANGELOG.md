# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0-rc.1] - 2026-01-01

### Added

- Initial SDK implementation with mock encryption
- Core types: `FhevmInstance`, `EncryptedInput`, `EncryptResult`
- Sepolia testnet configuration
- Network client with retry logic and exponential backoff
- Ethereum address validation and checksumming
- Public key caching with TTL support
- Chainable builder API for encrypted inputs:
  - `Add64()` - Encrypt uint64 values
  - `AddBool()` - Encrypt boolean values
  - `AddAddress()` - Encrypt Ethereum addresses
- Comprehensive test suite with >80% coverage
- GitHub Actions CI/CD workflows
- Pre-commit hooks for code quality
- Examples for basic and batch encryption

### Notes

- This is an alpha release with mock encryption
- Real TFHE encryption will be added in v0.2.0
- API is subject to change before v1.0.0

## Future Releases

See [ROADMAP.md](./ROADMAP.md) for planned features.

[Unreleased]: https://github.com/PrivaraXYZ/zama-golang-relayer-sdk/compare/v0.1.0-rc.1...HEAD
[0.1.0-rc.1]: https://github.com/PrivaraXYZ/zama-golang-relayer-sdk/releases/tag/v0.1.0-rc.1
