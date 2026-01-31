# 🔒 Kebijakan Keamanan (Security Policy)

## 1. Versi yang Didukung

Repositori ini berisi **implementasi protokol inti (core protocol)** dari blockchain QuantumPay.

Hanya rilis bertag terbaru (misalnya `v1.x.x-core`) yang menerima pembaruan keamanan.

| Versi            | Status       | Catatan                 |
|------------------|--------------|--------------------------|
| `main` / `master`| ✅ Aktif      | Branch pengembangan      |
| `v1.x.x-core`    | ✅ Dipelihara | Rilis stabil             |
| Versi lama       | ❌ Tidak didukung | Harap lakukan pembaruan |


## 2. Pelaporan Kerentanan Keamanan

Jika Anda menemukan kerentanan keamanan **pada protokol inti**, mohon **laporkan secara privat**.

### Kontak Keamanan:
📧 **quantumpaysec [at] gmail [dot] com**

(Jika alamat email tidak dapat diakses, silakan gunakan fitur **GitHub Private Security Advisory**.)

❗ **Jangan** membuka *issue publik* untuk laporan keamanan.

Kami menghargai praktik *responsible disclosure* dan berkomitmen untuk memberikan konfirmasi awal dalam waktu **maksimal 48 jam**.

## 3. Ruang Lingkup Tanggung Jawab

Repositori ini mencakup:
- Mesin blockchain inti (core blockchain engine)
- Transisi state deterministik
- Produksi dan validasi blok
- Penyimpanan chain dan *read-only views*

Repositori ini **tidak mencakup**:
- Wallet
- UI / frontend
- SDK
- Smart contract VM
- Aplikasi lapisan atas (application layer)

Isu keamanan di luar ruang lingkup tersebut tidak ditangani melalui repositori ini.

## 4. Prinsip Keamanan

QuantumPay dirancang dengan prinsip:
- Determinisme penuh
- Permukaan serangan minimal
- Arsitektur modular
- Stabilitas protokol jangka panjang

Keamanan protokol adalah prioritas utama dan setiap perubahan kritikal mengikuti proses evaluasi yang ketat.
