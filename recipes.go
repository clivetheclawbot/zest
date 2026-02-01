package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Ingredient represents a component of a recipe.
type Ingredient struct {
	Name   string  `json:"name"`   // Specific name "London Dry Gin"
	Tag    string  `json:"tag"`    // Fallback tag "gin"
	Amount float64 `json:"amount"` // Amount in ml (or dashes)
	Unit   string  `json:"unit"`   // "ml", "dash", "splash"
}

// Recipe represents a cocktail recipe.
type Recipe struct {
	Name         string       `json:"name"`
	Ingredients  []Ingredient `json:"ingredients"`
	Instructions []string     `json:"instructions"`
	Tags         []string     `json:"tags"`
}

// RecipeBook holds all loaded recipes.
type RecipeBook struct {
	Recipes []Recipe
}

// LoadRecipes loads recipes from a JSON file (or hardcoded defaults for now).
func LoadRecipes(path string) (*RecipeBook, error) {
	// For MVP, we'll embed some defaults if file doesn't exist,
	// but let's try to load from disk first.
	var recipes []Recipe

	data, err := os.ReadFile(path)
	if err == nil {
		if err := json.Unmarshal(data, &recipes); err != nil {
			return nil, err
		}
	} else {
		// Fallback to embedded classics
		recipes = getEmbeddedRecipes()
	}

	return &RecipeBook{Recipes: recipes}, nil
}

func getEmbeddedRecipes() []Recipe {
	return []Recipe{
		{
			Name: "Bramble",
			Ingredients: []Ingredient{
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
			Ingredients: []Ingredient{
				{Tag: "light_rum", Amount: 60, Unit: "ml"},
				{Name: "Lime Super Juice", Tag: "lime_juice", Amount: 30, Unit: "ml"},
				{Name: "Sugar Syrup", Tag: "simple_syrup", Amount: 20, Unit: "ml"},
			},
			Instructions: []string{"Shake all ingredients with ice.", "Double strain into chilled coupe."},
			Tags:         []string{"classic", "rum", "sour"},
		},
		{
			Name: "Tommy's Margarita",
			Ingredients: []Ingredient{
				{Tag: "tequila", Amount: 60, Unit: "ml"},
				{Name: "Lime Super Juice", Tag: "lime_juice", Amount: 30, Unit: "ml"},
				{Name: "Agave Syrup", Tag: "agave_syrup", Amount: 15, Unit: "ml"}, // Missing in inventory, good test case
			},
			Instructions: []string{"Shake all ingredients with ice.", "Strain into rocks glass over ice."},
			Tags:         []string{"classic", "tequila", "sour"},
		},
	}
}

// FindRecipe searches for a recipe by name.
func (rb *RecipeBook) FindRecipe(name string) *Recipe {
	name = strings.ToLower(name)
	for _, r := range rb.Recipes {
		if strings.ToLower(r.Name) == name {
			return &r
		}
	}
	return nil
}

// PrintPretty prints the recipe to stdout.
func (r *Recipe) PrintPretty() {
	fmt.Printf("🍸 %s\n", r.Name)
	fmt.Println(strings.Repeat("-", len(r.Name)+3))
	for _, ing := range r.Ingredients {
		target := ing.Name
		if target == "" {
			target = "[" + ing.Tag + "]"
		}
		fmt.Printf("• %.1f %s %s\n", ing.Amount, ing.Unit, target)
	}
	fmt.Println("\nInstructions:")
	for i, step := range r.Instructions {
		fmt.Printf("%d. %s\n", i+1, step)
	}
}
