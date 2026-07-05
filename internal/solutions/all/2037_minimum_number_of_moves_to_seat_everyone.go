package all

// Problem: 2037
// Title: Minimum Number of Moves to Seat Everyone
// Category: all
// Tags: all


import "sort"

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	n := len(students)
	moves := 0
	for i := 0; i < n; i++ {
		moves += abs(seats[i], students[i])
	}
	return moves
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
