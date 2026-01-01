package core

import "encoding/binary"

// Uint64ToBytes converts uint64 to big-endian bytes
func Uint64ToBytes(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
