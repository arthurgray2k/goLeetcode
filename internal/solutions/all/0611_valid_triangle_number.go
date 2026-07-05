package all

// Problem: 611
// Title: Valid Triangle Number
// Category: all
// Tags: all


import "sort"

func triangleNumber(nums []int) int {
	res := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		k := i + 2
		for j := i + 1; j < len(nums)-1 && nums[i] != 0; j++ {
			for k < len(nums) && nums[i]+nums[j] > nums[k] {
				k++
			}
			res += k - j - 1
		}
	}
	return res
}
