package all

// Problem: 134
// Title: Gas Station
// Category: all
// Tags: all


func canCompleteCircuit(gas []int, cost []int) int {
	totalGas := 0
	totalCost := 0
	currGas := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]
		currGas += gas[i] - cost[i]

		if currGas < 0 {
			start = i + 1
			currGas = 0
		}
	}

	if totalGas < totalCost {
		return -1
	}

	return start

}
