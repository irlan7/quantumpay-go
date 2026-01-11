# Quantumpay Mainnet Launch Checklist

This checklist defines the final conditions required to declare
the Quantumpay Mainnet live.

This document is executed **once**.
If any critical item fails, the launch must be postponed.

## 1. Governance & Declaration (MANDATORY)

- [ ] GOVERNANCE.md finalized and published
- [ ] GOVERNANCE_CHANGE_CONTROL_SOP.md finalized
- [ ] MAINNET_DECLARATION.md reviewed and ready
- [ ] GENESIS_FREEZE.md prepared
- [ ] ARCHIVE.md updated
- [ ] No pending governance proposals
- [ ] Governance freeze active

❌ Launch is blocked if governance is incomplete.


## 2. Chain Identity (MANDATORY)

- [ ] Chain Name confirmed
- [ ] Chain ID confirmed: **77001**
- [ ] Chain ID documented everywhere
- [ ] No testnet or devnet flags enabled
- [ ] Genesis configuration final

❌ Chain identity must never change after launch.


## 3. Genesis & Ledger Integrity (MANDATORY)

- [ ] Genesis block generated
- [ ] Genesis hash verified
- [ ] Genesis files archived
- [ ] All nodes using identical genesis
- [ ] No manual state modification

❌ Genesis mismatch = immediate NO-GO.


## 4. Network Readiness (MANDATORY)

- [ ] Minimum 2 independent nodes running
- [ ] Nodes connected via P2P
- [ ] Block production active
- [ ] Block height consistent across nodes
- [ ] No chain stalls observed (≥72h recommended)

❌ Unstable network = NO-GO.


## 5. Node Operations (MANDATORY)

- [ ] Nodes managed via systemd
- [ ] Auto-restart enabled
- [ ] Logs flowing to journald
- [ ] Disk space sufficient
- [ ] Clock synchronization verified (NTP)

❌ Manual or fragile operation is unacceptable.


## 6. Security Baseline (MANDATORY)

- [ ] Node runs as non-root user
- [ ] SSH key-only access
- [ ] Root login disabled
- [ ] Firewall configured
- [ ] Only required ports exposed

❌ Security shortcuts block launch.


## 7. Snapshot & Recovery (MANDATORY)

- [ ] Initial snapshot taken
- [ ] Snapshot stored off-node
- [ ] Recovery tested successfully
- [ ] Snapshot procedure documented

❌ No recovery plan = NO-GO.


## 8. Incident Readiness (MANDATORY)

- [ ] OPS_PLAYBOOK.md reviewed
- [ ] INCIDENT_RESPONSE.md reviewed
- [ ] Incident escalation path known
- [ ] No unresolved critical incidents

❌ Unresolved incidents block launch.


## 9. Documentation & Public Readiness

- [ ] README.md accurate and current
- [ ] NODE_OPERATOR_HANDBOOK.md published
- [ ] NETWORK_PORTS.md published
- [ ] SUPPORTED_PLATFORMS.md published
- [ ] AUDIT_READINESS.md published

Documentation must reflect reality.


## 10. Legal & Posture Check

- [ ] No investment claims
- [ ] No profit guarantees
- [ ] No custodial services
- [ ] Protocol positioned as neutral infrastructure

❌ Misrepresentation blocks launch.


## 11. Final GO / NO-GO Decision

### GO if:
- All mandatory items are checked
- Network stable and boring
- No pressure to rush
- Team confident to operate long-term

### NO-GO if:
- Any mandatory item is unchecked
- Any uncertainty exists
- Any shortcut is proposed


## Final Statement

Mainnet launch is not a marketing event.

It is a long-term operational commitment.

If there is doubt, delay.
