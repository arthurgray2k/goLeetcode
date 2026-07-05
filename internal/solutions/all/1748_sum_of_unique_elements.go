package all

// Problem: 1748
// Title: Sum of Unique Elements
// Category: all
// Tags: all


func sumOfUnique(nums []int) int {
	freq, res := make(map[int]int), 0
	for _, v := range nums {
		if _, ok := freq[v]; !ok {
			freq[v] = 0
		}
		freq[v]++
	}
	for k, v := range freq {
		if v == 1 {
			res += k
		}
	}
	return res
}
