package all

// Problem: 1700
// Title: Number of Students Unable to Eat Lunch
// Category: all
// Tags: all


func countStudents(students []int, sandwiches []int) int {
	tmp, n, i := [2]int{}, len(students), 0
	for _, v := range students {
		tmp[v]++
	}
	for i < n && tmp[sandwiches[i]] > 0 {
		tmp[sandwiches[i]]--
		i++
	}
	return n - i
}
