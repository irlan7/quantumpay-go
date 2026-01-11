# Governance

This document describes the governance model of the Quantumpay protocol.

Quantumpay is designed as an infrastructure-grade blockchain network.  
Governance prioritizes safety, determinism, and long-term stability over rapid change.

## Governance Model

Quantumpay currently follows a **maintainer-driven governance model**.

- The protocol is open source
- Anyone may review, audit, or propose changes
- Final decision-making authority rests with the maintainers

This model is intentional and aligned with early-stage protocol maturity.

## Scope of Governance

Governance applies to:

- Consensus rules
- State transition logic
- Transaction formats
- Network protocol behavior
- Protocol-level security decisions

Governance does **not** apply to:
- Off-chain services
- Commercial offerings
- Managed infrastructure
- UI, dashboards, or tooling


## Decision Principles

All governance decisions follow these principles:

1. **Stability over velocity**
2. **Determinism over flexibility**
3. **Safety over features**
4. **Explicit changes over implicit behavior**
5. **Backward compatibility where possible**

Changes that introduce ambiguity, excessive complexity, or increased attack surface will not be accepted.

## Change Process

Protocol changes follow this general process:

1. Proposal or discussion is raised (issue or design note)
2. Technical evaluation by maintainers
3. Security and compatibility assessment
4. Implementation and testing
5. Explicit release and versioning

There are no automatic upgrades.  
Node operators are responsible for upgrading their software.

## Authority and Control

- There are no hidden administrator keys
- There is no emergency kill switch
- Maintainers cannot arbitrarily modify ledger state

Protocol rules are enforced by node software and consensus, not by individuals.


## Evolution of Governance

The governance model may evolve as the network matures.

Any transition to broader governance (e.g. foundation stewardship or community processes) will be:
- Publicly documented
- Gradual
- Backward-compatible
- Focused on preserving protocol neutrality

## Non-Goals

Quantumpay governance explicitly avoids:

- Popularity-based voting
- Token-weighted governance
- Rapid or speculative protocol changes
- Governance mechanisms that compromise network safety

## Summary

Quantumpay governance is intentionally conservative.

The goal is to maintain a stable, neutral, and predictable protocol that can serve as long-term infrastructure rather than a rapidly changing platform.
