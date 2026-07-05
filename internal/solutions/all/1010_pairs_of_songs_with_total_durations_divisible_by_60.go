package all

// Problem: 1010
// Title: Pairs of Songs With Total Durations Divisible by 60
// Category: all
// Tags: all


func numPairsDivisibleBy60(time []int) int {
	counts := make([]int, 60)
	for _, v := range time {
		v %= 60
		counts[v]++
	}
	res := 0
	for i := 1; i < len(counts)/2; i++ {
		res += counts[i] * counts[60-i]
	}
	res += (counts[0] * (counts[0] - 1)) / 2
	res += (counts[30] * (counts[30] - 1)) / 2
	return res
}
