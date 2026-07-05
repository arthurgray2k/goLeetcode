package all

// Problem: 1160
// Title: Find Words That Can Be Formed by Characters
// Category: all
// Tags: all


func countCharacters(words []string, chars string) int {
	count, res := make([]int, 26), 0
	for i := 0; i < len(chars); i++ {
		count[chars[i]-'a']++
	}
	for _, w := range words {
		if canBeFormed(w, count) {
			res += len(w)
		}
	}
	return res
}
func canBeFormed(w string, c []int) bool {
	count := make([]int, 26)
	for i := 0; i < len(w); i++ {
		count[w[i]-'a']++
		if count[w[i]-'a'] > c[w[i]-'a'] {
			return false
		}
	}
	return true
}
