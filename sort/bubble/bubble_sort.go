package bubble

/*
SortArray

Sorts an array of integers using bubble sort algorithm
*/
func SortArray(dataArray []int) []int {
	arrayLen := len(dataArray)
	for i := 0; i < arrayLen; i++ {
		swapped := false
		for j := 0; j < arrayLen-i-1; j++ {
			if dataArray[j] > dataArray[j+1] {
				dataArray[j], dataArray[j+1] = dataArray[j+1], dataArray[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return dataArray
}
