package all

// Problem: 1614
// Title: Maximum Nesting Depth of the Parentheses
// Category: all
// Tags: all


func maxDepth(s string) int {
	res, cur := 0, 0
	for _, c := range s {
		if c == '(' {
			cur++
			res = max(res, cur)
		} else if c == ')' {
			cur--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
