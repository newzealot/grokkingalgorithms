package breadthfirstsearch

import (
	"fmt"
)

func BreadthFirstSearch(testArray map[string][]string, searchFor string) string {
	queue := []string{}
	queue = append(queue, testArray[searchFor]...)
	searched := map[string]bool{}
	for x := 0; len(queue) != 0; x++ {
		fmt.Println("Looking at", queue[0])
		if searched[queue[0]] {
			fmt.Println("Skipping", queue[0])
			queue = queue[1:]
			continue
		}
		if queue[0][len(queue[0])-1:] == "m" {
			fmt.Printf("Found %#v\n\n", queue[0])
			return queue[0]
		}
		searched[queue[0]] = true
		queue = append(queue, testArray[queue[0]]...)
		fmt.Println("Not found\nAdded", testArray[queue[0]])
		queue = queue[1:]
	}
	fmt.Printf("Cannot find\n\n")
	return ""
}
