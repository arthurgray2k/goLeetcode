package all

// Problem: 1137
// Title: N th Tribonacci Number
// Category: all
// Tags: all


func tribonacci(n int) int {
	if n < 2 {
		return n
	}
	trib, prev, prev2 := 1, 1, 0
	for n > 2 {
		trib, prev, prev2 = trib+prev+prev2, trib, prev
		n--
	}
	return trib
}
