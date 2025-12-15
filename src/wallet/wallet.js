const EC = require('elliptic').ec;
const ec = new EC('secp256k1');
const fs = require('fs');
const path = require('path');

const WALLET_FILE = path.join(__dirname, 'wallet.json');

class Wallet {
    constructor(privateKey) {
        this.keyPair = privateKey
            ? ec.keyFromPrivate(privateKey, 'hex')
            : ec.genKeyPair();

        this.publicKey = this.keyPair.getPublic('hex');
    }

    sign(dataHash) {
        return this.keyPair.sign(dataHash).toDER('hex');
    }

    createTransaction({ to, amount }) {
        const tx = {
            from: this.publicKey,
            to,
            amount,
            timestamp: Date.now()
        };

        const hash = Wallet.hash(tx);
        tx.signature = this.sign(hash);
        return tx;
    }

    static hash(data) {
        return require('crypto')
            .createHash('sha256')
            .update(JSON.stringify(data))
            .digest('hex');
    }

    save() {
        fs.writeFileSync(WALLET_FILE, JSON.stringify({
            privateKey: this.keyPair.getPrivate('hex'),
            publicKey: this.publicKey
        }, null, 2));
    }

    static load() {
        const data = JSON.parse(fs.readFileSync(WALLET_FILE));
        return new Wallet(data.privateKey);
    }
}

module.exports = Wallet;

