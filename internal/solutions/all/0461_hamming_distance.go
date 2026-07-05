package all

// Problem: 461
// Title: Hamming Distance
// Category: all
// Tags: all


func hammingDistance(x int, y int) int {
	distance := 0
	for xor := x ^ y; xor != 0; xor &= (xor - 1) {
		distance++
	}
	return distance
}
