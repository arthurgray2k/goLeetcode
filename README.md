# goLeetcode

A repository containing hundreds of well-tested LeetCode problem solutions written in Go (Golang), complete with an interactive command-line utility for searching and exporting problems.

## Features

- **800+ Solutions**: Extensively tested and formatted solutions covering a wide variety of algorithms (Array, DP, Graph, Tree, etc.).
- **Categorized Structure**: Problems are stored inside `internal/solutions/` and can be easily sorted by tags and categories.
- **Interactive CLI**: Includes a lightning-fast heuristic search tool built purely with standard Go libraries.

## CLI Usage

The project includes an interactive CLI (`lc`) to quickly find, read, and export problems.

To build and run the CLI:
```bash
cd cmd/lc
go build -o lc main.go
./lc
```

### Searching and Viewing Problems

When you run the CLI, you will see an interactive prompt: `lc>`.

1. **Search heuristically**: Type any string (e.g., `array`, `Two Sum`, `hash-table`) to search. The CLI will list all matching problems and their problem numbers.
2. **Select and Show**: Type the **exact problem number** (e.g., `1`, `200`) and press Enter. The CLI will immediately print the full Go source code and description for that problem directly to your terminal.
3. **Export to File**: After printing a problem to the console, the CLI will ask:
   `Would you like to write this to a file in the current directory? (y/N):`
   Simply type `y` and press Enter. It will generate a `.go` file containing the problem code right in your current directory, making it easy to test or modify yourself!

## Testing

You can run the full test suite using standard Go tooling:
```bash
go test ./...
```
