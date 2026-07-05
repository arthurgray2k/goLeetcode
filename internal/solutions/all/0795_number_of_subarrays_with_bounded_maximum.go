package all

// Problem: 795
// Title: Number of Subarrays with Bounded Maximum
// Category: all
// Tags: all


func numSubarrayBoundedMax(nums []int, left int, right int) int {
	return getAnswerPerBound(nums, right) - getAnswerPerBound(nums, left-1)
}

func getAnswerPerBound(nums []int, bound int) int {
	res, count := 0, 0
	for _, num := range nums {
		if num <= bound {
			count++
		} else {
			count = 0
		}
		res += count
	}
	return res
}
