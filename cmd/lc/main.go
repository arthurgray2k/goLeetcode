package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"goleetcode/internal/search"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Attempt to locate solutions dir
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting cwd: %v\n", err)
		return
	}
	
	solutionsDir := filepath.Join(cwd, "internal", "solutions")
	if _, err := os.Stat(solutionsDir); os.IsNotExist(err) {
		// If running from cmd/lc directory
		solutionsDir = filepath.Join("..", "..", "internal", "solutions")
	}

	problems, err := search.SearchProblems(solutionsDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error searching problems: %v\n", err)
		return
	}
	fmt.Printf("Loaded %d problems from %s\n", len(problems), solutionsDir)
	fmt.Println("Type a search term (e.g., 'array', 'Two Sum'), or a problem number.")
	fmt.Println("Type 'exit' to quit.")

	for {
		fmt.Print("lc> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		input = strings.TrimSpace(input)
		
		if input == "exit" || input == "quit" {
			break
		}
		if input == "" {
			continue
		}
		
		// If input is a number, print exact problem
		num, err := strconv.Atoi(input)
		if err == nil {
			found := false
			for _, p := range problems {
				if p.Number == num {
					printProblem(p, reader)
					found = true
					break
				}
			}
			if !found {
				fmt.Printf("Problem %d not found.\n", num)
			}
			continue
		}
		
		// Otherwise, search
		var matched []search.Problem
		for _, p := range problems {
			if search.Matches(p, input) {
				matched = append(matched, p)
			}
		}
		
		if len(matched) == 0 {
			fmt.Println("No matching problems found.")
		} else {
			fmt.Printf("Found %d problems:\n", len(matched))
			for _, p := range matched {
				fmt.Printf("[%d] %s (Category: %s, Tags: %s)\n", p.Number, p.Title, p.Category, strings.Join(p.Tags, ", "))
			}
			fmt.Println("Type the problem number to view it, or type a new search.")
		}
	}
}

func printProblem(p search.Problem, reader *bufio.Reader) {
	fmt.Printf("\n--- Problem %d: %s ---\n", p.Number, p.Title)
	fmt.Println(p.Content)
	fmt.Println("------------------------")
	
	fmt.Print("Would you like to write this to a file in the current directory? (y/N): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))
	if input == "y" || input == "yes" {
		outFilename := filepath.Base(p.Filepath)
		err := os.WriteFile(outFilename, []byte(p.Content), 0644)
		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
		} else {
			fmt.Printf("Successfully wrote %s\n", outFilename)
		}
	}
}
