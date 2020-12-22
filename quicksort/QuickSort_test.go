package quicksort

import (
	"reflect"
	"testing"
)

type TestData struct {
	TestArray []int
	Want      []int
}

func TestQuickSort(t *testing.T) {
	tests := []TestData{
		TestData{TestArray: []int{3, 4, 2, 5, 1, 7, 6, 8, 9}, Want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		TestData{TestArray: []int{3, 4, 2, 0, 5, 1, 7, -1, 6}, Want: []int{-1, 0, 1, 2, 3, 4, 5, 6, 7}},
	}
	for _, test := range tests {
		got := QuickSort(test.TestArray)
		if !reflect.DeepEqual(got, test.Want) {
			t.Errorf("Got: #%v, Want: #%v", got, test.Want)
		}
	}
}
