package all

// Problem: 1287
// Title: Element Appearing More Than In Sorted Array
// Category: all
// Tags: all


func findSpecialInteger(arr []int) int {
	n := len(arr)
	for i := 0; i < n-n/4; i++ {
		if arr[i] == arr[i+n/4] {
			return arr[i]
		}
	}
	return -1
}
