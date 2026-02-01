package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/clivetheclawbot/zest/internal/adapters/storage"
	"github.com/clivetheclawbot/zest/internal/core/services"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	// Wiring up the Hexagon
	// Adapters (Infrastructure)
	inventoryRepo := storage.NewFileInventoryRepository("inventory.yaml")
	recipeRepo := storage.NewMemoryRecipeRepository()

	// Services (Application Logic)
	bartender := services.NewBartenderService(inventoryRepo, recipeRepo)
	judge := services.NewJudgeService()
	shopper := services.NewShoppingService(inventoryRepo, recipeRepo)

	command := os.Args[1]

	switch command {
	case "status":
		stats, topShelf, err := bartender.GetStatus()
		if err != nil {
			fmt.Printf("Error loading inventory: %v\n", err)
			fmt.Println("Tip: Create 'inventory.yaml' or run 'zest init'")
			os.Exit(1)
		}
		fmt.Println("🍋 zest v0.0.1")
		fmt.Println(stats)
		
		// The Judge Speaks
		fmt.Printf("\n⚖️  Clive's Ruling: %s\n", judge.JudgeSession())

		if len(topShelf) > 0 {
			fmt.Println("\nTop Shelf:")
			for _, item := range topShelf {
				fmt.Printf("- %s\n", item)
			}
		}

	case "make":
		if len(os.Args) < 3 {
			fmt.Println("Error: What are we making? Usage: zest make <drink>")
			os.Exit(1)
		}
		drink := os.Args[2]
		recipe, missing, err := bartender.MakeDrink(drink)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Presentation Logic (View)
		fmt.Printf("🍸 %s\n", recipe.Name)
		fmt.Println(strings.Repeat("-", len(recipe.Name)+3))

		// Ingredients
		for _, ing := range recipe.Ingredients {
			target := ing.Name
			if target == "" {
				target = "[" + ing.Tag + "]"
			}
			fmt.Printf("• %.1f %s %s\n", ing.Amount, ing.Unit, target)
		}

		// Status Check
		if len(missing) > 0 {
			fmt.Println("\n❌ Missing Ingredients:")
			for _, m := range missing {
				fmt.Printf("  - %s\n", m)
			}
			fmt.Println("\nCannot make this drink. Go shopping.")
			os.Exit(1)
		}

		fmt.Println("\n✅ All ingredients available!")
		fmt.Println("\nInstructions:")
		for i, step := range recipe.Instructions {
			fmt.Printf("%d. %s\n", i+1, step)
		}

	case "shop":
		recs, err := shopper.GetRecommendations()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("🛍️  Zest Shopping List Optimizer")
		fmt.Println("--------------------------------")
		if len(recs) == 0 {
			fmt.Println("You have everything needed for all known recipes! (Or no recipes loaded)")
		} else {
			count := 0
			for _, r := range recs {
				// Filter out noisy low-value recs if list is long
				if count >= 5 {
					break
				}
				fmt.Printf("👉 Buy %-20s -> Unlocks %d recipes: %v\n", r.Ingredient, r.UnlockCount, r.UnlockList)
				count++
			}
		}

	case "help":
		printHelp()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("zest - the unpretentious mixology engine")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  zest status    Check your bar inventory")
	fmt.Println("  zest make      Attempt to mix a drink")
	fmt.Println("  zest shop      See what bottles to buy next")
	fmt.Println("  zest help      Show this help")
}
