#!/bin/bash

# =================================================================
# QuantumPay Node Automated Installer v1.1-alpha
# Network Identity: Chain ID 77077 (Mainnet-Alpha)
# =================================================================

# Warna untuk output
BLUE='\033[0;34m'
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}🌐 Memulai Instalasi QuantumPay Node...${NC}"

# 1. Update Sistem & Instalasi Dependency Dasar
echo -e "${BLUE}📦 Menginstal dependency sistem...${NC}"
sudo apt-get update && sudo apt-get install -y git wget build-essential curl

# 2. Instalasi Go-Lang (v1.21+)
if ! command -v go &> /dev/null; then
    echo -e "${BLUE}🐹 Menginstal Go-Lang...${NC}"
    wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
    echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.bashrc
    source $HOME/.bashrc
    export PATH=$PATH:/usr/local/go/bin
else
    echo -e "${GREEN}✅ Go-Lang sudah terpasang.${NC}"
fi

# 3. Instalasi Node.js & PM2 (Untuk Manajemen Proses)
if ! command -v pm2 &> /dev/null; then
    echo -e "${BLUE}🚀 Menginstal PM2 untuk manajemen proses...${NC}"
    sudo apt-get install -y nodejs npm
    sudo npm install pm2 -g
else
    echo -e "${GREEN}✅ PM2 sudah terpasang.${NC}"
fi

# 4. Clone Repositori Core
echo -e "${BLUE}🏗️ Mengunduh QuantumPay Core dari GitHub...${NC}"
rm -rf quantumpay-go
git clone https://github.com/irlan7/quantumpay-go
cd quantumpay-go

# 5. Build Binary
echo -e "${BLUE}⚙️ Membangun (Build) QuantumPay Node...${NC}"
go build -o quantumpay-node ./cmd/node

# 6. Menjalankan Node dengan PM2
echo -e "${BLUE}⚡ Menjalankan Node...${NC}"
pm2 start ./quantumpay-node --name "qp-node"

# 7. Finalisasi & Instruksi Verifikasi
echo -e "${GREEN}====================================================${NC}"
echo -e "${GREEN}✅ INSTALASI BERHASIL!${NC}"
echo -e "${GREEN}====================================================${NC}"
echo -e "Node Anda sekarang berjalan di latar belakang via PM2."
echo -e ""
echo -e "Silakan verifikasi Genesis Hash Anda sekarang:"
echo -e "${BLUE}pm2 logs qp-node --lines 100 | grep \"Genesis Hash\"${NC}"
echo -e ""
echo -e "Expected Hash: ${GREEN}0x1d58599424f1159828236111f1f9e83063f66345091a99540c4989679269491a${NC}"
echo -e "===================================================="
