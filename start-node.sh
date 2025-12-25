#!/bin/bash

ROOT_DIR=$(pwd)
SRC="$ROOT_DIR/src/server.js"

# Safety check
if [ ! -f "$SRC" ]; then
  echo "❌ server.js not found at $SRC"
  exit 1
fi

echo "🚀 Starting QuantumPay multi-node testnet..."

# Kill old nodes (optional)
pkill -f "node src/server.js" 2>/dev/null

sleep 1

# -------- NODE 1 --------
NODE_ID=1 \
NODE_NAME=node1 \
DATA_DIR="$ROOT_DIR/data/node1" \
P2P_PORT=4001 \
RPC_PORT=5001 \
node "$SRC" > "$ROOT_DIR/data/node1/node.log" 2>&1 &

echo "✅ Node1 started (P2P:4001 RPC:5001)"

# -------- NODE 2 --------
NODE_ID=2 \
NODE_NAME=node2 \
DATA_DIR="$ROOT_DIR/data/node2" \
P2P_PORT=4002 \
RPC_PORT=5002 \
node "$SRC" > "$ROOT_DIR/data/node2/node.log" 2>&1 &

echo "✅ Node2 started (P2P:4002 RPC:5002)"

# -------- NODE 3 --------
NODE_ID=3 \
NODE_NAME=node3 \
DATA_DIR="$ROOT_DIR/data/node3" \
P2P_PORT=4003 \
RPC_PORT=5003 \
node "$SRC" > "$ROOT_DIR/data/node3/node.log" 2>&1 &

echo "✅ Node3 started (P2P:4003 RPC:5003)"

echo ""
echo "🌐 QuantumPay testnet is LIVE"
echo "📂 Logs:"
echo "   tail -f data/node1/node.log"
echo "   tail -f data/node2/node.log"
echo "   tail -f data/node3/node.log"

