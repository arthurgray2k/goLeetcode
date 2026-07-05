package all

// Problem: 136
// Title: Single Number
// Category: all
// Tags: all


func singleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}
