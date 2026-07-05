package all

// Problem: 1657
// Title: Determine if Two Strings Are Close
// Category: all
// Tags: all


import (
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	freqCount1, freqCount2 := make([]int, 26), make([]int, 26)
	for _, c := range word1 {
		freqCount1[c-97]++
	}
	for _, c := range word2 {
		freqCount2[c-97]++
	}
	for i := 0; i < 26; i++ {
		if (freqCount1[i] == freqCount2[i]) ||
			(freqCount1[i] > 0 && freqCount2[i] > 0) {
			continue
		}
		return false
	}
	sort.Ints(freqCount1)
	sort.Ints(freqCount2)
	for i := 0; i < 26; i++ {
		if freqCount1[i] != freqCount2[i] {
			return false
		}
	}
	return true
}
