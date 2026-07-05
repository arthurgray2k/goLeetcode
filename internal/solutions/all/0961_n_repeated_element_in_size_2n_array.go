package all

// Problem: 961
// Title: N Repeated Element in Size 2N Array
// Category: all
// Tags: all


func repeatedNTimes(A []int) int {
	kv := make(map[int]struct{})
	for _, val := range A {
		if _, ok := kv[val]; ok {
			return val
		}
		kv[val] = struct{}{}
	}
	return 0
}
