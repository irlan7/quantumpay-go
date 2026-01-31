# Quantumpay Network Ports

This document defines the network ports used by the Quantumpay blockchain node.

The purpose of this document is to provide clarity for node operators,
infrastructure providers, and security audits.

Only the ports listed here are required for normal node operation.

## Design Principles

- Minimal port exposure
- Explicit purpose per port
- No hidden or dynamic port usage
- Firewall-friendly configuration

If a port is not listed here, it should be considered **unused**.

## Required Ports

### P2P Network Port

| Item | Value |
|---|---|
| Purpose | Peer-to-peer communication |
| Protocol | TCP |
| Default Port | `7001` |
| Direction | Inbound & Outbound |
| Required | Yes |

Description:
- Used for block and transaction propagation
- Used for peer discovery and synchronization
- Must be reachable by other nodes

Firewall:
- Allow inbound TCP on this port
- Allow outbound TCP traffic

## Optional / Operator-Specific Ports

### Local RPC / IPC (Optional)

| Item | Value |
|---|---|
| Purpose | Local node control / integration |
| Protocol | UNIX socket or localhost TCP |
| Default | Disabled |
| Required | No |

Description:
- Intended for local tools or internal services
- Must NOT be exposed publicly
- Should be bound to `127.0.0.1` or UNIX socket only

## Ports Not Used

Quantumpay nodes **do not use**:

- UDP ports
- Dynamic port ranges
- Public HTTP ports
- WebSocket endpoints exposed to the internet

Any service exposing additional ports is **outside the core protocol**.

## Firewall Recommendations

Minimum recommended firewall rules:

- Allow inbound TCP on P2P port
- Allow outbound TCP traffic
- Deny all other inbound connections by default

Example (conceptual):
- `ALLOW tcp/7001`
- `ALLOW outbound`
- `DENY all inbound`

Exact commands depend on the operating system and firewall software.

## Multi-Node Deployment

When running multiple nodes on the same host:

- Each node must use a unique P2P port
- Ports must be explicitly configured
- Port conflicts must be avoided

## Change Policy

- Network port changes are considered **protocol-affecting**
- Any modification requires governance approval
- Changes must be documented and announced

## Security Notes

- Never expose local control interfaces publicly
- Restrict access using firewall rules
- Monitor for unexpected inbound traffic

## Summary

Quantumpay network communication is intentionally simple and minimal.

Reducing exposed ports reduces operational complexity and attack surface.
