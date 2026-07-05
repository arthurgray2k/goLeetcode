package all

// Problem: 1437
// Title: Check If All 1s Are at Least Length K Places Away
// Category: all
// Tags: all


func kLengthApart(nums []int, k int) bool {
	prevIndex := -1
	for i, num := range nums {
		if num == 1 {
			if prevIndex != -1 && i-prevIndex-1 < k {
				return false
			}
			prevIndex = i
		}
	}
	return true
}
