package all

// Problem: 1480
// Title: Running Sum of 1d Array
// Category: all
// Tags: all


func runningSum(nums []int) []int {
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	for i := 1; i <= len(nums); i++ {
		dp[i] = dp[i-1] + nums[i-1]
	}
	return dp[1:]
}
