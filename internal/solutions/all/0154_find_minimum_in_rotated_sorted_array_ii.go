package all

// Problem: 154
// Title: Find Minimum in Rotated Sorted Array II
// Category: all
// Tags: all


func findMin154(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		if nums[low] < nums[high] {
			return nums[low]
		}
		mid := low + (high-low)>>1
		if nums[mid] > nums[low] {
			low = mid + 1
		} else if nums[mid] == nums[low] {
			low++
		} else {
			high = mid
		}
	}
	return nums[low]
}
