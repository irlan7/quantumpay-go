function getLeader(validators, height) {
  const ids = Object.keys(validators).sort();
  if (ids.length === 0) return null;
  return ids[height % ids.length];
}

module.exports = { getLeader };

