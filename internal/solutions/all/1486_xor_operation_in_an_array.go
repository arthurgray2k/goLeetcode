package all

// Problem: 1486
// Title: XOR Operation in an Array
// Category: all
// Tags: all


func xorOperation(n int, start int) int {
	res := 0
	for i := 0; i < n; i++ {
		res ^= start + 2*i
	}
	return res
}
