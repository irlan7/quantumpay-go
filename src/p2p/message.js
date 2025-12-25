function encodeMessage(type, payload) {
  return JSON.stringify({
    type,
    payload,
    timestamp: Date.now()
  });
}

function decodeMessage(data) {
  try {
    const msg = JSON.parse(data.toString());
    if (!msg.type || !msg.payload) return null;
    return msg;
  } catch {
    return null;
  }
}

module.exports = { encodeMessage, decodeMessage };

