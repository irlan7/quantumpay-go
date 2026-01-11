# Snapshot & Recovery SOP  
Quantumpay Protocol

**Scope:** Node Data Protection & Recovery  
**Phase:** P1 (Frozen)  
**Audience:** Operators, Maintainers  
**Status:** Active

## 1. Purpose

This SOP defines a **safe, deterministic procedure** for creating snapshots and recovering a Quantumpay node.

Goals:
- Protect node data
- Enable recovery after failure
- Prevent state corruption
- Preserve network determinism

Snapshots are **operational safeguards**, not protocol features.

## 2. Core Principles

- **Cold snapshot only** (node must be stopped)
- **Manual recovery only** (no automation in P1)
- **No protocol or state modification**
- **Safety over speed**

---

## 3. Snapshot Procedure (Cold Snapshot)

Step 1 — Stop Node
bash
sudo systemctl stop quantumpay-node

Confirm:

systemctl status quantumpay-node

Step 2 — Snapshot Data Directory

Default data path:

/opt/quantumpay/data

Create snapshot archive:

cd /opt/quantumpay
sudo tar -czf quantumpay-snapshot-YYYYMMDD.tar.gz data/

Alternative (disk backup):

sudo rsync -a data/ /backup/quantumpay-data/


Step 3 — Restart Node

sudo systemctl start quantumpay-node

Verify logs:

journalctl -u quantumpay-node -n 20 --no-pager

4. Snapshot Storage Rules

Store snapshots off-node when possible

Label clearly (date, node ID)

Retain:

Latest snapshot

One previous snapshot


Protect from accidental overwrite or deletion


5. Recovery Procedure

Step 1 — Stop Node

sudo systemctl stop quantumpay-node


Step 2 — Preserve Existing Data

sudo mv /opt/quantumpay/data /opt/quantumpay/data.backup.YYYYMMDD


Step 3 — Restore Snapshot

sudo tar -xzf quantumpay-snapshot-YYYYMMDD.tar.gz -C /opt/quantumpay/
sudo chown -R quantumpay:quantumpay /opt/quantumpay/data

Step 4 — Restart Node

sudo systemctl start quantumpay-node

Step 5 — Verify Recovery

Confirm:

Node starts cleanly

Block height increases

No panic or corruption logs

Peer connections restored


If verification fails, stop immediately and escalate.

6. When NOT to Restore

Do NOT restore if:

Snapshot source is untrusted

Snapshot version mismatches protocol

Genesis or Chain ID differs

Governance freeze would be violated

7. Incident Escalation

If recovery fails:

1. Stop node

2. Preserve all data

3. Notify maintainer

4. Follow INCIDENT_RESPONSE.md

5. Record in CHANGE_CONTROL_LOG.md

   
8. Governance & Change Control

Snapshot & recovery must never alter protocol behavior

No automated restore allowed during P1

SOP changes require governance approval


9. Operational Philosophy

> Recovery must be boring, predictable, and reversible.

Fast recovery is useful.
Correct recovery is mandatory.


Owner: Quantumpay Operations
Review Cycle: Pre-Mainnet / As Needed

