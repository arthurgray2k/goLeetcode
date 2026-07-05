package all

// Problem: 1018
// Title: Binary Prefix Divisible By 5
// Category: all
// Tags: all


func prefixesDivBy5(a []int) []bool {
	res, num := make([]bool, len(a)), 0
	for i, v := range a {
		num = (num<<1 | v) % 5
		res[i] = num == 0
	}
	return res
}
