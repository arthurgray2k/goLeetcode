package all

// Problem: 991
// Title: Broken Calculator
// Category: all
// Tags: all


func brokenCalc(X int, Y int) int {
	res := 0
	for Y > X {
		res++
		if Y&1 == 1 {
			Y++
		} else {
			Y /= 2
		}
	}
	return res + X - Y
}
