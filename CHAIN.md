
🔒 Chain ID Freeze (Final)

Chain ID

Name: QuantumPay Mainnet

Chain ID: 77077

Status: ✅ FINAL & FROZEN

Effective Date: 16 January 2026


Policy Statement

> The Chain ID 77077 is permanently frozen and MUST NOT be changed under any circumstances.


Changing the Chain ID after this point would:

Break wallet compatibility

Invalidate signed transactions

Risk chain replay attacks

Fragment the network

Violate mainnet immutability guarantees

Therefore, Chain ID 77077 is final for the lifetime of the QuantumPay mainnet.


Scope of Applicability

The frozen Chain ID applies to:

Core blockchain protocol

Validator nodes

RPC endpoints

Wallet integrations

Explorer services

Smart contract deployment

Cross-service and cross-node communication


All environments referring to QuantumPay Mainnet must use:

chain_id = 77077


Governance & Upgrade Rules

No governance proposal may alter the Chain ID

No hard fork may reuse or change this Chain ID

Any future network requiring a different Chain ID must be launched as a new chain, not a replacement


Security Rationale

Freezing the Chain ID ensures:

Replay-attack prevention

Deterministic transaction signing

Long-term ecosystem stability

Compatibility with external tools (wallets, explorers, exchanges)

Clear separation between testnet, devnet, and mainnet


Audit Note

This Chain ID freeze is part of the Mainnet Readiness Checklist and is considered irreversible.

Auditors and integrators may rely on the following invariant:

> QuantumPay Mainnet Chain ID will always be 77077.


Change Log

Date	Change

2026-01-16	Chain ID 77077 frozen (final)

