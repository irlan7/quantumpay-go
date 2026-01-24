
📘 Panduan Instalasi

Menjalankan QuantumPay Go Core Node (Lokal)

Panduan ini menjelaskan langkah-langkah untuk menginstal dan menjalankan QuantumPay Core Node di komputer lokal menggunakan Go.
Dokumen ini ditujukan untuk developer, auditor, dan operator node.


1️⃣ Persyaratan Sistem

Sistem Operasi

Linux (Ubuntu 20.04 / 22.04 direkomendasikan)

macOS (Intel / Apple Silicon)

Windows (via WSL2 direkomendasikan)


Perangkat Lunak

Go versi 1.21 atau lebih baru

Git

Akses terminal (bash/zsh)


Spesifikasi Minimum

CPU: 2 Core

RAM: 4 GB (8 GB direkomendasikan)

Storage: ≥ 20 GB kosong

Koneksi internet stabil


2️⃣ Instalasi Go

Linux / macOS

go version

Jika belum terpasang, instal Go dari:

https://go.dev/dl/

Pastikan Go sudah terdeteksi:

export PATH=$PATH:/usr/local/go/bin


3️⃣ Clone Repository QuantumPay

git clone https://github.com/irlan7/quantumpay-go.git
cd quantumpay-go/quantumpay-go

Pastikan struktur direktori terlihat seperti ini:

ls
cmd/  internal/  config/  go.mod  go.sum


4️⃣ Konfigurasi Node

Masuk ke folder konfigurasi:

cd config

Pastikan file genesis.json tersedia dan sesuai jaringan yang digunakan (testnet / mainnet).

Jika menjalankan node lokal:

Tidak perlu mengubah konfigurasi P2P

Port default aman untuk localhost


Kembali ke root:

cd ..


5️⃣ Build & Jalankan Node

Menjalankan Langsung (Development Mode)

go run ./cmd/node

Jika berhasil, akan muncul log seperti:

[NODE] Starting QuantumPay node
[ENGINE] Block produced height=1



Menjalankan dengan gRPC (Read-Only – Aman)

go run ./cmd/node --grpc --grpc-addr 127.0.0.1:9090

> ⚠️ gRPC bersifat read-only, tidak mengubah state blockchain.


6️⃣ Verifikasi Node Berjalan

Cek Produksi Block

tail -f logs/node.log

atau langsung dari terminal output:

[ENGINE] New block produced height=...

Cek Port Aktif

ss -tulpen | grep 7001



7️⃣ Struktur Modul Core (Ringkas)

internal/
├── blockchain/
├── consensus/
├── core/
├── engine/
├── coin/          # Economic & supply logic
├── state/
├── p2p/
├── rpc/
└── grpc_disabled/ # gRPC placeholder (aman)

> ⚠️ Core protocol berjalan tanpa API publik
API dan explorer akan hadir di fase berikutnya.


8️⃣ Hentikan Node

Tekan:

CTRL + C

Node akan shutdown dengan aman:

[NODE] Shutdown signal received
[NODE] Exiting cleanly


9️⃣ Catatan Keamanan

Jangan expose port P2P ke publik tanpa firewall

Jangan ubah file di internal/ tanpa audit

Gunakan branch master untuk stabilitas



🔒 Scope Notice (Penting)

> Rilis mainnet awal ini berfokus pada stabilitas protokol inti.
API publik, block explorer, dan tooling ekosistem akan diperkenalkan pada fase berikutnya.


✅ Status

Core engine: Stabil

Multi-node: Aktif

Economic (coin): Aktif

Cocok untuk:

Audit

Integrasi internal

Persiapan ekosistem

