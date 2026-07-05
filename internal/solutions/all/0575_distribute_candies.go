package all

// Problem: 575
// Title: Distribute Candies
// Category: all
// Tags: all


func distributeCandies(candies []int) int {
	n, m := len(candies), make(map[int]struct{}, len(candies))
	for _, candy := range candies {
		m[candy] = struct{}{}
	}
	res := len(m)
	if n/2 < res {
		return n / 2
	}
	return res
}
