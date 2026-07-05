package all

// Problem: 1295
// Title: Find Numbers with Even Number of Digits
// Category: all
// Tags: all


import "strconv"

func findNumbers(nums []int) int {
	res := 0
	for _, n := range nums {
		res += 1 - len(strconv.Itoa(n))%2
	}
	return res
}
