package all

// Problem: 922
// Title: Sort Array By Parity II
// Category: all
// Tags: all


func sortArrayByParityII(A []int) []int {
	if len(A) == 0 || len(A)%2 != 0 {
		return []int{}
	}
	res := make([]int, len(A))
	oddIndex := 1
	evenIndex := 0
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			res[evenIndex] = A[i]
			evenIndex += 2
		} else {
			res[oddIndex] = A[i]
			oddIndex += 2
		}
	}
	return res
}
