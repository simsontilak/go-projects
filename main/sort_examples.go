package main

import (
	"algorithms/common"
	"algorithms/sort/bubble"
	"algorithms/sort/selection"
	"fmt"
	"time"
)

func main() {
	randomData := common.GetRandomNumbers(20, 21, 99)

	dataForSelection := make([]int, len(randomData))
	dataForBubble := make([]int, len(randomData))
	copy(dataForSelection, randomData)
	copy(dataForBubble, randomData)

	common.PrintIntArrayMsg(20, randomData, "Unsorted")

	sortedSelectionData := selection.SortArray(dataForSelection)
	common.PrintIntArrayMsg(20, sortedSelectionData, "Selection Sort")

	sortedBubbleData := bubble.SortArray(dataForBubble)
	common.PrintIntArrayMsg(20, sortedBubbleData, "Bubble Sort")

	largeRandomData := common.GetRandomNumbers(100000, 1, 200000)
	dataForSelection = make([]int, len(largeRandomData))
	dataForBubble = make([]int, len(largeRandomData))
	copy(dataForSelection, largeRandomData)
	copy(dataForBubble, largeRandomData)

	selectionTime := time.Now()
	selection.SortArray(dataForSelection)
	selectionTimeTaken := time.Since(selectionTime)
	fmt.Println("Selection sort time =", selectionTimeTaken.Milliseconds())

	bubbleTime := time.Now()
	bubble.SortArray(dataForBubble)
	bubbleTimeTaken := time.Since(bubbleTime)
	fmt.Println("Bubble sort time =", bubbleTimeTaken.Milliseconds())

}
