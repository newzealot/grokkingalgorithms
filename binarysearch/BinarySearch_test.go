package binarysearch

import (
	"testing"
)

type TestData struct {
	TestArray []int
	Want      int
}

func TestBinarySearch(t *testing.T) {
	tests := []TestData{
		TestData{TestArray: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, Want: 8},
		TestData{TestArray: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, Want: 5},
		TestData{TestArray: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, Want: 2},
	}
	for _, test := range tests {
		got, err := BinarySearch(test.TestArray, test.Want)
		if err != nil {
			t.Error(err)
		}
		if got != test.Want {
			t.Errorf("Got: %#v, Want: %#v", got, test.Want)
		}
	}
}
