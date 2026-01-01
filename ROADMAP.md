# Roadmap

This document outlines the planned development path for the Zama Golang Relayer SDK.

## Current Version: v0.1.0-rc.1

**Status:** Release Candidate - API stabilizing

### What's Included

- Mock encryption (returns deterministic handles)
- Core SDK structure and types
- Sepolia testnet configuration
- Input encryption API (Add64, AddBool, AddAddress)
- Network client with retry logic
- Comprehensive test suite

### Known Limitations

- Encryption is mocked (not real TFHE)
- No decryption support
- No mainnet configuration
- No EIP-712 signing utilities

---

## v0.2.0 - Real TFHE Encryption

**Target:** Q1 2026

### Features

- [ ] Real TFHE encryption via CGO bindings
- [ ] Integration with tfhe-rs library
- [ ] Live Sepolia testnet testing
- [ ] Performance benchmarks
- [ ] Memory optimization

### Technical

- CGO bindings to Zama's tfhe-rs
- Build tags for CGO/no-CGO modes
- Cross-compilation support (Linux, macOS, Windows)

---

## v0.3.0 - Decryption Support

**Target:** Q2 2026

### Features

- [ ] User decryption (with private key)
- [ ] Public decryption (via gateway)
- [ ] Re-encryption support
- [ ] Key derivation utilities

### Technical

- EIP-712 typed data signing
- Keypair generation and management
- Secure key storage recommendations

---

## v0.4.0 - Production Ready

**Target:** Q2 2026

### Features

- [ ] Mainnet configuration
- [ ] Additional integer types (uint8, uint16, uint32, uint128, uint256)
- [ ] Batch operations optimization
- [ ] Connection pooling
- [ ] Metrics and observability

### Technical

- Production hardening
- Security audit preparation
- Performance optimization
- Documentation completion

---

## v1.0.0 - Stable Release

**Target:** Q3 2026

### Requirements

- [ ] Security audit completed
- [ ] API stability guarantee
- [ ] Full documentation
- [ ] Production deployments validated
- [ ] Community feedback incorporated

### Guarantees

- Semantic versioning compliance
- Backwards compatibility within major version
- Long-term support (LTS) consideration

---

## Future Considerations

### Potential Features

- WebAssembly (WASM) compilation target
- gRPC API support
- Prometheus metrics exporter
- OpenTelemetry tracing
- Hardware Security Module (HSM) integration

### Community Requests

We track feature requests via [GitHub Issues](https://github.com/PrivaraXYZ/zama-golang-relayer-sdk/issues).

---

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md) for how to contribute to this roadmap.

## Disclaimer

This roadmap represents our current plans and is subject to change based on:
- Community feedback
- Technical discoveries
- Resource availability
- Zama protocol updates

Dates are estimates and not commitments.
