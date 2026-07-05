package all

// Problem: 1678
// Title: Goal Parser Interpretation
// Category: all
// Tags: all


func interpret(command string) string {
	if command == "" {
		return ""
	}
	res := ""
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			res += "G"
		} else {
			if command[i] == '(' && command[i+1] == 'a' {
				res += "al"
				i += 3
			} else {
				res += "o"
				i++
			}
		}
	}
	return res
}
