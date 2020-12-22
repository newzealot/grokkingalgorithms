package quicksort

import (
	"math/rand"
	"time"
)

func QuickSort(intArray []int) []int {

	if len(intArray) < 2 {
		return intArray
	}
	smallerArray := []int{}
	biggerArray := []int{}
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	index := rand.Intn(len(intArray))
	pivot := intArray[index]
	intArray = append(intArray[:index], intArray[index+1:]...)
	for _, element := range intArray {
		if element <= pivot {
			smallerArray = append(smallerArray, element)
		} else if element > pivot {
			biggerArray = append(biggerArray, element)
		}
	}
	sortedArray := []int{}
	sortedArray = append(sortedArray, QuickSort(smallerArray)...)
	sortedArray = append(sortedArray, pivot)
	sortedArray = append(sortedArray, QuickSort(biggerArray)...)
	return sortedArray

}
