package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	srcDir := `D:\golang_toolshed\temp_leetcode_go\leetcode`
	destDir := `D:\golang_toolshed\goLeetcode\internal\solutions\all`

	if err := os.MkdirAll(destDir, 0755); err != nil {
		fmt.Printf("Failed to create dest dir: %v\n", err)
		return
	}

	entries, err := os.ReadDir(srcDir)
	if err != nil {
		fmt.Printf("Failed to read src dir: %v\n", err)
		return
	}

	reDirName := regexp.MustCompile(`^(\d+)\.(.+)$`)
	
	count := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		
		match := reDirName.FindStringSubmatch(entry.Name())
		if match == nil {
			continue
		}
		
		numStr := match[1]
		titleStr := match[2]
		
		num, err := strconv.Atoi(numStr)
		if err != nil {
			continue
		}
		
		// find the .go file in this directory
		problemDir := filepath.Join(srcDir, entry.Name())
		files, err := os.ReadDir(problemDir)
		if err != nil {
			continue
		}
		
		var goFile string
		var testFile string
		for _, f := range files {
			if strings.HasSuffix(f.Name(), "_test.go") {
				testFile = f.Name()
			} else if strings.HasSuffix(f.Name(), ".go") {
				goFile = f.Name()
			}
		}
		
		if goFile == "" {
			continue
		}
		
		// generate snake case name for filename
		cleanTitle := strings.ReplaceAll(titleStr, "-", "_")
		cleanTitle = strings.ToLower(cleanTitle)
		
		destFile := filepath.Join(destDir, fmt.Sprintf("%04d_%s.go", num, cleanTitle))
		destTestFile := filepath.Join(destDir, fmt.Sprintf("%04d_%s_test.go", num, cleanTitle))
		
		// Read and modify the source file
		srcContent, err := os.ReadFile(filepath.Join(problemDir, goFile))
		if err != nil {
			continue
		}
		
		newContent := rewriteContent(string(srcContent), num, strings.ReplaceAll(titleStr, "-", " "))
		os.WriteFile(destFile, []byte(newContent), 0644)
		
		// Read and modify test file if exists
		if testFile != "" {
			testContent, err := os.ReadFile(filepath.Join(problemDir, testFile))
			if err == nil {
				newTestContent := rewriteTestContent(string(testContent))
				os.WriteFile(destTestFile, []byte(newTestContent), 0644)
			}
		}
		
		count++
	}
	
	fmt.Printf("Successfully imported %d problems!\n", count)
}

func rewriteContent(content string, num int, title string) string {
	lines := strings.Split(content, "\n")
	var out []string
	
	// Add header
	out = append(out, "package all")
	out = append(out, "")
	out = append(out, fmt.Sprintf("// Problem: %d", num))
	out = append(out, fmt.Sprintf("// Title: %s", title))
	out = append(out, "// Category: all")
	out = append(out, "// Tags: all")
	out = append(out, "")
	
	for _, line := range lines {
		if strings.HasPrefix(line, "package ") {
			continue
		}
		out = append(out, line)
	}
	
	return strings.Join(out, "\n")
}

func rewriteTestContent(content string) string {
	lines := strings.Split(content, "\n")
	var out []string
	out = append(out, "package all")
	for _, line := range lines {
		if strings.HasPrefix(line, "package ") {
			continue
		}
		out = append(out, line)
	}
	return strings.Join(out, "\n")
}
