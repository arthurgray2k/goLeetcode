package all

// Problem: 1846
// Title: Maximum Element After Decreasing and Rearranging
// Category: all
// Tags: all


func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	n := len(arr)
	count := make([]int, n+1)
	for _, v := range arr {
		count[min(v, n)]++
	}
	miss := 0
	for _, c := range count[1:] {
		if c == 0 {
			miss++
		} else {
			miss -= min(c-1, miss)
		}
	}
	return n - miss
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
