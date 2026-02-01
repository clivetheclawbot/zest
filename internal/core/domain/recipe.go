package domain

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
