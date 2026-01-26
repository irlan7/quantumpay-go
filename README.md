QuantumPay is a high-performance, sovereign Layer-1 blockchain infrastructure built for national digital sovereignty and global scalability. Engineered in Go-Lang, it achieves near-instant finality through a customized BFT-based consensus mechanism while maintaining an ultra-light resource footprint.
🌐 Network Identity (SSoT)
As the Single Source of Truth, these parameters define the current state of the frozen mainnet-alpha:
| Parameter | Value |
|-------------------|----------------|
| Chain ID          | 77001 [FROZEN] |
| Genesis Hash      | acc395137e5d0c28c609d011ea99d89405f07009c0bbf8933711e1a7f184edc6 |
| Core Engine       | Go-Lang (quantumpay-go-v1.1) |
| Official Provider | VPS-9451c332 |
⚡ Technical Features
 * High Efficiency: Optimized to consume as little as 5% RAM on standard hardware, ensuring maximum decentralization.
 * Fast Consensus: Block production and finality achieved in under 5 seconds.
 * Production Process Management: Utilizes PM2 for 24/7 bridge and node uptime with minimal memory overhead (~24.7MB).
🛠 Run a Node (Join the Decentralization)
Follow these steps to synchronize your validator with the QuantumPay Network:
Hardware Requirements
 * CPU: 2 Cores (Minimum)
 * RAM: 4GB (Optimized usage: ~5%)
 * Storage: 40GB SSD
 * OS: Ubuntu 22.04 LTS / 24.04 LTS
Installation
# Clone the core repository
git clone https://github.com/irlan7/quantumpay-go
cd quantumpay-go

# Build the node
go build -o quantumpay-node ./cmd/node

# Start the node (Recommended with PM2)
pm2 start ./quantumpay-node --name "qp-node"

📡 Official Channels
 * Website: https://quantumpaychain.org
 * Explorer: /explorer
 * X (Twitter): @quantumpaychain
 * Email: quantumpaysec@gmail.com
📜 License & Vision
This project is open-source under the MIT License. We follow the vision of Satoshi Nakamoto and Vitalik Buterin: building a trustless, transparent, and permissionless world where every individual can run a node and verify the state of the truth.
Would you like me to create a bash script (install.sh) that automates this entire process for new users who want to join your network?
