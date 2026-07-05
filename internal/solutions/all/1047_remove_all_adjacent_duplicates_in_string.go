package all

// Problem: 1047
// Title: Remove All Adjacent Duplicates In String
// Category: all
// Tags: all


func removeDuplicates1047(S string) string {
	stack := []rune{}
	for _, s := range S {
		if len(stack) == 0 || len(stack) > 0 && stack[len(stack)-1] != s {
			stack = append(stack, s)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
