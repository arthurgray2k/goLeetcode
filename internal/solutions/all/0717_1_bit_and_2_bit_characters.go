package all

// Problem: 717
// Title: 1 bit and 2 bit Characters
// Category: all
// Tags: all


func isOneBitCharacter(bits []int) bool {
	var i int
	for i = 0; i < len(bits)-1; i++ {
		if bits[i] == 1 {
			i++
		}
	}
	return i == len(bits)-1
}
