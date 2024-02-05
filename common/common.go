package common

import "math/rand"

func GetRandomNumbers(size int, min int, max int) []int {
	numArray := make([]int, size)
	for i := 0; i < size; i++ {
		numArray[i] = rand.Intn(max-min+1) + min
	}
	return numArray
}
