package selection

/*
SortArray

Sorts an array of integers using selection sort algorithm
*/
func SortArray(dataArray []int) []int {

	for i := 0; i < len(dataArray)-1; i++ {

		//assume minIndex is i in the beginning of the loop
		minIndex := i

		for j := i + 1; j < len(dataArray); j++ {
			if dataArray[minIndex] > dataArray[j] {
				minIndex = j
				//change minIndex when the value in minIndex is not the minimum value
			}
		}

		//minIndex has changed, swap the array values
		if minIndex != i {
			dataArray[minIndex], dataArray[i] = dataArray[i], dataArray[minIndex]
		}
	}

	return dataArray
}
