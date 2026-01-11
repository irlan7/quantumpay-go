# Governance

This document defines the governance model of the Quantumpay protocol.

Quantumpay is an infrastructure-grade blockchain network designed for long-term
stability, determinism, and regulatory-aware operation. Governance exists to
protect the protocol, not to optimize for speed or popularity.

## Governance Model

Quantumpay follows a **maintainer-driven governance model**.

- The protocol is open source and publicly auditable
- Anyone may review, discuss, or propose changes
- Final decision-making authority rests with the maintainers

This model is intentional and appropriate for an early-stage mainnet protocol.

## Scope of Governance

Governance applies to protocol-level behavior, including:

- Consensus rules
- State transition logic
- Transaction formats and validation
- Network protocol behavior
- Security-critical parameters

Governance does **not** apply to:
- Off-chain services
- Commercial offerings
- Managed infrastructure
- User interfaces, dashboards, or tooling

## Decision Principles

All governance decisions are guided by the following principles:

1. **Stability over velocity**
2. **Determinism over flexibility**
3. **Safety over features**
4. **Explicit changes over implicit behavior**
5. **Backward compatibility where possible**

Changes that increase ambiguity, complexity, or attack surface are rejected by default.

## Change Classification

### Non-Breaking Changes
Examples:
- Internal refactoring with no behavior change
- Logging or observability improvements
- Documentation updates

Approval: Maintainer  
Network notice: Optional

### Soft Changes (Backward Compatible)
Examples:
- Parameter tuning
- Performance improvements
- Optional features

Approval:
- Lead Maintainer
- Governance reviewer

Process:
- Proposal
- Review
- Scheduled release

---

### Hard Changes (Consensus / State Affecting)
Examples:
- Consensus logic changes
- State transition rule changes
- Fee model changes
- Fork-related changes

Approval:
- Governance Committee (quorum ≥ ⅔)

Process:
1. Formal proposal (QCP)
2. Public review window (minimum 14 days)
3. Final approval
4. Scheduled activation (block height or timestamp)

Emergency bypass for hard changes is **not permitted**.

## Chain ID Immutability

The Quantumpay Mainnet Chain ID is permanently fixed.

Once the Quantumpay Mainnet is live, the Chain ID **MUST NOT** be changed, reused,
overridden, or reassigned under any circumstances.

Any future protocol evolution, upgrade, or network change must occur through
explicit versioning and governance processes, not by modifying the Chain ID.

Changing the Chain ID after mainnet launch is considered a protocol violation.

---

## Authority and Control

- There are no hidden administrator keys
- There is no emergency kill switch
- Maintainers cannot arbitrarily modify ledger state

Protocol rules are enforced by node software and consensus, not by individuals.

## Change Freeze Policy

### Architecture Freeze
- Core architecture is considered **frozen** after mainnet launch
- Changes require formal governance approval

### Emergency Freeze
Activated if:
- Network instability
- Security incidents
- Chain divergence

During emergency freeze:
- All changes are suspended except security mitigation

## Transparency and Auditability

- All approved changes are documented
- Decisions are recorded with timestamps and approvers
- Releases reference corresponding governance decisions
- No undocumented or silent changes are allowed


## Evolution of Governance

Governance may evolve as the network matures.

Any evolution will be:
- Publicly documented
- Gradual
- Backward-compatible
- Focused on preserving protocol neutrality

Governance evolution will not compromise protocol safety.

## Non-Goals

Quantumpay governance explicitly avoids:

- Token-weighted voting
- Popularity-based governance
- Rapid or speculative protocol changes
- Governance mechanisms that increase risk

## Summary

Quantumpay governance is intentionally conservative.

The objective is to maintain a stable, neutral, and predictable protocol suitable
for long-term public and institutional use, rather than a rapidly changing platform.
