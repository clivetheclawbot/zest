package storage

import (
	"os"

	"github.com/clivetheclawbot/zest/internal/core/domain"
	"gopkg.in/yaml.v3"
)

type FileInventoryRepository struct {
	Path string
}

func NewFileInventoryRepository(path string) *FileInventoryRepository {
	return &FileInventoryRepository{Path: path}
}

func (r *FileInventoryRepository) Load() (*domain.Inventory, error) {
	data, err := os.ReadFile(r.Path)
	if err != nil {
		return nil, err
	}

	var inv domain.Inventory
	err = yaml.Unmarshal(data, &inv)
	if err != nil {
		return nil, err
	}

	return &inv, nil
}

func (r *FileInventoryRepository) Save(inv *domain.Inventory) error {
	data, err := yaml.Marshal(inv)
	if err != nil {
		return err
	}
	return os.WriteFile(r.Path, data, 0644)
}
