const fs = require('fs');
const path = require('path');
const crypto = require('crypto');

const CERT_DIR = path.join(__dirname, '../../certs');
const KEY_PATH = path.join(CERT_DIR, 'node.key');
const CERT_PATH = path.join(CERT_DIR, 'node.crt');

function ensureCertDir() {
  if (!fs.existsSync(CERT_DIR)) {
    fs.mkdirSync(CERT_DIR, { recursive: true });
  }
}

function generateSelfSignedCert() {
  console.log('🔐 Generating self-signed TLS cert for P2P node');

  const key = crypto.generateKeyPairSync('rsa', {
    modulusLength: 2048,
  });

  const privateKey = key.privateKey.export({ type: 'pkcs1', format: 'pem' });

  // Simple placeholder cert (DEV / TESTNET)
  const cert = privateKey; 
  // ⛔ NOTE: Untuk mainnet → pakai OpenSSL / CA / validator cert

  fs.writeFileSync(KEY_PATH, privateKey);
  fs.writeFileSync(CERT_PATH, cert);
}

function loadTLSConfig() {
  ensureCertDir();

  if (!fs.existsSync(KEY_PATH) || !fs.existsSync(CERT_PATH)) {
    generateSelfSignedCert();
  }

  return {
    key: fs.readFileSync(KEY_PATH),
    cert: fs.readFileSync(CERT_PATH),
    rejectUnauthorized: false // P2P trust handled at protocol layer
  };
}

module.exports = { loadTLSConfig };

