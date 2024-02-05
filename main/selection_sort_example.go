package main

import (
	"algorithms/common"
	"algorithms/sort/selection"
)

func main() {
	randomData := common.GetRandomNumbers(20, 21, 99)
	common.PrintIntArray(20, randomData)

	sortedData := selection.SortArray(randomData)
	common.PrintIntArray(20, sortedData)
}
