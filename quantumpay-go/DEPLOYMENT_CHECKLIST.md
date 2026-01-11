# Quantumpay Deployment Checklist

This checklist defines the minimum steps required to deploy and operate
a Quantumpay node safely in a production environment.

It is intended to be used **before**, **during**, and **after** deployment.

Do not skip steps.

## 1. Pre-Deployment Checklist

### Governance & Protocol
- [ ] Governance documents finalized and published
- [ ] Chain ID confirmed and immutable
- [ ] Genesis configuration finalized
- [ ] Genesis hash verified
- [ ] No pending protocol changes


### Infrastructure Preparation
- [ ] Supported OS installed (Ubuntu LTS recommended)
- [ ] System fully updated
- [ ] Correct timezone and NTP enabled
- [ ] Dedicated node user created
- [ ] SSH key-based access configured
- [ ] Root login disabled


### Security Baseline
- [ ] Firewall configured (default deny inbound)
- [ ] Only required ports opened (`NETWORK_PORTS.md`)
- [ ] Password authentication disabled
- [ ] Node runs as non-root user
- [ ] No unnecessary services running


## 2. Installation Checklist

### Binary & Files
- [ ] Official release binary obtained
- [ ] Binary checksum verified (if provided)
- [ ] Binary placed in correct directory
- [ ] File permissions verified


### Configuration
- [ ] Genesis file placed correctly
- [ ] Chain ID verified in config
- [ ] Data directory initialized
- [ ] P2P port configured and unique
- [ ] No testnet/devnet flags enabled


### Service Management
- [ ] systemd service file installed
- [ ] Service runs under correct user
- [ ] Auto-restart enabled
- [ ] Service enabled at boot


## 3. Initial Startup Checklist

- [ ] Node starts without errors
- [ ] Logs show normal startup sequence
- [ ] P2P port reachable
- [ ] Peer connections established
- [ ] Block height progresses

Do not proceed if errors are observed.


## 4. Post-Deployment Verification

### Network Health
- [ ] Block height consistent with peers
- [ ] No repeated crashes or restarts
- [ ] Peer count stable
- [ ] Resource usage within limits

---

### Logging & Monitoring
- [ ] Logs flowing to journald
- [ ] Disk usage monitored
- [ ] Memory usage monitored
- [ ] Alerts or manual checks defined
      

## 5. Snapshot & Backup

- [ ] Initial snapshot taken
- [ ] Snapshot stored off-node
- [ ] Snapshot restore tested (at least once)
- [ ] Snapshot schedule defined


## 6. Upgrade Readiness

- [ ] Upgrade procedure understood
- [ ] Rollback plan documented
- [ ] Previous binary retained
- [ ] Release notes reviewed


## 7. Incident Readiness

- [ ] Incident Response document reviewed
- [ ] Logs access confirmed
- [ ] Escalation path known
- [ ] Emergency freeze procedure understood


## 8. Final Go/No-Go Check

### GO if:
- [ ] Node stable for ≥24 hours
- [ ] No critical errors observed
- [ ] Governance freeze active
- [ ] Snapshot verified

### NO-GO if:
- [ ] State inconsistency detected
- [ ] Repeated crashes occur
- [ ] Genesis mismatch found
- [ ] Unclear operational status

## Final Note

A successful deployment is quiet and uneventful.

If something feels rushed, stop and reassess.
