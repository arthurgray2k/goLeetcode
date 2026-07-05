package all

// Problem: 1732
// Title: Find the Highest Altitude
// Category: all
// Tags: all


func largestAltitude(gain []int) int {
	max, height := 0, 0
	for _, g := range gain {
		height += g
		if height > max {
			max = height
		}
	}
	return max
}
