
🚀 QuantumPay Mainnet — Release v1.0.0-core

🔖 Tag

v1.0.0-core

📅 Release Date

2026-01-22 (UTC)

🧩 Overview

This release marks the first stable version of the QuantumPay Mainnet Core, a blockchain protocol designed and implemented in Go to ensure deterministic economic logic, secure validator participation, and auditable consensus.

> 🏁 Mainnet Core is live-ready.
Economic model, coin invariants, and staking logic have all passed deterministic test validation.


⚙️ Highlights

💰 Economic Layer

Core QP coin (internal/coin/)

Safe mint/burn/transfer invariants

Supply cap defined via config/genesis.json

Gas accounting finalized (internal/coin/gas.go)


🛡️ Staking Layer

Validator reward distribution

Slashing framework

Integration-tested with coin supply


🧱 Architecture

Modular & import-cycle-free

Deterministic ledger and state layer

CLI entrypoint: cmd/quantumpay/main.go


🧾 Genesis

Chain ID: 77001

Network: quantumpay-mainnet

Supply: 210,000,000 QP

Genesis Hash: acc395137e5d0c28c609d011ea99d89405f07009c0bbf8933711e1a7f184edc6


🔒 Security & Policy

All deprecated modules moved to _disabled directories

API and core logic frozen

No breaking changes allowed post-release

Governance & updates through QuantumPay Foundation


🧪 Verified Tests

Module	Status	Description

internal/coin	✅ Passed	Core monetary invariants
internal/staking	✅ Passed	Reward / penalty logic
cmd/quantumpay	✅ Build OK	CLI transactions


🧰 Build Instructions

git clone https://github.com/irlan7/quantumpay-go
cd quantumpay-go
go mod tidy
go build ./cmd/quantumpay


📜 Documentation

ECONOMICS.md

GAS_MODEL.md

STAKING_MODEL.md

MAINNET_DECLARATION_v1.1.md


🕊️ Acknowledgment

This milestone marks the completion of Indonesia’s first L1 blockchain developed from zero — designed to be auditable, regulator-safe, and globally scalable.

> “With knowledge, discipline, and faith — technology becomes a blessing, not speculation.”
