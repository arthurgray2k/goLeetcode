package all

// Problem: 523
// Title: Continuous Subarray Sum
// Category: all
// Tags: all


func checkSubarraySum(nums []int, k int) bool {
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	for i, n := range nums {
		sum += n
		if r, ok := m[sum%k]; ok {
			if i-2 >= r {
				return true
			}
		} else {
			m[sum%k] = i
		}
	}
	return false
}
