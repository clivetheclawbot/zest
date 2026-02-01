package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	command := os.Args[1]

	switch command {
	case "status":
		inv, err := LoadInventory("inventory.yaml")
		if err != nil {
			fmt.Printf("Error loading inventory: %v\n", err)
			fmt.Println("Tip: Create 'inventory.yaml' or run 'zest init'")
			os.Exit(1)
		}
		fmt.Println("🍋 zest v0.0.1")
		fmt.Println(inv.Stats())
		// Basic listing for now
		if len(inv.Items) > 0 {
			fmt.Println("\nTop Shelf:")
			for i, item := range inv.Items {
				if i >= 5 {
					break
				}
				fmt.Printf("- %s (%s)\n", item.Name, item.Category)
			}
			if len(inv.Items) > 5 {
				fmt.Printf("...and %d more.\n", len(inv.Items)-5)
			}
		}
	case "make":
		if len(os.Args) < 3 {
			fmt.Println("Error: What are we making? Usage: zest make <drink>")
			os.Exit(1)
		}
		drink := os.Args[2]
		if drink == "vodka redbull" {
			panic("clive: absolutely not. runtime error: bad taste detected")
		}
		fmt.Printf("Searching memory for '%s'...\n", drink)
		fmt.Println("Error: Recipe database empty. Please implement inventory.yaml.")
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
	fmt.Println("  zest help      Show this help")
}
