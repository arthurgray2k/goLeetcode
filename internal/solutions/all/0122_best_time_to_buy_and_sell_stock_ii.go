package all

// Problem: 122
// Title: Best Time to Buy and Sell Stock II
// Category: all
// Tags: all


func maxProfit122(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}
