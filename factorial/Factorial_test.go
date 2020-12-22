package factorial

import (
	"testing"
)

type TestData struct {
	testInt int
	want    int
}

func TestFactorial(t *testing.T) {
	tests := []TestData{
		TestData{testInt: 8, want: 40320},
		TestData{testInt: 5, want: 120},
	}
	for _, test := range tests {
		got := Factorial(test.testInt)
		if got != test.want {
			t.Errorf("got: #%v, want: #%v", got, test.want)
		}
	}

}
