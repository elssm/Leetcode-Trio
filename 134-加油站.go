package goLeetCode

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	var carOil, minCarOil, minIndex int
	for i := 0; i < len(gas); i++ {
		carOil = carOil + gas[i] - cost[i]
		fmt.Println(carOil, minCarOil)
		if carOil < minCarOil {
			minCarOil = carOil
			minIndex = i + 1
		}
	}
	if carOil < 0 {
		return -1
	}
	return minIndex % len(gas)

}
