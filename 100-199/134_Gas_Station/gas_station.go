package gasstation

func canCompleteCircuit(gas []int, cost []int) int {
	leftAmountOfGas := 0
	resultIndex := -1
	sumGas := 0
	sumCost := 0
    for i := 0; i < len(gas); i++ {
		if leftAmountOfGas == 0 {
			resultIndex = i
		}
		leftAmountOfGas += gas[i]-cost[i]
		if leftAmountOfGas < 0 {
			leftAmountOfGas = 0
			resultIndex = -1
		}
		sumGas += gas[i]
		sumCost += cost[i]
	}
	if sumGas >= sumCost {
		return resultIndex
	}
	return -1
}