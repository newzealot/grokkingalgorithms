package breadthfirstsearch

import (
	"testing"
)

type TestData struct {
	testArray map[string][]string
	searchFor string
	want      string
}

func TestBreadFirstSearch(t *testing.T) {
	tests := []TestData{
		TestData{
			testArray: map[string][]string{
				"you":    []string{"alice", "bob", "claire"},
				"alice":  []string{"peggy"},
				"bob":    []string{"anuj", "peggy"},
				"claire": []string{"thom", "johnny"},
				"anuj":   []string{},
				"peggy":  []string{},
				"thom":   []string{},
				"johnny": []string{},
			},
			searchFor: "you",
			want: "thom",
		},
		TestData{
			testArray: map[string][]string{
				"you":    []string{"alice", "bob", "claire"},
				"alice":  []string{"peggy"},
				"bob":    []string{"anuj", "peggy"},
				"claire": []string{"thomas", "johnny"},
				"anuj":   []string{},
				"peggy":  []string{},
				"thom":   []string{},
				"johnny": []string{},
			},
			searchFor: "you",
			want: "",
		},
	}
	for _, test := range tests {
		got := BreadthFirstSearch(test.testArray, test.searchFor)
		if got != test.want {
			t.Errorf("\ngot: %#v\nwant: %#v\n", got, test.want)
		}
	}
}
