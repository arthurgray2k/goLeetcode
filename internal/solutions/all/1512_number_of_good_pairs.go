package all

// Problem: 1512
// Title: Number of Good Pairs
// Category: all
// Tags: all


func numIdenticalPairs(nums []int) int {
	total := 0
	for x := 0; x < len(nums); x++ {
		for y := x + 1; y < len(nums); y++ {
			if nums[x] == nums[y] {
				total++
			}
		}
	}
	return total
}
