const Wallet = require('../wallet');

const { publicKey, privateKey } = Wallet.generateKeyPair();

console.log("\n=== PUBLIC KEY PEM ===\n", publicKey);
console.log("\n=== PRIVATE KEY PEM ===\n", privateKey);

const address = Wallet.getAddressFromPublicKey(publicKey);
console.log("\nWallet Address:", address);

// test sign
const message = "Hello QuantumPay!";
const signature = Wallet.signData(privateKey, message);
console.log("\nSignature Hex:", signature);

// verify
const valid = Wallet.verifySignature(publicKey, message, signature);
console.log("\nSignature Valid?", valid);
