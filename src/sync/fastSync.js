const axios = require('axios');
const StateSnapshot = require('../state/snapshot');

class FastSync {
  /**
   * peers: array URL peer (http)
   */
  static async sync({ blockchain, peers }) {
    for (const peer of peers) {
      try {
        console.log('🔍 Trying snapshot from', peer);

        // 1. download snapshot
        const snapshotResp = await axios.get(`${peer}/snapshot`);
        const snapshot = snapshotResp.data;

        // 2. verify signature
        if (snapshot.signature) {
          const valid = StateSnapshot.verify(snapshot);
          if (!valid) {
            console.warn('❌ Invalid snapshot signature from', peer);
            continue;
          }
        }

        // 3. download blocks after snapshot
        const blocksResp = await axios.get(
          `${peer}/blocks?from=${snapshot.height + 1}`
        );

        const blocks = blocksResp.data;

        // 4. apply snapshot
        blockchain.chain = blockchain.chain.slice(0, snapshot.height + 1);
        blockchain.state = {
          balances: snapshot.balances,
          nonces: snapshot.nonces,
          contracts: snapshot.contracts
        };

        // 5. apply blocks incrementally
        for (const block of blocks) {
          blockchain.addBlock({ data: block.data });
        }

        console.log('✅ Fast sync success from', peer);
        return true;
      } catch (err) {
        console.warn('Fast sync failed from', peer, err.message);
      }
    }

    throw new Error('Fast sync failed from all peers');
  }
}

module.exports = FastSync;
