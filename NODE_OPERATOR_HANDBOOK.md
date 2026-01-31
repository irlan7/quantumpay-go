# Quantumpay Node Operator Handbook

This handbook provides practical guidance for running a Quantumpay node.

It is intended for independent node operators, infrastructure providers,
and organizations participating in the Quantumpay network.

Operating a node is a responsibility.  
This handbook prioritizes safety, predictability, and long-term network health.

## 1. What Is a Quantumpay Node

A Quantumpay node:

- Participates in peer-to-peer networking
- Validates and propagates blocks and transactions
- Maintains a copy of the blockchain state
- Enforces protocol rules locally

A node does **not**:
- Control the network
- Have special privileges
- Receive guaranteed rewards
- Act as a custodian for users

## 2. Operator Responsibilities

Node operators are responsible for:

- Maintaining reliable infrastructure
- Monitoring node health
- Applying upgrades intentionally
- Securing their operating environment
- Following governance and operational guidelines

Running a node is voluntary and at the operator’s own discretion.

## 3. Minimum System Requirements

Recommended minimums for a production node:

- 2 CPU cores
- 4 GB RAM
- SSD storage
- Stable internet connection
- Linux-based OS (Ubuntu LTS recommended)

Higher resources improve reliability but are not mandatory.

## 4. Installation Overview

A typical node installation includes:

- Linux system preparation
- Installation of Quantumpay node binary
- Configuration files
- systemd service setup
- Network port configuration

Refer to official documentation or release notes for exact commands.

## 5. Running the Node

### Service Management
Nodes are expected to run under systemd:

- Automatic restart enabled
- No manual background execution
- Controlled start and stop

Manual execution methods (nohup, screen, tmux) are discouraged.

## 6. Monitoring Your Node

Operators should regularly check:

- Service status (`systemctl status`)
- Logs (`journalctl`)
- Disk usage
- Memory usage
- Peer connectivity
- Block height progression

There is no requirement for complex monitoring systems.

## 7. Upgrades

- Upgrades are **never automatic**
- Operators decide when to upgrade
- Read release notes before upgrading
- Stop the node before applying upgrades
- Monitor behavior after restart

Running older versions is allowed unless explicitly deprecated.

## 8. Snapshot and Recovery

Operators are encouraged to:

- Take periodic snapshots of node data
- Store snapshots off-node
- Test recovery procedures occasionally

Snapshot and recovery protect against hardware failure and corruption.

## 9. Security Best Practices

Operators should:

- Keep the OS up to date
- Restrict SSH access
- Use firewall rules
- Separate operator accounts from system services
- Avoid exposing unnecessary ports

Security incidents should follow the process in `SECURITY.md`.

## 10. Incident Handling

If an issue is detected:

1. Observe and collect logs
2. Avoid repeated restarts
3. Do not modify state files
4. Escalate if necessary
5. Follow `INCIDENT_RESPONSE.md`

Improper handling can worsen network impact.

## 11. Governance Awareness

Node operators should be aware that:

- Protocol changes follow formal governance processes
- There are no forced upgrades
- Chain ID and genesis are immutable
- Operators are not expected to make governance decisions

Refer to `GOVERNANCE.md` for details.

## 12. What Not to Do

Operators must not:

- Modify protocol code arbitrarily
- Change genesis files
- Attempt to influence consensus behavior
- Run experimental builds on mainnet
- Publicly speculate during incidents

## 13. Philosophy

Quantumpay nodes are part of shared public infrastructure.

Responsible operation values restraint, discipline, and consistency
over experimentation or optimization for short-term gain.

## Final Note

Running a node contributes to network resilience.

The most valuable operator behavior is quiet reliability.
