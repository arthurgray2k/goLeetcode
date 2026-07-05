package all

// Problem: 709
// Title: To Lower Case
// Category: all
// Tags: all


func toLowerCase(s string) string {
	runes := []rune(s)
	diff := 'a' - 'A'
	for i := 0; i < len(s); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += diff
		}
	}
	return string(runes)
}
