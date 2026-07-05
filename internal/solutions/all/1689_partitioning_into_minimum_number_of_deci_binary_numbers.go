package all

// Problem: 1689
// Title: Partitioning Into Minimum Number Of Deci Binary Numbers
// Category: all
// Tags: all


func minPartitions(n string) int {
	res := 0
	for i := 0; i < len(n); i++ {
		if int(n[i]-'0') > res {
			res = int(n[i] - '0')
		}
	}
	return res
}
