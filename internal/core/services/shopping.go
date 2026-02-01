package services

import (
	"sort"

	"github.com/clivetheclawbot/zest/internal/core/ports"
)

type ShoppingRecommendation struct {
	Ingredient string
	UnlockCount int
	UnlockList  []string
}

type ShoppingService struct {
	inventoryRepo ports.InventoryRepository
	recipeRepo    ports.RecipeRepository
}

func NewShoppingService(inv ports.InventoryRepository, recipes ports.RecipeRepository) *ShoppingService {
	return &ShoppingService{
		inventoryRepo: inv,
		recipeRepo:    recipes,
	}
}

func (s *ShoppingService) GetRecommendations() ([]ShoppingRecommendation, error) {
	inv, err := s.inventoryRepo.Load()
	if err != nil {
		return nil, err
	}

	recipes, err := s.recipeRepo.LoadAll()
	if err != nil {
		return nil, err
	}

	// Map of Missing Ingredient -> List of Locked Recipes
	missingStats := make(map[string][]string)

	for _, recipe := range recipes {
		var missingForThisRecipe []string
		
		for _, ing := range recipe.Ingredients {
			if !inv.Has(ing.Name, ing.Tag) {
				key := ing.Name
				if key == "" {
					key = "[" + ing.Tag + "]"
				}
				missingForThisRecipe = append(missingForThisRecipe, key)
			}
		}

		// Logic: If we are missing EXACTLY ONE ingredient, that ingredient is high value.
		// If we are missing 5, buying one won't help much yet (simplification).
		// For now, let's just track everything we're missing.
		if len(missingForThisRecipe) > 0 {
			for _, missingItem := range missingForThisRecipe {
				missingStats[missingItem] = append(missingStats[missingItem], recipe.Name)
			}
		}
	}

	// Convert map to slice for sorting
	var recommendations []ShoppingRecommendation
	for ing, unlocks := range missingStats {
		recommendations = append(recommendations, ShoppingRecommendation{
			Ingredient: ing,
			UnlockCount: len(unlocks),
			UnlockList: unlocks,
		})
	}

	// Sort by UnlockCount descending
	sort.Slice(recommendations, func(i, j int) bool {
		return recommendations[i].UnlockCount > recommendations[j].UnlockCount
	})

	return recommendations, nil
}
