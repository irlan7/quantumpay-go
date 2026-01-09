# Quantumpay systemd Service

## Purpose
Menjalankan Quantumpay node sebagai background service yang:
- Auto-restart
- Tahan reboot
- Aman untuk operasi 24/7

## Deployment
File ini **TIDAK dieksekusi di WSL**.

Langkah deploy di VPS:
1. Copy file ke `/etc/systemd/system/`
2. Reload systemd
3. Enable & start service

## Example
```bash
sudo cp quantumpay-node.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable quantumpay-node
sudo systemctl start quantumpay-node

