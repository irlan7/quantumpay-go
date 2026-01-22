
QuantumPay Blockchain — Mainnet

QuantumPay is a modular, lightweight blockchain core designed for secure value transfer, predictable gas economics, and validator-based staking incentives.
This repository contains the Mainnet Core implementation written in Go.

> Status: ✅ Mainnet Core Stable
Tag: v1.0.0-core
Language: Go
License: MIT

✨ Features

Deterministic Coin & Gas Model

Amount based on math/big.Int

Safe mint / burn / transfer invariants

Predictable gas accounting


Staking & Rewards

Validator staking

Reward distribution

Slashing-ready architecture


Modular Architecture

Clean separation: coin, staking, state, consensus

Disabled experimental modules safely archived

No import cycles


Production-Ready Core

All core packages build successfully

Unit-tested economic invariants

CLI binary ready


📦 Repository Structure

quantumpay-go/
├── cmd/
│   └── quantumpay/          # CLI entrypoint
│
├── internal/
│   ├── coin/                # Coin, gas, supply, balance logic
│   ├── staking/             # Staking, rewards, slashing
│   ├── blockchain/          # Block & chain logic
│   ├── consensus/           # Consensus abstraction
│   ├── state/               # State management
│   ├── crypto/              # Cryptography primitives
│   ├── grpc_disabled/       # Archived (non-mainnet)
│   └── p2pv1_disabled/      # Archived experimental P2P
│
├── .gitignore
├── go.mod
├── go.sum
└── README.md


🚀 Build & Test

Build CLI

go build ./cmd/quantumpay

Run Tests

go test ./internal/coin
go test ./internal/staking

All mainnet-critical modules pass build.


🧪 Economic Safety

The following invariants are enforced and tested:

Total supply consistency

No negative balances

Gas deducted before execution

Rewards never exceed minted supply

Slashing cannot underflow stake


Economic logic lives in:

internal/coin/

internal/staking/


🔒 Mainnet Policy

Experimental modules are disabled, not deleted

Core APIs are stable (no breaking changes)

Economic logic is frozen for v1

Future upgrades require explicit versioning


🏷 Versioning

v1.0.0-core — Mainnet Core (current)

Future protocol upgrades will follow semantic versioning

📖 Documentation

ECONOMICS.md — Coin & gas model

GAS_MODEL.md — Gas accounting

STAKING_MODEL.md — Validator economics


🤝 Contribution

Mainnet core is frozen.
Development continues on feature branches only.

git checkout -b feature/<name>


🕌 Acknowledgment

Built with discipline, audit-first mindset, and responsibility.
May this technology bring benefit, fairness, and trust.


📜 License

MIT License
© QuantumPay Contributors

