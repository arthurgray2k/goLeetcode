package all

// Problem: 191
// Title: Number of 1 Bits
// Category: all
// Tags: all


import "math/bits"

// 解法一
func hammingWeight(num uint32) int {
	return bits.OnesCount(uint(num))
}

// 解法二
func hammingWeight1(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
