package all

// Problem: 1470
// Title: Shuffle the Array
// Category: all
// Tags: all


func shuffle(nums []int, n int) []int {
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		result = append(result, nums[i])
		result = append(result, nums[n+i])
	}
	return result
}
