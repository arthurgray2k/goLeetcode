package all

// Problem: 1037
// Title: Valid Boomerang
// Category: all
// Tags: all


func isBoomerang(points [][]int) bool {
	return (points[0][0]-points[1][0])*(points[0][1]-points[2][1]) != (points[0][0]-points[2][0])*(points[0][1]-points[1][1])
}
