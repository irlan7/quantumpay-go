package p2pv1

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

// ===========================
// Node Identity (Persistent)
// ===========================

const (
	defaultKeyDir  = ".quantumpay"
	defaultKeyFile = "nodekey.json"
)

type storedKey struct {
	PrivateKey []byte `json:"private_key"` // DER encoded
}

// LoadOrCreateNodeKey loads node key from disk or creates a new one
func LoadOrCreateNodeKey(basePath string) (*ecdsa.PrivateKey, error) {
	if basePath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		basePath = filepath.Join(home, defaultKeyDir)
	}

	if err := os.MkdirAll(basePath, 0700); err != nil {
		return nil, err
	}

	keyPath := filepath.Join(basePath, defaultKeyFile)

	// 🔁 Load existing key
	if _, err := os.Stat(keyPath); err == nil {
		return loadKey(keyPath)
	}

	// ✨ Create new key
	priv, err := GenerateKey()
	if err != nil {
		return nil, err
	}

	if err := saveKey(keyPath, priv); err != nil {
		return nil, err
	}

	return priv, nil
}

// ===========================
// Helpers
// ===========================

func saveKey(path string, priv *ecdsa.PrivateKey) error {
	der, err := x509.MarshalECPrivateKey(priv)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(storedKey{
		PrivateKey: der,
	}, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0600)
}

func loadKey(path string) (*ecdsa.PrivateKey, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var sk storedKey
	if err := json.Unmarshal(raw, &sk); err != nil {
		return nil, err
	}

	priv, err := x509.ParseECPrivateKey(sk.PrivateKey)
	if err != nil {
		return nil, err
	}

	if priv.Curve != elliptic.P256() {
		return nil, errors.New("unsupported curve")
	}

	return priv, nil
}
