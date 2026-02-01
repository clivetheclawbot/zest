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
			{
				Name: "Negroni",
				Ingredients: []domain.Ingredient{
					{Tag: "gin", Amount: 30, Unit: "ml"},
					{Name: "Campari", Tag: "campari", Amount: 30, Unit: "ml"},
					{Name: "Sweet Vermouth", Tag: "sweet_vermouth", Amount: 30, Unit: "ml"},
				},
				Instructions: []string{"Stir all ingredients with ice.", "Strain into rocks glass over ice.", "Garnish with orange peel."},
				Tags:         []string{"classic", "gin", "bitter", "stirred"},
			},
			{
				Name: "Old Fashioned",
				Ingredients: []domain.Ingredient{
					{Tag: "whiskey", Amount: 60, Unit: "ml"},
					{Name: "Sugar Syrup", Tag: "simple_syrup", Amount: 10, Unit: "ml"},
					{Name: "Angostura Bitters", Tag: "angostura_bitters", Amount: 2, Unit: "dash"},
				},
				Instructions: []string{"Stir sugar and bitters with a splash of water.", "Add whiskey and ice.", "Stir until chilled and diluted.", "Garnish with orange peel."},
				Tags:         []string{"classic", "whiskey", "stirred"},
			},
			{
				Name: "Corpse Reviver No. 2",
				Ingredients: []domain.Ingredient{
					{Tag: "gin", Amount: 25, Unit: "ml"},
					{Name: "Lillet Blanc", Tag: "lillet_blanc", Amount: 25, Unit: "ml"},
					{Name: "Triple Sec", Tag: "orange_liqueur", Amount: 25, Unit: "ml"},
					{Name: "Lemon Super Juice", Tag: "lemon_juice", Amount: 25, Unit: "ml"},
					{Name: "Absinthe", Tag: "absinthe", Amount: 1, Unit: "dash"},
				},
				Instructions: []string{"Rinse chilled coupe with absinthe.", "Shake other ingredients with ice.", "Fine strain into glass.", "Garnish with lemon twist."},
				Tags:         []string{"classic", "gin", "sour", "strong"},
			},
			{
				Name: "Mai Tai",
				Ingredients: []domain.Ingredient{
					{Tag: "aged_rum", Amount: 60, Unit: "ml"},
					{Name: "Lime Super Juice", Tag: "lime_juice", Amount: 30, Unit: "ml"},
					{Name: "Triple Sec", Tag: "orange_liqueur", Amount: 15, Unit: "ml"},
					{Name: "Orgeat", Tag: "orgeat", Amount: 15, Unit: "ml"},
					{Name: "Sugar Syrup", Tag: "simple_syrup", Amount: 7.5, Unit: "ml"},
				},
				Instructions: []string{"Shake all ingredients with ice.", "Strain into rocks glass over crushed ice.", "Garnish with mint sprig and spent lime shell."},
				Tags:         []string{"tiki", "rum", "sour"},
			},
			{
				Name: "Whiskey Sour",
				Ingredients: []domain.Ingredient{
					{Tag: "whiskey", Amount: 60, Unit: "ml"},
					{Name: "Lemon Super Juice", Tag: "lemon_juice", Amount: 30, Unit: "ml"},
					{Name: "Sugar Syrup", Tag: "simple_syrup", Amount: 15, Unit: "ml"},
					{Name: "Egg White", Tag: "egg_white", Amount: 15, Unit: "ml"},
				},
				Instructions: []string{"Dry shake all ingredients.", "Add ice and shake again.", "Strain into chilled glass.", "Garnish with Angostura drops."},
				Tags:         []string{"classic", "whiskey", "sour"},
			},
			{
				Name: "Espresso Martini",
				Ingredients: []domain.Ingredient{
					{Tag: "vodka", Amount: 50, Unit: "ml"},
					{Name: "Coffee Liqueur", Tag: "coffee_liqueur", Amount: 25, Unit: "ml"},
					{Name: "Espresso", Tag: "espresso", Amount: 25, Unit: "ml"},
					{Name: "Sugar Syrup", Tag: "simple_syrup", Amount: 5, Unit: "ml"},
				},
				Instructions: []string{"Shake vigorously with ice.", "Strain into chilled coupe.", "Garnish with 3 coffee beans."},
				Tags:         []string{"modern_classic", "vodka", "coffee"},
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
