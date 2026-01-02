# QuantumPay Go Core

Status: 🔒 ARCHITECTURE FROZEN

This repository contains the **Go Core Blockchain Engine** for QuantumPay.

## Scope
- Pure blockchain engine
- Deterministic state machine
- No networking consensus
- No node authority

## Architecture
- core: immutable data + hashing
- state: world state & transitions
- block: block building
- blockchain: chain storage & views
- engine: orchestrator

## Freeze Tag
v0.1.1-architecture-stable

⚠️ This layer MUST NOT depend on:
- P2P
- RPC
- Node.js
- Consensus
