
CONSENSUS_ID.md

QuantumPay Consensus Specification

PoS + BFT Finality

1. Overview

QuantumPay menggunakan mekanisme konsensus Proof of Stake (PoS) yang dikombinasikan dengan Byzantine Fault Tolerant (BFT) Finality.

Desain ini bertujuan untuk:

Menyediakan finalitas cepat dan deterministik

Menghindari konsumsi energi berlebih seperti pada Proof of Work

Mendukung jaringan publik jangka panjang yang stabil, aman, dan dapat diskalakan

Memastikan konsistensi state pada seluruh node tanpa fork berkepanjangan

Konsensus ini diimplementasikan sepenuhnya di lapisan core (on-chain) dan tidak bergantung pada logika off-chain atau pihak terpusat.


2. Consensus Model

2.1 High-Level Model

QuantumPay consensus terdiri dari dua lapisan utama:

1. Block Production (PoS)


2. Block Finality (BFT)


Stakeholders → Validator Set
Validator Set → Block Proposer
Proposed Block → BFT Voting → Finalized Block


2.2 Proof of Stake (PoS)

Validator dipilih berdasarkan stake yang dikunci (bonded stake) di jaringan.

Karakteristik:

Tidak ada mining berbasis hash

Tidak ada kompetisi komputasi

Pemilihan proposer bersifat deterministik / pseudo-random

Stake berfungsi sebagai economic security


Stake yang terkunci dapat dikenakan slashing apabila validator berperilaku jahat atau lalai.


2.3 BFT Finality

Setelah blok diproduksi, blok tersebut tidak langsung final.

Blok harus melewati proses voting BFT oleh mayoritas validator.

Finality dicapai jika:

≥ 2/3 validator (berdasarkan total stake) menyetujui blok

Tidak ada konflik state

Tidak terjadi double-sign atau equivocation


Setelah final:

Blok tidak dapat di-reorg

Transaksi dianggap irreversible

State jaringan menjadi konsisten secara global


3. Validator Responsibilities

Validator bertanggung jawab atas:

Menjaga node online dan sinkron

Memvalidasi transaksi dan state

Berpartisipasi dalam voting finality

Menghindari perilaku byzantine


Validator tidak memiliki otoritas governance absolut dan tunduk pada aturan protokol.


4. Slashing & Fault Handling

4.1 Slashing Conditions

Validator dapat dikenai sanksi jika melakukan:

Double signing

Voting konflik pada ronde finality yang sama

Extended downtime (offline)


4.2 Slashing Effects

Pengurangan stake

Temporary jailing

Permanent removal (untuk pelanggaran berat)


5. Security Assumptions

Konsensus aman selama:

≤ 1/3 total stake bersifat byzantine

Validator mengikuti aturan protokol

Kunci privat validator dijaga dengan benar


Desain ini mengikuti model keamanan klasik Practical Byzantine Fault Tolerance (PBFT).


6. Comparison with Proof of Work (PoW)

Aspect                PoW	          QuantumPay PoS + BFT

Energy Usage  	    Sangat tinggi	  Sangat rendah
Finality	          Probabilistik	  Deterministik
Fork                Risk	Tinggi	  Sangat rendah
Hardware Arms Race  Ada	            Tidak ada
Time to Finality    Menit–jam	      Detik
Sustainability	    Rendah	        Tinggi


7. Determinism Guarantee

QuantumPay menjamin:

> Same input → same state → same block hash → same finalized chain


Tidak ada:

Randomness non-deterministik

Time-based logic

External entropy


8. Upgrade & Evolution

Perubahan mekanisme konsensus:

Tidak dilakukan melalui patch sembarangan

Harus melalui deklarasi protokol dan jaringan baru

Konsensus lama dianggap frozen setelah mainnet


9. Scope Clarification

Dokumen ini hanya mendeskripsikan konsensus layer, tidak mencakup:

Governance sosial

Token economics detail

Smart contract VM

Application-layer logic


10. Status

Consensus Design: Final

Implementation Layer: Go Core (on-chain)

Mainnet Readiness: Verified

Last Updated: Mainnet (Post-Genesis)
