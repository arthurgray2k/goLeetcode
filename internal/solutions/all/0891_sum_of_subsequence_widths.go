package all

// Problem: 891
// Title: Sum of Subsequence Widths
// Category: all
// Tags: all


import (
	"sort"
)

func sumSubseqWidths(A []int) int {
	sort.Ints(A)
	res, mod, n, p := 0, 1000000007, len(A), 1
	for i := 0; i < n; i++ {
		res = (res + (A[i]-A[n-1-i])*p) % mod
		p = (p << 1) % mod
	}
	return res
}
