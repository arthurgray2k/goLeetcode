package all

// Problem: 1816
// Title: Truncate Sentence
// Category: all
// Tags: all


func truncateSentence(s string, k int) string {
	end := 0
	for i := range s {
		if k > 0 && s[i] == ' ' {
			k--
		}
		if k == 0 {
			end = i
			break
		}
	}
	if end == 0 {
		return s
	}
	return s[:end]
}
