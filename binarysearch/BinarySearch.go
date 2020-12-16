package binarysearch

import (
	"errors"
	"fmt"
)

func BinarySearch(intArray []int, toFind int) (int, error) {
	midIndex := len(intArray) / 2
	minIndex := 0
	maxIndex := len(intArray) - 1
	for i := 0; i < len(intArray); i++ {
		fmt.Printf("Pass: %d, intArray: %#v, minIndex: %#v, midIndex: %#v, maxIndex: %#v\n", i+1, intArray, minIndex, midIndex, maxIndex)
		switch {
		case toFind == intArray[midIndex]:
			fmt.Printf("Found: %#v\n", toFind)
			return toFind, nil
		case toFind > intArray[midIndex]:
			minIndex = midIndex
			midIndex = (maxIndex + minIndex) / 2
		case toFind < intArray[midIndex]:
			maxIndex = midIndex
			midIndex = (maxIndex + minIndex) / 2
		}
	}
	err := fmt.Sprintf("Not Found: %#v\n", toFind)
	return 0, errors.New(err)
}
