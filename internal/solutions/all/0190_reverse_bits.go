package all

// Problem: 190
// Title: Reverse Bits
// Category: all
// Tags: all


func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		res = res<<1 | num&1
		num >>= 1
	}
	return res
}
