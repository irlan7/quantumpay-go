# QuantumPay (QP) – Staking & Slashing Model

**Version:** v1.0 (Frozen)  
**Network:** QuantumPay Mainnet  
**Chain ID:** 77001  
**Status:** Mainnet Candidate – Economics Locked

## 1. Purpose & Design Goals

The QuantumPay staking and slashing model is designed to:

- Secure the network through economic incentives
- Align validator and delegator behavior with protocol integrity
- Penalize malicious or faulty actions deterministically
- Remain simple, auditable, and resistant to economic exploits

This model is **fixed for mainnet launch** and may only be changed via formal governance proposals.

## 2. Participants

### 2.1 Validators
Validators are entities responsible for:
- Block production
- Transaction validation
- Network consensus participation

Validators must lock (stake) QP tokens to participate.

### 2.2 Delegators
Delegators stake QP tokens by delegating to validators.
They:
- Share rewards proportionally
- Share slashing risk proportionally
- Do not participate directly in block production


## 3. Staking Mechanics

### 3.1 Minimum Stake
- **Validator minimum stake:** Protocol-defined (configurable, non-zero)
- **Delegator minimum stake:** ≥ 1 QP

### 3.2 Stake Lockup
- Staked QP is locked and non-transferable
- Unstaking requires an **unbonding period**

### 3.3 Unbonding Period
- Fixed unbonding delay (e.g., N epochs)
- Stake remains slashable during unbonding
- Prevents instant exit after malicious behavior

## 4. Reward Distribution

### 4.1 Reward Source
Validator rewards are funded from:
- Gas fees (validator share after burn)
- Protocol incentives (if enabled)

### 4.2 Distribution Formula
For each epoch:
