package main

import (
	"algorithms/common"
	"algorithms/sort/selection"
	"fmt"
)

func main() {
	randomData := common.GetRandomNumbers(20, 7, 17)

	fmt.Println("Random Data:")
	for i := 0; i < len(randomData); i++ {
		fmt.Println(randomData[i])
	}

	sortedData := selection.SortArray(randomData)

	fmt.Println("============")
	fmt.Println("Sorted Data:")
	for i := 0; i < len(sortedData); i++ {
		fmt.Println(sortedData[i])
	}

}
