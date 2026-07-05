package all

// Problem: 775
// Title: Global and Local Inversions
// Category: all
// Tags: all


func isIdealPermutation(A []int) bool {
	for i := range A {
		if abs(A[i]-i) > 1 {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
