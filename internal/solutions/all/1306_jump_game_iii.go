package all

// Problem: 1306
// Title: Jump Game III
// Category: all
// Tags: all


func canReach(arr []int, start int) bool {
	if start >= 0 && start < len(arr) && arr[start] < len(arr) {
		jump := arr[start]
		arr[start] += len(arr)
		return jump == 0 || canReach(arr, start+jump) || canReach(arr, start-jump)
	}
	return false
}
