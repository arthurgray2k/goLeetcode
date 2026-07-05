package all

// Problem: 344
// Title: Reverse String
// Category: all
// Tags: all


func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
