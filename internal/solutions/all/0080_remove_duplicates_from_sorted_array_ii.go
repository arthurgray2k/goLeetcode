package all

// Problem: 80
// Title: Remove Duplicates from Sorted Array II
// Category: all
// Tags: all


func removeDuplicates(nums []int) int {
	slow := 0
	for fast, v := range nums {
		if fast < 2 || nums[slow-2] != v {
			nums[slow] = v
			slow++
		}
	}
	return slow
}
