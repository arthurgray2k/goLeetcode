package all

// Problem: 1299
// Title: Replace Elements with Greatest Element on Right Side
// Category: all
// Tags: all


func replaceElements(arr []int) []int {
	j, temp := -1, 0
	for i := len(arr) - 1; i >= 0; i-- {
		temp = arr[i]
		arr[i] = j
		j = max(j, temp)
	}
	return arr
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
