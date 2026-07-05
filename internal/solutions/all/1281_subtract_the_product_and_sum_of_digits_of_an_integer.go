package all

// Problem: 1281
// Title: Subtract the Product and Sum of Digits of an Integer
// Category: all
// Tags: all


func subtractProductAndSum(n int) int {
	sum, product := 0, 1
	for ; n > 0; n /= 10 {
		sum += n % 10
		product *= n % 10
	}
	return product - sum
}
