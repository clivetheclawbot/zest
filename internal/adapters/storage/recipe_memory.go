package storage

import (
	"strings"

	"github.com/clivetheclawbot/zest/internal/core/domain"
)

type MemoryRecipeRepository struct {
	recipes []domain.Recipe
}

func NewMemoryRecipeRepository() *MemoryRecipeRepository {
	// Seeding with the embedded list for now
	return &MemoryRecipeRepository{
		recipes: []domain.Recipe{
			{
				Name: "Bramble",
				Ingredients: []domain.Ingredient{
					{Tag: "gin", Amount: 50, Unit: "ml"},
					{Name: "Lemon Super Juice", Tag: "lemon_juice", Amount: 25, Unit: "ml"},
					{Name: "Sugar Syrup", Tag: "simple_syrup", Amount: 12.5, Unit: "ml"},
					{Name: "Creme de Mure", Tag: "blackberry_liqueur", Amount: 15, Unit: "ml"},
				},
				Instructions: []string{"Shake gin, lemon, sugar with ice.", "Strain into glass with crushed ice.", "Drizzle Creme de Mure over top.", "Garnish with blackberry and lemon."},
				Tags:         []string{"classic", "gin", "fruity"},
			},
			{
				Name: "Daiquiri",
				Ingredients: []domain.Ingredient{
					{Tag: "light_rum", Amount: 60, Unit: "ml"},
					{Name: "Lime Super Juice", Tag: "lime_juice", Amount: 30, Unit: "ml"},
					{Name: "Sugar Syrup", Tag: "simple_syrup", Amount: 20, Unit: "ml"},
				},
				Instructions: []string{"Shake all ingredients with ice.", "Double strain into chilled coupe."},
				Tags:         []string{"classic", "rum", "sour"},
			},
			{
				Name: "Tommy's Margarita",
				Ingredients: []domain.Ingredient{
					{Tag: "tequila", Amount: 60, Unit: "ml"},
					{Name: "Lime Super Juice", Tag: "lime_juice", Amount: 30, Unit: "ml"},
					{Name: "Agave Syrup", Tag: "agave_syrup", Amount: 15, Unit: "ml"},
				},
				Instructions: []string{"Shake all ingredients with ice.", "Strain into rocks glass over ice."},
				Tags:         []string{"classic", "tequila", "sour"},
			},
		},
	}
}

func (r *MemoryRecipeRepository) LoadAll() ([]domain.Recipe, error) {
	return r.recipes, nil
}

func (r *MemoryRecipeRepository) FindByName(name string) (*domain.Recipe, error) {
	search := strings.ToLower(name)
	for _, recipe := range r.recipes {
		if strings.ToLower(recipe.Name) == search {
			return &recipe, nil
		}
	}
	return nil, nil // Not found
}
