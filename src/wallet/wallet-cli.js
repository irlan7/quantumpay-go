#!/usr/bin/env node

/**
 * QuantumPay Wallet CLI
 * Hybrid Signature Wallet
 * ECDSA (secp256k1) + Dilithium (PQC Placeholder)
 */

const fs = require('fs');
const path = require('path');
const crypto = require('crypto');
const axios = require('axios');
const EC = require('elliptic').ec;
const yargs = require('yargs');

const ec = new EC('secp256k1');

const NODE_URL = 'http://localhost:3001';
const WALLET_DIR = path.join(__dirname, '.wallet');
const WALLET_FILE = path.join(WALLET_DIR, 'wallet.json');

/* =====================================================
   UTIL
===================================================== */

function sha256(data) {
  return crypto.createHash('sha256').update(data).digest('hex');
}

/**
 * Dilithium placeholder (PQC-ready)
 * GANTI dengan liboqs / wasm nanti
 */
function dilithiumSign(message, privateKey) {
  return sha256(message + privateKey);
}

/* =====================================================
   WALLET CORE
===================================================== */

function ensureWalletDir() {
  if (!fs.existsSync(WALLET_DIR)) {
    fs.mkdirSync(WALLET_DIR);
  }
}

function createWallet() {
  ensureWalletDir();

  if (fs.existsSync(WALLET_FILE)) {
    console.log('⚠ Wallet already exists');
    return;
  }

  // ECDSA keypair
  const keyPair = ec.genKeyPair();
  const ecdsaPrivateKey = keyPair.getPrivate('hex');
  const ecdsaPublicKey = keyPair.getPublic('hex');

  // Dilithium (PQC placeholder keys)
  const dilithiumPrivateKey = crypto.randomBytes(32).toString('hex');
  const dilithiumPublicKey = sha256(dilithiumPrivateKey);

  // Address (hybrid derived)
  const address = sha256(ecdsaPublicKey + dilithiumPublicKey).slice(0, 40);

  const wallet = {
    address,
    ecdsa: {
      privateKey: ecdsaPrivateKey,
      publicKey: ecdsaPublicKey
    },
    dilithium: {
      privateKey: dilithiumPrivateKey,
      publicKey: dilithiumPublicKey
    },
    createdAt: Date.now()
  };

  fs.writeFileSync(WALLET_FILE, JSON.stringify(wallet, null, 2));

  console.log('✅ Wallet created');
  console.log('Address:', address);
}

function loadWallet() {
  if (!fs.existsSync(WALLET_FILE)) {
    console.error('❌ Wallet not found. Run: wallet-cli.js create');
    process.exit(1);
  }
  return JSON.parse(fs.readFileSync(WALLET_FILE));
}

/* =====================================================
   HYBRID SIGN
===================================================== */

function signHybridTransaction(tx, wallet) {
  const txHash = sha256(
    JSON.stringify({
      from: tx.from,
      to: tx.to,
      amount: tx.amount,
      timestamp: tx.timestamp
    })
  );

  // ECDSA SIGN
  const key = ec.keyFromPrivate(wallet.ecdsa.privateKey);
  const signatureECDSA = key.sign(txHash).toDER('hex');

  // DILITHIUM SIGN (placeholder)
  const signatureDilithium = dilithiumSign(
    txHash,
    wallet.dilithium.privateKey
  );

  return {
    signatureECDSA,
    signatureDilithium
  };
}

/* =====================================================
   SEND TX
===================================================== */

async function sendTransaction(to, amount) {
  const wallet = loadWallet();

  const tx = {
    from: wallet.address,
    to,
    amount,
    timestamp: Date.now()
  };

  const signatures = signHybridTransaction(tx, wallet);

  const payload = {
    ...tx,
    publicKey: wallet.ecdsa.publicKey,
    signatureECDSA: signatures.signatureECDSA,
    publicKeyDilithium: wallet.dilithium.publicKey,
    signatureDilithium: signatures.signatureDilithium
  };

  try {
    const res = await axios.post(`${NODE_URL}/transactions/new`, payload);
    console.log('🚀 Transaction sent');
    console.log(res.data);
  } catch (err) {
    console.error(
      '❌ Failed:',
      err.response?.data || err.message
    );
  }
}

/* =====================================================
   CLI
===================================================== */

yargs
  .command('create', 'Create new wallet', {}, createWallet)
  .command(
    'send',
    'Send transaction',
    y => {
      y.option('to', {
        type: 'string',
        describe: 'Destination address',
        demandOption: true
      });
      y.option('amount', {
        type: 'number',
        describe: 'Amount to send',
        demandOption: true
      });
    },
    argv => sendTransaction(argv.to, argv.amount)
  )
  .demandCommand(1)
  .help()
  .argv;

