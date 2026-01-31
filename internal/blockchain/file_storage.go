package blockchain

import (
	"encoding/json"
	"os"

	"github.com/irlan/quantumpay-go/internal/core"
)

type FileStorage struct {
	BasePath string
}

func NewFileStorage(base string) *FileStorage {
	return &FileStorage{BasePath: base}
}

func (fs *FileStorage) LoadBlockchain() ([]*core.Block, error) {
	path := fs.BasePath + "/chain.json"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []*core.Block{}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var blocks []*core.Block
	if err := json.Unmarshal(data, &blocks); err != nil {
		return nil, err
	}
	return blocks, nil
}

func (fs *FileStorage) SaveBlockchain(blocks []*core.Block) error {
	path := fs.BasePath + "/chain.json"

	data, err := json.MarshalIndent(blocks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
