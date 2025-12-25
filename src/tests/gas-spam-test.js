const NODE_URL = 'http://localhost:3001/transaction';
const TOTAL_TX = 50;

async function sendSpamTx(i) {
  const tx = {
    from: 'SPAMMER',
    to: 'VICTIM',
    amount: 1,
    gas: 1, // deliberately LOW gas
    nonce: i,
    signature: 'INVALID',
    chainId: 'quantumpay-testnet'
  };

  try {
    const res = await fetch(NODE_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(tx)
    });

    const text = await res.text();
    console.log(`TX ${i} → status ${res.status}: ${text}`);
  } catch (err) {
    console.error(`TX ${i} network error`, err.message);
  }
}

(async () => {
  console.log('🚀 Starting Gas Spam Stress Test');

  for (let i = 0; i < TOTAL_TX; i++) {
    await sendSpamTx(i);
  }

  console.log('✅ Stress test finished');
})();

