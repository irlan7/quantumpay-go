# 🔒 Security Policy

## 1. Supported Versions
This repository contains the **core protocol implementation** of the QuantumPay blockchain.  
Only the latest tagged release (e.g., `v1.x.x-core`) receives security updates.

| Version | Status | Notes |
|----------|---------|-------|
| `main` / `master` | ✅ Active | Development branch |
| `v1.x.x-core` | ✅ Maintained | Stable releases |
| Older versions | ❌ Unsupported | Please upgrade |

## 2. Reporting Vulnerabilities
If you discover a security vulnerability **in the core protocol**, please report it **privately**.

### Contact:
📧 quantumpaysec [at] gmail [dot] com
(If unavailable, please use GitHub’s private “Security Advisory” feature.)

Do **not** open public issues for security reports.

We appreciate responsible disclosure and aim to acknowledge reports within **48 hours**.

## 3. Scope of Responsibility
This repository maintains:
- Core consensus logic
- Economic and cryptographic primitives
- Transaction validation
- Node networking (p2p/gRPC)
- Genesis configuration and chain parameters

**Out of scope**:
- Wallet frontends
- Exchanges or dApps
- Third-party smart contracts

## 4. Responsible Disclosure Process
1. Submit the issue privately (email or GitHub Security Advisory)
2. Provide a minimal proof of concept
3. Our engineering and research team will review internally
4. If confirmed, a patch will be prepared and tested
5. Public disclosure after patch release, with attribution (if permitted)

## 5. Security Philosophy
QuantumPay Core follows a **“deterministic safety-first”** model:
- Minimal dependencies
- Deterministic builds
- Isolation of consensus logic
- No auto-update or hidden telemetry
- Reproducible build verification

We believe transparency, reproducibility, and independent review are the foundation of blockchain security.

## 6. Credits
QuantumPay thanks all independent auditors, researchers, and community contributors who help ensure protocol security.


*© 2026 QuantumPay Foundation — Security & Protocol Research Division*
