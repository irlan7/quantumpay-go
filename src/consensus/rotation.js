const {
  EPOCH_LENGTH,
  MIN_STAKE,
  MIN_REPUTATION
} = require("./constants");

const { selectLeader } = require("./leaderElection");

function getEpoch(height) {
  return Math.floor(height / EPOCH_LENGTH);
}

function rotateValidators({
  height,
  chainId,
  validatorSet
}) {
  const epoch = getEpoch(height);

  const active = validatorSet.getActive(
    MIN_STAKE,
    MIN_REPUTATION
  );

  const leader = selectLeader(active, epoch, chainId);

  return {
    epoch,
    leader,
    activeValidators: active.map(v => v.id)
  };
}

module.exports = {
  rotateValidators,
  getEpoch
};

