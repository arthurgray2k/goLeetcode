package all

// Problem: 973
// Title: K Closest Points to Origin
// Category: all
// Tags: all


import "sort"

// KClosest define
func KClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0]*points[i][0]+points[i][1]*points[i][1] <
			points[j][0]*points[j][0]+points[j][1]*points[j][1]
	})
	ans := make([][]int, K)
	for i := 0; i < K; i++ {
		ans[i] = points[i]
	}
	return ans
}
