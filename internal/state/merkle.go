package state

import (
	"crypto/sha256"
)

// MerkleTree struktur data sederhana
type MerkleTree struct {
	RootNode *MerkleNode
}

// MerkleNode adalah simpul pohon
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

// NewMerkleTree membuat pohon Merkle dari sekumpulan data (misal: ID Transaksi)
// FIX: Tidak lagi menerima WorldState, tapi raw data [][]byte
func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []*MerkleNode

	// Jika data ganjil, duplikasi data terakhir agar genap
	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	// Buat leaf nodes
	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, node)
	}

	// Bangun tree dari bawah ke atas
	for len(nodes) > 1 {
		var newLevel []*MerkleNode

		for i := 0; i < len(nodes); i += 2 {
			node1 := nodes[i]
			node2 := nodes[i+1]
			
			// Gabungkan hash kiri dan kanan
			combinedHash := append(node1.Data, node2.Data...)
			hash := sha256.Sum256(combinedHash)
			
			newNode := NewMerkleNode(node1, node2, hash[:])
			newLevel = append(newLevel, newNode)
		}

		nodes = newLevel
	}

	if len(nodes) == 0 {
		return &MerkleTree{RootNode: nil}
	}

	return &MerkleTree{RootNode: nodes[0]}
}

// NewMerkleNode membuat node baru
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	mNode := &MerkleNode{}

	if left == nil && right == nil {
		// Leaf Node: Hash data aslinya
		hash := sha256.Sum256(data)
		mNode.Data = hash[:]
	} else {
		// Parent Node: Data sudah di-hash di loop pembangun
		mNode.Data = data
	}

	mNode.Left = left
	mNode.Right = right

	return mNode
}
