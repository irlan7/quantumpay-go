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
ValidatorReward = TotalEpochRewards × ValidatorCommission DelegatorRewards = RemainingRewards × DelegatorStakeRatio

### 4.3 Validator Commission
- Set by validator (within protocol bounds)
- Applied before delegator distribution
- Immutable per epoch


## 5. Slashing Model

### 5.1 Slashing Principles
- Deterministic
- Proportional
- Immediate
- Non-reversible

Slashing affects:
- Validator self-stake
- Delegated stake


## 6. Slashable Offenses

| Offense | Description | Severity |
|-------|-------------|----------|
| Downtime | Extended validator inactivity | Low |
| Double Signing | Producing conflicting blocks | Critical |
| Invalid State Transition | Consensus rule violation | Critical |
| Byzantine Behavior | Malicious network actions | Critical |


## 7. Slashing Penalties

### 7.1 Penalty Rates (Indicative)

| Offense | Slashing Rate |
|-------|---------------|
| Downtime | 1–5% |
| Double Signing | ≥20% |
| Byzantine Behavior | Up to 100% |

> Exact values are defined in protocol parameters and frozen at genesis.

### 7.2 Slashing Distribution
- Slashed QP is **burned**
- No redistribution to other validators
- Ensures deflationary pressure and deterrence


## 8. Delegator Risk Model

- Delegators inherit validator slashing risk
- Slashing is applied proportionally to delegated stake
- Encourages informed validator selection


## 9. Security & Economic Invariants

The following invariants **must always hold**:

1. Total staked QP ≤ total circulating supply
2. Slashing never increases total supply
3. Rewards never exceed protocol-defined limits
4. No stake can exit without unbonding
5. Slashing applies during unbonding


## 10. Governance & Upgrades

- This staking model is **locked for mainnet**
- Any modification requires:
  - Governance proposal
  - Network-wide approval
  - Explicit version bump


## 11. Summary

The QuantumPay staking and slashing model provides:
- Strong economic security
- Clear incentives
- Predictable penalties
- Long-term network sustainability

This model balances decentralization, safety, and simplicity — suitable for production-grade blockchain infrastructure.


**Document Status:** ✅ Final  
**Audit Ready:** ✅ Yes  
**Mainnet Ready:** ✅ Yes

