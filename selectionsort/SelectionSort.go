package selectionsort

import (
	"fmt"
)

func findSmallest(intArray []int) int {
	smallestValue := intArray[0]
	smallestIndex := 0
	for i := 0; i < len(intArray); i++ {
		if intArray[i] < smallestValue {
			smallestValue = intArray[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func SelectionSort(intArray []int) []int {
	sortedArray := []int{}
	intArrayLength := len(intArray)
	for i := 0; i < intArrayLength; i++ {
		smallestIndex := findSmallest(intArray)
		sortedArray = append(sortedArray, intArray[smallestIndex])
		intArray[smallestIndex] = intArray[len(intArray)-1]
		intArray = intArray[:len(intArray)-1]
	}
	fmt.Println(sortedArray)
	return sortedArray
}
