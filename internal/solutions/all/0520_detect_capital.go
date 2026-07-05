package all

// Problem: 520
// Title: Detect Capital
// Category: all
// Tags: all


import "strings"

func detectCapitalUse(word string) bool {
	wLower := strings.ToLower(word)
	wUpper := strings.ToUpper(word)
	wCaptial := strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))
	if wCaptial == word || wLower == word || wUpper == word {
		return true
	}
	return false
}
