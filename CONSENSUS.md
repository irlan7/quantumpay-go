# Quantumpay Consensus Protocol

**Document:** CONSENSUS.md  
**Network:** Quantumpay Mainnet  
**Chain ID:** 77077  
**Consensus Type:** Proof-of-Stake (PoS) with BFT Finality  
**Status:** Mainnet (Active)  
**Last Updated:** 2026-01-25

## 1. Overview

Quantumpay uses a **Proof-of-Stake (PoS)** consensus mechanism combined with **Byzantine Fault Tolerant (BFT) finality** to achieve secure, energy-efficient, and deterministic block confirmation.

This design prioritizes:
- Safety and consistency
- Fast finality
- Economic accountability
- Long-term sustainability
- Post-quantum readiness


## 2. Consensus Model

Quantumpay consensus consists of two integrated layers:

1. **Proof-of-Stake (PoS)**  
   Determines validator eligibility and economic security.

2. **BFT Finality Engine**  
   Ensures blocks are finalized deterministically once consensus is reached.

This hybrid approach combines the economic incentives of PoS with the strong safety guarantees of BFT.


## 3. Validators

Validators are network participants who:
- Stake native QP coins
- Propose and validate blocks
- Participate in BFT voting rounds

### Validator Properties
- **Stake-based weight**
- **Bounded validator set**
- **Deterministic leader selection**
- **On-chain accountability**

The maximum validator count is capped to preserve BFT safety and network performance.

## 4. Block Production

- Block producers are selected based on stake-weighted rotation.
- Each block proposal enters a BFT consensus round.
- A block is finalized when it receives sufficient validator signatures (≥2/3 by stake).

Once finalized, a block **cannot be reverted**.

## 5. Finality Guarantees

Quantumpay provides **deterministic finality**, meaning:
- No probabilistic confirmations
- No chain reorganizations after finality
- Immediate settlement assurance

This makes Quantumpay suitable for:
- Payment systems
- Financial settlement
- Machine-to-machine transactions
- Regulatory-compliant infrastructure

## 6. Byzantine Fault Tolerance

The consensus engine tolerates:
- Up to **1/3 malicious or faulty validators**
- Network partitions
- Message delays

Safety is preserved as long as ≥2/3 of the validator stake behaves honestly.

---

## 7. Slashing & Accountability

Validators may be penalized (slashed) for:
- Double-signing
- Equivocation
- Extended downtime
- Protocol violations

Slashing is enforced on-chain and affects staked QP balances.


## 8. Energy Efficiency

Unlike Proof-of-Work systems, Quantumpay:
- Does not require mining hardware
- Consumes minimal energy
- Scales without increasing environmental cost

This aligns with global sustainability standards.


## 9. Post-Quantum Readiness

The consensus design is **crypto-agile**:
- Signature algorithms are abstracted
- Supports future migration to post-quantum cryptography
- No dependence on hash-based mining assumptions


## 10. Governance & Evolution

Changes to consensus parameters require:
- On-chain governance process
- Validator approval
- Public change documentation

Backward compatibility and network stability are prioritized.

---

## 11. Comparison Summary

| Feature                 | Quantumpay (PoS+BFT)    | Proof-of-Work      |
|-------------------------|-------------------------|--------------------|
| Finality                | Deterministic           | Probabilistic     |
| Energy Use              | Low                     | High              |
| Reorg Risk              | None after finality     | Always possible   |
| Settlement Speed        | Fast                    | Slow              |
| Economic Accountability | Explicit                | Implicit          |


## 12. Scope Notice

This mainnet release focuses on **core consensus stability**.  
Public APIs, explorer services, and extended tooling will be introduced in subsequent phases.


## 13. Disclaimer

Quantumpay is an open protocol.  
Participation as a validator involves economic risk and operational responsibility.

This document describes protocol mechanics and does not constitute financial advice.
