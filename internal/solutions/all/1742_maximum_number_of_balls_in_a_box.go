package all

// Problem: 1742
// Title: Maximum Number of Balls in a Box
// Category: all
// Tags: all


func countBalls(lowLimit int, highLimit int) int {
	buckets, maxBall := [46]int{}, 0
	for i := lowLimit; i <= highLimit; i++ {
		t := 0
		for j := i; j > 0; {
			t += j % 10
			j = j / 10
		}
		buckets[t]++
		if buckets[t] > maxBall {
			maxBall = buckets[t]
		}
	}
	return maxBall
}
