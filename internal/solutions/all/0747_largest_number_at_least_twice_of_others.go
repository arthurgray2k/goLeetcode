package all

// Problem: 747
// Title: Largest Number At Least Twice of Others
// Category: all
// Tags: all


func dominantIndex(nums []int) int {
	maxNum, flag, index := 0, false, 0
	for i, v := range nums {
		if v > maxNum {
			maxNum = v
			index = i
		}
	}
	for _, v := range nums {
		if v != maxNum && 2*v > maxNum {
			flag = true
		}
	}
	if flag {
		return -1
	}
	return index
}
