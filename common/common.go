package common

import (
	"fmt"
	"math/rand"
)

/*
GetRandomNumbers

Generated random integers withing a min and max range with min and max inclusive
*/
func GetRandomNumbers(size int, min int, max int) []int {
	numArray := make([]int, size)
	for i := 0; i < size; i++ {
		numArray[i] = rand.Intn(max-min+1) + min
	}
	return numArray
}

/*
PrintIntArray

Print array of integers in rows with the number of columns specified as a parameter
*/
func PrintIntArray(cols int, arrayData []int) {
	colCounter := 0
	for i := 0; i < len(arrayData); i++ {
		fmt.Printf("%6d", arrayData[i])
		colCounter++
		if colCounter == cols {
			fmt.Println()
			colCounter = 0
		} else {
			fmt.Printf(", ")
		}
	}
	fmt.Println()
}
