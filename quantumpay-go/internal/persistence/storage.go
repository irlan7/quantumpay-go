package persistence

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/irlan/quantumpay-go/internal/block"
	"github.com/irlan/quantumpay-go/internal/state"
)

type Storage struct {
	BasePath string
}

func NewStorage(basePath string) *Storage {
	return &Storage{
		BasePath: basePath,
	}
}

/* ---------------- BLOCK ---------------- */

// SaveBlock menyimpan block ke disk (append-only)
func (s *Storage) SaveBlock(b *block.Block) error {
	blocksDir := filepath.Join(s.BasePath, "blocks")
	err := os.MkdirAll(blocksDir, 0755)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("block_%d.json", b.Header.Height)
	path := filepath.Join(blocksDir, filename)

	data, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return err
	}

	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0644); err != nil {
		return err
	}

	return os.Rename(tmp, path)
}

// LoadBlocks memuat seluruh block dari disk (berurutan)
func (s *Storage) LoadBlocks() ([]*block.Block, error) {
	blocksDir := filepath.Join(s.BasePath, "blocks")
	files, err := os.ReadDir(blocksDir)
	if err != nil {
		return nil, err
	}

	var blocks []*block.Block

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		data, err := os.ReadFile(filepath.Join(blocksDir, f.Name()))
		if err != nil {
			return nil, err
		}

		var b block.Block
		if err := json.Unmarshal(data, &b); err != nil {
			return nil, err
		}

		blocks = append(blocks, &b)
	}

	return blocks, nil
}

/* ---------------- STATE ---------------- */

// SaveState menyimpan world state snapshot
func (s *Storage) SaveState(ws *state.WorldState) error {
	path := filepath.Join(s.BasePath, "state.json")

	data, err := json.MarshalIndent(ws, "", "  ")
	if err != nil {
		return err
	}

	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0644); err != nil {
		return err
	}

	return os.Rename(tmp, path)
}

// LoadState memuat world state dari disk
func (s *Storage) LoadState() (*state.WorldState, error) {
	path := filepath.Join(s.BasePath, "state.json")

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var ws state.WorldState
	if err := json.Unmarshal(data, &ws); err != nil {
		return nil, err
	}

	return &ws, nil
}
