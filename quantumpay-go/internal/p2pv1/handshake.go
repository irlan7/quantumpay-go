package p2pv1

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"time"
)

// HandshakeMessage is signed identity proof
type HandshakeMessage struct {
	NodeID    string `json:"node_id"`
	Timestamp int64  `json:"timestamp"`
	PubKey    []byte `json:"pubkey"` // marshaled
	Signature []byte `json:"signature"`
}

// BuildHandshake creates signed handshake payload
func BuildHandshake(
	nodeID string,
	priv *ecdsa.PrivateKey,
	pubBytes []byte,
) (*HandshakeMessage, error) {

	msg := HandshakeMessage{
		NodeID:    nodeID,
		Timestamp: time.Now().Unix(),
		PubKey:    pubBytes,
	}

	payload, err := json.Marshal(struct {
		NodeID    string
		Timestamp int64
	}{
		NodeID:    msg.NodeID,
		Timestamp: msg.Timestamp,
	})
	if err != nil {
		return nil, err
	}

	sig, err := SignMessage(priv, payload)
	if err != nil {
		return nil, err
	}

	msg.Signature = sig
	return &msg, nil
}

// VerifyHandshake validates signature & timestamp
func VerifyHandshake(
	msg *HandshakeMessage,
	pub *ecdsa.PublicKey,
	maxSkewSec int64,
) error {

	if msg == nil {
		return errors.New("nil handshake")
	}

	now := time.Now().Unix()
	if abs(now-msg.Timestamp) > maxSkewSec {
		return errors.New("handshake timestamp skew too large")
	}

	payload, err := json.Marshal(struct {
		NodeID    string
		Timestamp int64
	}{
		NodeID:    msg.NodeID,
		Timestamp: msg.Timestamp,
	})
	if err != nil {
		return err
	}

	if !VerifyMessage(pub, payload, msg.Signature) {
		return errors.New("invalid handshake signature")
	}

	return nil
}

func abs(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}
