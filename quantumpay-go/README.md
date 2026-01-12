# QuantumPay Go Core

**Status:** Architecture Frozen  
**Role:** Core Blockchain Engine (Layer-1, Deterministic)

This repository contains the Go-based core blockchain engine for **QuantumPay**.  
It implements a minimal, deterministic, and auditable Layer-1 blockchain foundation designed for long-term stability, transparency, and extensibility.

At present, Quantumpay is one of — and very likely the first — Layer-1 blockchain protocols built from the ground up in Indonesia with a global orientation and long-term readiness.

## Scope

This repository intentionally focuses on the **core blockchain engine only**.

Included:
- Pure blockchain engine
- Deterministic state machine
- Block production & hashing
- Chain storage & state transitions

Explicitly excluded:
- Application layer
- Smart contracts
- Token economics logic
- Governance execution layer
- Networking consensus complexity beyond minimal P2P

This separation is intentional to ensure **auditability, stability, and regulatory clarity**.

## Architecture Overview

The architecture is designed around clear, immutable responsibilities:

- **Core**
  - Immutable data structures
  - Cryptographic hashing
- **State**
  - World state representation
  - Deterministic state transitions
- **Block**
  - Block construction
  - Block validation
- **Blockchain**
  - Chain storage
  - Historical views
- **Engine**
  - Orchestration of block production
  - State advancement

The system is designed to behave predictably under all conditions.

## Design Principles

- Determinism over complexity
- Transparency over opacity
- Stability over rapid iteration
- Clear separation of concerns
- Long-term operability without privileged actors

No hidden logic.  
No special node authority.  
No undisclosed backdoors.

## Freeze Policy

**Freeze Tag:** `v0.1.1-architecture-stable`

From this tag forward:
- Core architecture is frozen
- Changes require formal governance review
- Backward compatibility is prioritized
- Stability takes precedence over new features

This layer **MUST NOT** depend on:
- External services
- Centralized authorities
- Mutable runtime assumptions

## Network Status

- Nodes operate as system services (systemd)
- Block production verified on-chain
- Multi-node operation validated
- Long-running soak tests in progress (7–14 days, no changes)

The network is currently in a **stability observation phase**.

## Governance & Transparency

- Governance principles are documented separately
- Founder allocation disclosure is public
- Change control is logged and auditable
- Incident response procedures are defined

Quantumpay is developed with a **compliance-first and public-interest mindset**.

## Regulatory Posture

Quantumpay does not claim:
- Legal tender status
- Monetary authority
- Regulatory approval

All future financial, token, or commercial activities will follow:
- Applicable laws and regulations
- Guidance from relevant authorities
- Jurisdiction-specific compliance requirements

This repository represents **infrastructure**, not financial products.

## Disclaimer

This software is provided **as-is**, for research, infrastructure development, and public review.

Nothing in this repository constitutes:
- Investment advice
- Financial products
- Solicitation or offering of securities

## Contribution

See `CONTRIBUTING.md` for contribution guidelines.

## Final Note

QuantumPay is built with the belief that **strong infrastructure precedes adoption**.

Quietly.
Carefully.
Correctly.
