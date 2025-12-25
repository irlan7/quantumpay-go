'use strict';

class KeyRotation {
  constructor({ state }) {
    this.state = state;
  }

  rotateKey({
    validatorId,
    newKey
  }) {
    const validator = this.state.validators[validatorId];
    if (!validator || validator.jailed) {
      throw new Error('Invalid or jailed validator');
    }

    validator.keys = validator.keys || [];

    // retire active key
    for (const key of validator.keys) {
      if (key.status === 'active') {
        key.status = 'retired';
      }
    }

    const version =
      validator.keys.length > 0
        ? Math.max(...validator.keys.map(k => k.version)) + 1
        : 1;

    const now = Math.floor(Date.now() / 1000);

    validator.keys.push({
      version,
      type: newKey.type,       // 'ecdsa' | 'dilithium'
      pubKey: newKey.pubKey,
      createdAt: now,
      expiresAt: newKey.expiresAt,
      status: 'active'
    });

    return {
      validatorId,
      newVersion: version
    };
  }

  isKeyValid(validatorId, keyVersion) {
    const validator = this.state.validators[validatorId];
    if (!validator) return false;

    const key = validator.keys.find(
      k => k.version === keyVersion
    );

    if (!key) return false;
    if (key.status !== 'active') return false;

    const now = Math.floor(Date.now() / 1000);
    if (now > key.expiresAt) return false;

    return true;
  }

  getActiveKey(validatorId) {
    const validator = this.state.validators[validatorId];
    if (!validator) return null;

    return validator.keys.find(
      k => k.status === 'active'
    );
  }
}

module.exports = KeyRotation;

