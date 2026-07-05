package all

// Problem: 1446
// Title: Consecutive Characters
// Category: all
// Tags: all


func maxPower(s string) int {
	cur, cnt, ans := s[0], 1, 1
	for i := 1; i < len(s); i++ {
		if cur == s[i] {
			cnt++
		} else {
			if cnt > ans {
				ans = cnt
			}
			cur = s[i]
			cnt = 1
		}
	}
	if cnt > ans {
		ans = cnt
	}
	return ans
}
