
Quantumpay — Mainnet Ops Playbook (1 Page)

Scope: Operational readiness for Quantumpay Mainnet
Excludes: Core protocol logic, consensus changes, tokenomics
Goal: Stable, secure, compliant mainnet operations


1. Operational Principles (Non-Negotiable)

1. Stability > Features
No feature additions during mainnet phase without governance approval.


2. Determinism First
Ops changes must not affect block determinism or state transitions.


3. Recoverability is Mandatory
Every failure must have a documented recovery path.


4. Least Privilege
Node, ops, and governance access strictly separated.


2. Node Operations Standard

Node Requirements

Dedicated VPS (no shared workloads)

systemd-managed service (auto-restart)

Time sync enabled (chrony / ntpd)

Disk IO monitored


Mandatory Checks (Daily)

systemctl status quantumpay-node

Block height progressing

Disk usage < 70%

Memory stable (no leaks)


Uptime Target

≥ 99.5% per node (rolling 30 days)


3. Snapshot & Recovery (Critical)

Snapshot Policy

Frequency:

Full snapshot: weekly

State snapshot: daily


Storage:

Off-node (separate VPS / object storage)


Retention:

Minimum 14 days



Recovery SLA

Node restore ≤ 2 hours

Full network recovery ≤ 24 hours


> No snapshot = mainnet NOT ready.


4. Monitoring & Alerting (Minimal but Sufficient)

Required Metrics

Block height

Block interval

Peer count

Disk usage

Memory usage


Alert Rules

Block stalled > 2 block intervals

Disk > 80%

Service restart loop (>3/hour)


Logging via journald is acceptable for phase 1.



5. Governance & Change Control

Change Categories

Type	Examples	Approval

Ops	Config, ports, infra	Core Ops
Governance	Validators, policy	Governance Council
Protocol	Core logic	Frozen


Change Rules

Written proposal required

24–72h review window

Change log mandatory

Rollback plan required


6. Security Operations

Key Management

Validator keys never exposed to internet

No key reuse across environments

Offline backup required


Incident Response

1. Freeze affected nodes


2. Preserve logs


3. Snapshot state


4. Governance notification


5. Public transparency report (post-mortem)


7. Compliance & Public Posture

No profit guarantees

No price discussion

No influencer promotion

Transparent technical communication only


Quantumpay positions itself as infrastructure, not speculation.


---

8. Mainnet Readiness Checklist (Ops)

✅ systemd active on all nodes
✅ ≥72h continuous testnet uptime
✅ Snapshot & restore tested
✅ Monitoring active
✅ Governance SOP documented
⏳ Incident drill (recommended before public launch)


9. Mainnet Go / No-Go Criteria

GO if:

Nodes stable ≥ 7 days

Recovery tested

Governance active


NO-GO if:

Snapshot missing

Manual intervention required daily

Undocumented config differences



Final Note

> Quantumpay wins by being boring, stable, and predictable.
Mainnet success is operational discipline, not hype.




