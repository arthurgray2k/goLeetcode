package all

// Problem: 1005
// Title: Maximize Sum Of Array After K Negations
// Category: all
// Tags: all


import (
	"sort"
)

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	minIdx := 0
	for i := 0; i < K; i++ {
		A[minIdx] = -A[minIdx]
		if A[minIdx+1] < A[minIdx] {
			minIdx++
		}
	}
	sum := 0
	for _, a := range A {
		sum += a
	}
	return sum
}
