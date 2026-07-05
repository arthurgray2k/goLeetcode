package all

// Problem: 810
// Title: Chalkboard XOR Game
// Category: all
// Tags: all


func xorGame(nums []int) bool {
	if len(nums)%2 == 0 {
		return true
	}
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return xor == 0
}
