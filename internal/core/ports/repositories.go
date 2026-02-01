package ports

import "github.com/clivetheclawbot/zest/internal/core/domain"

type InventoryRepository interface {
	Load() (*domain.Inventory, error)
	Save(*domain.Inventory) error
}

type RecipeRepository interface {
	LoadAll() ([]domain.Recipe, error)
	FindByName(name string) (*domain.Recipe, error)
}
