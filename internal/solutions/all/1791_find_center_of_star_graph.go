package all

// Problem: 1791
// Title: Find Center of Star Graph
// Category: all
// Tags: all


func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
}
