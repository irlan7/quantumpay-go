
QuantumPay is a high-performance, sovereign Layer-1 blockchain infrastructure built for national digital sovereignty and global scalability. Engineered in Go-Lang, it achieves near-instant finality through a customized BFT-based consensus mechanism while maintaining an ultra-light resource footprint.
🌐 Network Identity (SSoT)
As the Single Source of Truth, these parameters define the current state of the frozen mainnet-alpha:
| Parameter | Value |
|-------------------|----------------|
| Chain ID          | 77077 [FROZEN] |
| Genesis Hash      | 0x1d58599424f1159828236111f1f9e83063f66345091a99540c4989679269491a |
| Core Engine       | Go-Lang (quantumpay-go-v1.1) |
| Official Provider | VPS-9451c332 |
⚡ Technical Features
 * High Efficiency: Optimized to consume as little as 5% RAM on standard hardware, ensuring maximum decentralization.
 * Fast Consensus: Block production and finality achieved in under 5 seconds.
 * Production Process Management: Utilizes PM2 for 24/7 bridge and node uptime with minimal memory overhead (~24.7MB).
🛠 Run a Node (Join the Decentralization)
Follow these steps to synchronize your validator with the QuantumPay Network:
Hardware Requirements
 * CPU: 2 Cores (Minimum)
 * RAM: 4GB (Optimized usage: ~5%)
 * Storage: 40GB SSD
 * OS: Ubuntu 22.04 LTS / 24.04 LTS
Installation
# Clone the core repository
git clone https://github.com/irlan7/quantumpay-go
cd quantumpay-go

# Build the node
go build -o quantumpay-node ./cmd/node

# Start the node (Recommended with PM2)
pm2 start ./quantumpay-node --name "qp-node"

#Genesis Verification ,Genesis Hash via logs
pm2 logs qp-node --lines 100 | grep "Genesis Hash"

Verification Check:
​Expected Hash: 0x1d58599424f1159828236111f1f9e83063f66345091a99540c4989679269491a

read Install.sh 

📡 Official Channels
 * Website: https://quantumpaychain.org
 * Explorer: /explorer
 * X (Twitter): @quantumpaychain
 * Email: quantumpaysec@gmail.com
📜 License & Vision
This project is open-source under the MIT License. We follow the vision of Satoshi Nakamoto and Vitalik Buterin: building a trustless, transparent, and permissionless world where every individual can run a node and verify the state of the truth.
Would you like me to create a bash script (install.sh) that automates this entire process for new users who want to join your network?
=======

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

