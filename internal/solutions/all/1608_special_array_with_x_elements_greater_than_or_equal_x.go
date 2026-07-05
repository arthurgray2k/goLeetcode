package all

// Problem: 1608
// Title: Special Array With X Elements Greater Than or Equal X
// Category: all
// Tags: all


import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	x := len(nums)
	for _, num := range nums {
		if num >= x {
			return x
		}
		x--
		if num >= x {
			return -1
		}
	}
	return -1
}
