package all

// Problem: 1716
// Title: Calculate Money in Leetcode Bank
// Category: all
// Tags: all


func totalMoney(n int) int {
	res := 0
	for tmp, count := 1, 7; n > 0; tmp, count = tmp+1, 7 {
		for m := tmp; n > 0 && count > 0; m++ {
			res += m
			n--
			count--
		}
	}
	return res
}
