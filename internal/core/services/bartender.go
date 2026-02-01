package services

import (
	"fmt"
	"strings"

	"github.com/clivetheclawbot/zest/internal/core/domain"
	"github.com/clivetheclawbot/zest/internal/core/ports"
)

type BartenderService struct {
	inventoryRepo ports.InventoryRepository
	recipeRepo    ports.RecipeRepository
}

func NewBartenderService(inv ports.InventoryRepository, recipes ports.RecipeRepository) *BartenderService {
	return &BartenderService{
		inventoryRepo: inv,
		recipeRepo:    recipes,
	}
}

func (s *BartenderService) GetStatus() (string, []string, error) {
	inv, err := s.inventoryRepo.Load()
	if err != nil {
		return "", nil, err
	}

	topShelf := []string{}
	for i, item := range inv.Items {
		if i >= 5 {
			break
		}
		topShelf = append(topShelf, fmt.Sprintf("%s (%s)", item.Name, item.Category))
	}
	return inv.Stats(), topShelf, nil
}

func (s *BartenderService) MakeDrink(name string) (*domain.Recipe, error) {
	if strings.ToLower(name) == "vodka redbull" {
		panic("clive: absolutely not. runtime error: bad taste detected")
	}

	recipe, err := s.recipeRepo.FindByName(name)
	if err != nil {
		return nil, err
	}
	if recipe == nil {
		return nil, fmt.Errorf("I don't know how to make a '%s'", name)
	}

	// TODO: Availability check will go here

	return recipe, nil
}
