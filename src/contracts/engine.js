const DOMAIN = require('../crypto/domain');
const { verifyWithDomain } = require('../crypto/verify');

class ContractEngine {
  static execute({ contract, method, args, caller, signature }) {
    const domain = DOMAIN.CONTRACT(contract, method);

    const ok = verifyWithDomain({
      publicKey: caller,
      domain,
      payload: args,
      signature
    });

    if (!ok) {
      throw new Error('Invalid contract signature');
    }

    // execute contract
  }
}

module.exports = ContractEngine;

