package all

// Problem: 1207
// Title: Unique Number of Occurrences
// Category: all
// Tags: all


func uniqueOccurrences(arr []int) bool {
	freq, m := map[int]int{}, map[int]bool{}
	for _, v := range arr {
		freq[v]++
	}
	for _, v := range freq {
		if _, ok := m[v]; !ok {
			m[v] = true
		} else {
			return false
		}
	}
	return true
}
