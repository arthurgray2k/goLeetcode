package all

// Problem: 540
// Title: Single Element in a Sorted Array
// Category: all
// Tags: all


func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			if nums[mid] == nums[mid-1] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return nums[left]
}
