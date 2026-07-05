package all

// Problem: 1009
// Title: Complement of Base 10 Integer
// Category: all
// Tags: all


func bitwiseComplement(n int) int {
	mask := 1
	for mask < n {
		mask = (mask << 1) + 1
	}
	return mask ^ n
}
