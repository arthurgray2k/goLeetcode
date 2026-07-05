package search

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Problem struct {
	Number   int
	Title    string
	Category string
	Tags     []string
	Filepath string
	Content  string
}

// SearchProblems parses files inside the given directory and returns a slice of Problems
func SearchProblems(rootDir string) ([]Problem, error) {
	var problems []Problem
	
	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || strings.HasSuffix(d.Name(), "_test.go") || !strings.HasSuffix(d.Name(), ".go") {
			return nil
		}
		
		problem, err := parseFile(path)
		if err == nil && problem.Number != 0 {
			problems = append(problems, problem)
		}
		return nil
	})
	
	return problems, err
}

func parseFile(path string) (Problem, error) {
	file, err := os.Open(path)
	if err != nil {
		return Problem{}, err
	}
	defer file.Close()
	
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return Problem{}, err
	}
	
	p := Problem{
		Filepath: path,
		Content:  string(contentBytes),
	}
	
	scanner := bufio.NewScanner(file)
	
	reNumber := regexp.MustCompile(`^//\s*Problem:\s*(\d+)`)
	reTitle := regexp.MustCompile(`^//\s*Title:\s*(.+)`)
	reCategory := regexp.MustCompile(`^//\s*Category:\s*(.+)`)
	reTags := regexp.MustCompile(`^//\s*Tags:\s*(.+)`)
	
	for scanner.Scan() {
		line := scanner.Text()
		
		if m := reNumber.FindStringSubmatch(line); m != nil {
			p.Number, _ = strconv.Atoi(m[1])
		} else if m := reTitle.FindStringSubmatch(line); m != nil {
			p.Title = strings.TrimSpace(m[1])
		} else if m := reCategory.FindStringSubmatch(line); m != nil {
			p.Category = strings.TrimSpace(m[1])
		} else if m := reTags.FindStringSubmatch(line); m != nil {
			tagsSplit := strings.Split(m[1], ",")
			for _, tag := range tagsSplit {
				p.Tags = append(p.Tags, strings.TrimSpace(tag))
			}
		}
	}
	
	return p, scanner.Err()
}

func Matches(p Problem, query string) bool {
    query = strings.ToLower(query)
    if strconv.Itoa(p.Number) == query {
        return true
    }
    if strings.Contains(strings.ToLower(p.Title), query) {
        return true
    }
    if strings.Contains(strings.ToLower(p.Category), query) {
        return true
    }
    for _, tag := range p.Tags {
        if strings.Contains(strings.ToLower(tag), query) {
            return true
        }
    }
    return false
}
