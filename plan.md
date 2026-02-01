# Zest Implementation Plan 🍋

## Phase 1: Core Mechanics (The "Backbar")
- [ ] **Inventory System:**
    - Schema for `~/.bar/inventory.yaml`.
    - `Inventory` struct and loader in Go.
    - CRUD commands: `zest bottle add <name>`, `zest bottle rm <name>`.
- [ ] **Recipe Schema:**
    - JSON/YAML format for drinks (ingredients, amounts, units, steps).
    - Hardcoded embedded library of classics (Martini, Negroni, Daiquiri).

## Phase 2: The "Bartender" (Logic)
- [ ] **Availability Engine:**
    - `CanMake(recipe, inventory) -> bool, missing_ingredients`.
    - Partial matching logic (substitutions TBD in Phase 3).
- [ ] **CLI Commands:**
    - `zest make <name>`: Checks inventory, prints instructions.
    - `zest list`: Shows what you can make *right now*.

## Phase 3: Advanced Features
- [ ] **Judge Module (`--judge`):**
    - Middleware to intercept requests.
    - Reject bad drinks with snarky errors.
    - Check system uptime/load for context-aware suggestions.
- [ ] **Shopping List:**
    - `zest shop`: Tells you what one bottle unlock the most new recipes.

## Phase 4: Polish
- [ ] **TUI Mode:** Bubbletea interface for browsing recipes.
- [ ] **Stats:** "You drink too much Gin" analytics.
