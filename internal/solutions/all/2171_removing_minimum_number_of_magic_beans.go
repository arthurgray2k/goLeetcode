package all

// Problem: 2171
// Title: Removing Minimum Number of Magic Beans
// Category: all
// Tags: all


import "sort"

func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)
	sum, mx := 0, 0
	for i, v := range beans {
		sum += v
		mx = max(mx, (len(beans)-i)*v)
	}
	return int64(sum - mx)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
