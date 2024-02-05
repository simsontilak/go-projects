package selection

func SortArray(dataArray []int) []int {

	for i := 0; i < len(dataArray)-1; i++ {
		minIndex := i

		for j := i + 1; j < len(dataArray); j++ {
			if dataArray[minIndex] > dataArray[j] {
				minIndex = j
			}
		}

		dataArray[minIndex], dataArray[i] = dataArray[i], dataArray[minIndex]
	}

	return dataArray
}
