# Network Quickstart  
Quantumpay Protocol  
(5–10 Minutes Node Setup)

## Purpose

This document provides a **minimal, safe, and fast** guide to run a
Quantumpay node for the first time.

It is intended for:
- Node operators
- Reviewers / auditors
- Early network participants

This guide assumes **Phase 1 (P1) is frozen** and focuses on **stability**, not features.

## System Requirements

Minimum recommended:

- OS: Linux (Ubuntu 20.04+ recommended)
- CPU: 2 cores
- RAM: 2 GB (4 GB recommended)
- Disk: 50 GB free space
- Network: Public IP, stable connection


## Step 1 — Create Service User

Run as root or via sudo:

bash
sudo useradd --system \
  --home /opt/quantumpay \
  --shell /usr/sbin/nologin \
  quantumpay

Create directories:

Bash
sudo mkdir -p /opt/quantumpay/{bin,data,config,logs}
sudo chown -R quantumpay:quantumpay /opt/quantumpay
sudo chmod 750 /opt/quantumpay


Step 2 — Install Binary
Copy the quantumpay-node binary to:

Bash
/opt/quantumpay/bin/quantumpay-node
Ensure permissions:

Bash
sudo chown quantumpay:quantumpay /opt/quantumpay/bin/quantumpay-node
sudo chmod 750 /opt/quantumpay/bin/quantumpay-node


Step 3 — Basic Configuration
Create a minimal config file:

Bash
sudo -u quantumpay nano /opt/quantumpay/config/node.toml
Example:

Toml
chain_id = 77077

[p2p]
listen_addr = "0.0.0.0:7001"
peers = [
  "NODE1_PUBLIC_IP:7001"
]
Replace NODE1_PUBLIC_IP with an existing peer.

Step 4 — Create systemd Service
Create service file:
Bash
sudo nano /etc/systemd/system/quantumpay-node.service
Contents:

Ini
[Unit]
Description=Quantumpay Blockchain Node
After=network.target

[Service]
User=quantumpay
Group=quantumpay
WorkingDirectory=/opt/quantumpay

ExecStart=/opt/quantumpay/bin/quantumpay-node \
  --config /opt/quantumpay/config/node.toml

Restart=always
RestartSec=5
LimitNOFILE=65535
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
Reload and start:

Bash
sudo systemctl daemon-reload
sudo systemctl enable quantumpay-node
sudo systemctl start quantumpay-node

Step 5 — Verify Node Health
Check service status:

Bash
sudo systemctl status quantumpay-node --no-pager
Check logs:

Bash
journalctl -u quantumpay-node -n 20 --no-pager
Healthy output indicators:
Node starts without panic
Block height increases
Peer connection established


Step 6 — Confirm Block Progress
Wait 1–2 minutes and confirm logs show:
New block produced or received
Height increasing
No repeated error messages
If block height advances, your node is successfully connected.
Notes & Safety Rules
This node runs in P1 (frozen) mode
Do NOT modify:
Chain ID
Genesis
Core configuration
No APIs, SDKs, or smart contracts are required
If unsure, stop and ask before changing anything.

Stopping the Node
To stop safely:

Bash
sudo systemctl stop quantumpay-node
Data remains intact.


Troubleshooting (Minimal)
Node does not start → check config syntax
No blocks → check peer address & firewall
Repeated restart → inspect logs carefully
If issues persist, do not attempt random fixes.
Final Note
A healthy Quantumpay node is quiet, stable, and predictable.
If your node runs without intervention, it is operating correctly.
End of Document
