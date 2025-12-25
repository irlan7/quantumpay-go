class Evidence {
  static doubleSign({ snapshotA, snapshotB }) {
    return (
      snapshotA.height === snapshotB.height &&
      snapshotA.signer === snapshotB.signer &&
      snapshotA.stateRoot !== snapshotB.stateRoot
    );
  }
}

module.exports = Evidence;

