package all

// Problem: 1576
// Title: Replace All s to Avoid Consecutive Repeating Characters
// Category: all
// Tags: all


func modifyString(s string) string {
	res := []byte(s)
	for i, ch := range res {
		if ch == '?' {
			for b := byte('a'); b <= 'z'; b++ {
				if !(i > 0 && res[i-1] == b || i < len(res)-1 && res[i+1] == b) {
					res[i] = b
					break
				}
			}
		}
	}
	return string(res)
}
