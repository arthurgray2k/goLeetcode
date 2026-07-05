package all

// Problem: 1389
// Title: Create Target Array in the Given Order
// Category: all
// Tags: all


func createTargetArray(nums []int, index []int) []int {
	result := make([]int, len(nums))
	for i, pos := range index {
		copy(result[pos+1:i+1], result[pos:i])
		result[pos] = nums[i]
	}
	return result
}
