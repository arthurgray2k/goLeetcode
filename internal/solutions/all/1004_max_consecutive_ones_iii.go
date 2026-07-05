package all

// Problem: 1004
// Title: Max Consecutive Ones III
// Category: all
// Tags: all


func longestOnes(A []int, K int) int {
	res, left, right := 0, 0, 0
	for left < len(A) {
		if right < len(A) && ((A[right] == 0 && K > 0) || A[right] == 1) {
			if A[right] == 0 {
				K--
			}
			right++
		} else {
			if K == 0 || (right == len(A) && K > 0) {
				res = max(res, right-left)
			}
			if A[left] == 0 {
				K++
			}
			left++
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
