# QuantumPay (QP) – Economic Model

**Version:** v1.0 (Locked)  
**Network:** QuantumPay Mainnet  
**Chain ID:** 77001  
**Status:** Mainnet Candidate – Economics Frozen

## 1. Overview

QuantumPay (QP) is the **native coin** of the QuantumPay blockchain, designed as:
- **Medium of exchange** for transactions
- **Gas fee** for network usage
- **Staking asset** for Proof-of-Stake security
- **Economic incentive** for validators and ecosystem participants

  ## Economic Model Diagram

![QuantumPay Tokenomics]( https://github.com/irlan7/quantumpay-go/blob/master/docs/assets/economics-diagram.jpeg)

The QP economic model is designed with:
- Predictable supply
- No token sale
- Long-term sustainability
- Validator-aligned incentives


## 2. Token Specification

| Parameter        | Value            |
|------------------|------------------|
| Symbol           | QP               |
| Decimals         | 18               |
| Total Supply     | 210,000,000 QP   |
| Inflation        | No (fixed supply)|
| Token Sale       | None             |

> **Note:** QP was not sold via ICO, IEO, IDO, or public sale.


## 3. Genesis Supply & Distribution

Total Genesis Supply: **210,000,000 QP**

### Initial Allocation

| Category              | Allocation | Amount (QP) |
|-----------------------|------------|-------------|
| Foundation            | 10%        | 21,000,000  |
| Validator Bootstrap   | 10%        | 21,000,000  |
| Ecosystem Reserve     | 80%        | 168,000,000 |

### Allocation Principles
- **Foundation**: protocol development, legal, security, and operations
- **Validator Bootstrap**: initial network security and decentralization
- **Ecosystem Reserve**: grants, partnerships, future incentives

All allocations are transparently defined in `genesis.json`.


## 4. Utility of QP

QP is required for:

### 4.1 Transaction Fees (Gas)
- All transactions consume gas paid in QP
- Gas pricing is deterministic and consensus-enforced
- Fees are burned or distributed per protocol rules (v1: distributed)

### 4.2 Staking
- Validators must stake QP to participate in consensus
- Delegators may stake QP to validators (optional in v1)
- Stake weight influences validator selection

### 4.3 Network Security
- QP aligns economic incentives with honest behavior
- Slashing penalties apply for protocol violations


## 5. Validator Rewards

Validator rewards originate from:
- Transaction fees
- Protocol-defined reward pool (no inflation)

Reward distribution:
- Proportional to validator stake and performance
- Enforced by on-chain logic
- Transparent and auditable


## 6. Slashing & Penalties

Validators may be penalized for:
- Double signing
- Extended downtime
- Malicious behavior

Penalties include:
- Partial stake slashing
- Temporary jailing
- Permanent removal (severe cases)


## 7. No Inflation Policy

QuantumPay uses a **fixed supply model**:
- No block inflation
- No hidden minting
- Supply changes only via governance upgrade (future)

This ensures:
- Predictable monetary policy
- Protection against dilution
- Long-term value stability


## 8. Governance (v1 Scope)

- No on-chain governance in v1
- Economic parameters are protocol-defined
- Future governance will require explicit mainnet upgrade


## 9. Transparency & Auditability

- Genesis hash publicly published
- `genesis.json` available on GitHub
- Economic logic implemented in open-source Go code
- Validator activity observable on-chain


## 10. Regulatory Positioning

- No token sale
- No promise of profit
- QP functions as a **network utility coin**
- Designed to comply with global blockchain best practices


## 11. Economics Status

| Component            | Status  |
|----------------------|---------|
| Supply Model         | Locked  |
| Genesis Allocation   | Locked  |
| Validator Rewards    | Active  |
| Inflation            | Disabled|
| Token Sale           | None    |


## 12. Disclaimer

QuantumPay (QP) is a **utility coin** used to operate the QuantumPay blockchain network.  
It is not an investment product, security, or financial instrument.


**End of ECONOMICS.md**
