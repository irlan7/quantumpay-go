package persistence

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/irlan/quantumpay-go/internal/core"
)

type FileStorage struct {
	BaseDir string
}

func NewFileStorage(baseDir string) *FileStorage {
	_ = os.MkdirAll(baseDir, 0755)
	return &FileStorage{BaseDir: baseDir}
}

func (fs *FileStorage) blocksFile() string {
	return filepath.Join(fs.BaseDir, "blocks.json")
}

func (fs *FileStorage) SaveBlocks(blocks []*core.Block) error {
	f, err := os.Create(fs.blocksFile())
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(blocks)
}

func (fs *FileStorage) LoadBlocks() ([]*core.Block, error) {
	path := fs.blocksFile()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []*core.Block{}, nil
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var blocks []*core.Block
	err = json.NewDecoder(f).Decode(&blocks)
	return blocks, err
}
