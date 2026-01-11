# Network FAQ  
Quantumpay Protocol

This document answers common questions for node operators,
reviewers, and early participants during Phase 1 (P1).

## 1. What is Quantumpay at this stage?

Quantumpay is a **Minimum Usable Blockchain** focused on:
- Network stability
- Deterministic execution
- Reliable block production

At P1, Quantumpay is intentionally minimal and conservative.

## 2. Is Quantumpay mainnet live?

Quantumpay is operating in **early mainnet / quiet mainnet mode**.

This means:
- The network is live and producing blocks
- Stability and operations are prioritized
- No public APIs, SDKs, or applications are enabled yet


## 3. Why are there no public APIs or gRPC endpoints?

Public APIs are intentionally **deferred**.

Reasons:
- Reduce attack surface
- Maintain determinism
- Focus on core network reliability

APIs will be introduced in later phases (P2+) after sufficient
operational maturity is proven.

## 4. Can anyone run a node?

Yes.

Anyone with:
- A compatible Linux system
- Stable internet connection
- Proper configuration

can run a Quantumpay node by following `NETWORK_QUICKSTART.md`.


## 5. How does peer discovery work?

At P1, Quantumpay uses a **static peer model**.

Nodes connect to explicitly configured peers.
Automatic discovery is intentionally disabled to reduce complexity
and unexpected network behavior.


## 6. What happens if a node goes offline?

If a node stops:
- The network continues operating
- The node can rejoin later
- No special intervention is required

This behavior is expected and safe.


## 7. How do I verify my node is healthy?

A healthy node:
- Runs continuously without restarts
- Shows increasing block height
- Produces no repeated errors in logs

Basic checks:
```bash
systemctl status quantumpay-node
journalctl -u quantumpay-node
