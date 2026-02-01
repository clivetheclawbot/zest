package domain

import "fmt"

// Item represents a single bottle or ingredient in the bar.
type Item struct {
	Name     string   `yaml:"name"`
	Category string   `yaml:"category"`
	Tags     []string `yaml:"tags,omitempty"`
}

// Inventory represents the collection of items available.
type Inventory struct {
	Items []Item `yaml:"items"`
}

// Has checks if the inventory contains an item by exact name or matching tag.
func (i *Inventory) Has(name, tag string) bool {
	for _, item := range i.Items {
		// 1. Exact Name Match
		if name != "" && item.Name == name {
			return true
		}
		// 2. Tag Match
		if tag != "" {
			for _, t := range item.Tags {
				if t == tag {
					return true
				}
			}
		}
	}
	return false
}

// Stats returns a summary of the inventory.
func (i *Inventory) Stats() string {
	counts := make(map[string]int)
	for _, item := range i.Items {
		counts[item.Category]++
	}

	return fmt.Sprintf("Inventory: %d items total (Spirits: %d, Liqueurs: %d, Mixers: %d)",
		len(i.Items), counts["spirit"], counts["liqueur"], counts["mixer"])
}
