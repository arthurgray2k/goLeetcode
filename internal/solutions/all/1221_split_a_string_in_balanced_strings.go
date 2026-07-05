package all

// Problem: 1221
// Title: Split a String in Balanced Strings
// Category: all
// Tags: all


func balancedStringSplit(s string) int {
	count, res := 0, 0
	for _, r := range s {
		if r == 'R' {
			count++
		} else {
			count--
		}
		if count == 0 {
			res++
		}
	}
	return res
}
