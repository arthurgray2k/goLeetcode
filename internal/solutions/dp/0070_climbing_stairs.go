package dp

// Problem: 70
// Title: Climbing Stairs
// Category: dynamic-programming
// Tags: math, dynamic-programming, memoization
//
// Description:
// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
