package all

// Problem: 1455
// Title: Check If a Word Occurs As a Prefix of Any Word in a Sentence
// Category: all
// Tags: all


import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	for i, v := range strings.Split(sentence, " ") {
		if strings.HasPrefix(v, searchWord) {
			return i + 1
		}
	}
	return -1
}
