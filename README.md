# QuantumPay Go Core

**Status:** 🔒 Architecture Frozen  
**Role:** Core Blockchain Engine (Deterministic Layer)

QuantumPay Go Core is the **foundational blockchain engine** of the QuantumPay network.

This repository intentionally contains **no UI, wallet, SDK, or application-layer logic**.  
Its sole purpose is to provide a **deterministic, minimal, and stable blockchain core** that higher-level protocols can safely build upon.

This repository is designed for **long-term production use** and **protocol stability**.

---

## Purpose

QuantumPay Go Core serves as the **technical foundation** for:

- QuantumPay Network
- QuantumSwap
- QuantumDEX
- Future protocol extensions

Once stabilized, this core is expected to change **very rarely**.

---

## Scope (What This Repo Is — and Is Not)

### Included
- Pure blockchain engine
- Deterministic state machine
- Block production & validation
- World state & state transitions
- Chain storage & read views
- Engine orchestration loop

### Explicitly NOT Included
- Wallets
- UI / frontend
- Smart contract VM
- Governance logic
- Consensus authority / validator politics
- Application-layer protocols

Networking exists **only to support block propagation**, not governance or authority.

---

## Architecture Overview

Transactions │ ▼ Deterministic Execution │ ▼ State Transitions ──► World State │ ▼ Block Building ──► Block Hashing │ ▼ Chain Storage ──► Views / Queries

**Key principle:**  
> Same input → same state → same block hash.

---

## Design Principles

### Determinism First
All state transitions are deterministic.  
No randomness, no time-based logic, no non-deterministic behavior.

### Minimal Surface Area
The core does one thing well: **produce and validate blocks**.

Everything else must live **outside** this repository.

### Long-Term Stability
Architecture is frozen to avoid chain-breaking changes.

Breaking changes require **new networks**, not patches.

### Composable by Design
Higher-level protocols are expected to build **on top of** the core, never inside it.

---

## Repository Structure

cmd/ └── node/            # quantumpay-node entrypoint
internal/ ├── core/            # immutable primitives & hashing ├── state/           # world state & transitions ├── blockchain/      # blocks, chain, views ├── engine/          # orchestration loop ├── tx/              # transaction handling └── p2p/             # peer-to-peer networking (transport only)
testnet/              # testnet configs & artifacts

---

## Build

```bash
go mod tidy
go build -o quantumpay-node ./cmd/node

