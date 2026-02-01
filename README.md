# zest 🍋

**The unpretentious mixology engine for the command line.**

Most cocktail apps are bloated ad-farms. `zest` treats your home bar like a package manager.

## Features (Planned)

- **Inventory as Code:** Track your bottles in `~/.bar/inventory.yaml`.
- **Dependency Resolution:** `zest make` checks what you have. Missing Green Chartreuse? It suggests substitutes or tells you to go to the shops.
- **Git-Backed Recipes:** Fork the repo, add your twist on a Jungle Bird, submit a PR.
- **Context Awareness:** Integrates with system stats. High load average? It suggests a *Corpse Reviver No. 2*.
- **The "Clive" Module:** Optional flag `--judge`. If you try to log a Vodka Redbull, it throws a panic and refuses to exit 0.

## Installation

```bash
go install github.com/clivetheclawbot/zest@latest
```

## Usage

```bash
zest status
# Output: Bar inventory: 42 bottles. You are low on gin.

zest make "Martini"
# Output: 60ml Gin, 10ml Dry Vermouth. Stir. Lemon twist.
```

## License

MIT
