# Quantumpay Network Health Check

**Document Type:** Operational Reference  
**Audience:** Node Operators, Core Maintainers, Governance  
**Scope:** Lightweight, repeatable health verification for Quantumpay network  
**Status:** Active (Pre-Mainnet & Mainnet)

## 1. Purpose

This document defines a **minimal, deterministic health check** to verify that the Quantumpay blockchain network is operating normally.

Health checks are designed to be:
- Simple
- Fast (≤ 2 minutes)
- Non-intrusive
- Executable without special tooling

## 2. Definition of a “Healthy Network”

A Quantumpay network is considered **healthy** when:

- Nodes are running continuously
- Blocks are produced at expected intervals
- Peers remain connected
- Resource usage is stable
- No error loops or crash patterns appear in logs

## 3. Mandatory Health Checks (Node-Level)

### 3.1 Service Status

```bash
sudo systemctl status quantumpay-node --no-pager
