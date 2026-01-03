# QuantumPay P2P v1

## Goals
- Simple
- Deterministic
- No consensus logic
- No state mutation

## Rules
- P2P MUST NOT import engine, blockchain, or state
- P2P only transports bytes
- Validation happens in higher layer
- One TCP connection = one peer

## Forbidden
- ❌ No mempool access
- ❌ No block execution
- ❌ No global state
