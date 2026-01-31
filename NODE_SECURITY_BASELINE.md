# Quantumpay Node Security Baseline

This document defines the minimum security baseline required to operate
a Quantumpay node in production environments.

The objective is to reduce attack surface, prevent common misconfigurations,
and ensure predictable node behavior.

This is a baseline, not a guarantee.

## 1. Security Philosophy

- Minimize exposed surfaces
- Prefer simplicity over complexity
- Assume the internet is hostile
- Assume misconfiguration is the biggest risk

Security is achieved through discipline, not tooling alone.

## 2. Operating System

### Required
- Linux (Ubuntu LTS recommended)
- Regular security updates applied
- Supported kernel version

### Not Allowed
- End-of-life OS versions
- Unmaintained custom kernels
- Desktop environments on production nodes

## 3. User and Privilege Model

- Dedicated system user for Quantumpay node
- Node process must not run as root
- No shared system accounts
- Principle of least privilege enforced

Example:
- `quantumpay` user
- `nologin` shell
- Limited filesystem access

## 4. SSH Access

Minimum SSH requirements:

- Key-based authentication only
- Password authentication disabled
- Root login disabled
- Limited number of authorized keys
- Strong key types (ed25519 recommended)

SSH access should be limited to known operators.

## 5. Network Security

- Only required ports exposed (see `NETWORK_PORTS.md`)
- Default-deny inbound firewall policy
- Outbound traffic allowed
- No public RPC or control interfaces

Unexpected open ports indicate misconfiguration.

## 6. Filesystem and Storage

- Dedicated data directory
- Correct ownership and permissions
- No world-writable directories
- Sufficient disk space monitoring

Do not use network-mounted storage for production nodes.

## 7. Process Management

- Node must run under systemd
- Automatic restart enabled
- No manual background execution
- Resource limits applied where appropriate

Systemd provides predictable lifecycle management.

## 8. Logging and Monitoring

- Logs must be retained
- Logs must not be modified during incidents
- Disk usage for logs monitored
- Access to logs restricted

Logs are security artifacts.

## 9. Software Integrity

- Use official release binaries or verified builds
- Verify checksums where provided
- Avoid custom or patched binaries on mainnet
- Track deployed version explicitly

Unknown binaries are a security risk.

## 10. Snapshot and Backup Security

- Snapshots stored off-node
- Snapshot access restricted
- Snapshot integrity verified periodically

Snapshots contain sensitive state data.

## 11. Incident Preparedness

Operators must be familiar with:
- `INCIDENT_RESPONSE.md`
- Snapshot and recovery procedures
- Escalation paths

Preparation reduces impact more than reaction.

## 12. What Not to Do

Operators must not:

- Expose admin interfaces publicly
- Modify state data manually
- Run experimental builds on mainnet
- Share credentials
- Ignore security updates indefinitely

## 13. Audit and Compliance

Operators should periodically review:
- Firewall rules
- User access
- Disk usage
- Node behavior

Security baselines should be revalidated after upgrades.

## Summary

This security baseline defines the minimum acceptable standard for running
a Quantumpay node.

Operators who cannot meet this baseline should not operate production nodes.

Security failures at the node level can affect the entire network.
