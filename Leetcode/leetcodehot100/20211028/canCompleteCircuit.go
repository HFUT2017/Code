package _0211028

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); {
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < len(gas) {
			j := (i + cnt) % len(gas)
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfGas < sumOfCost {
				break
			}
			cnt++
		}
		if cnt == len(gas) {
			return i
		}
		i = i + cnt + 1
	}
	return -1
}
