package main

import (
	"fmt"
	"strings"
)

// Greet takes one or more names and returns a greeting
// By the last test, this turned ugly
func Greet(names ...string) string {
	firstUpperIndex, names := sortNames(names)
	name := ""
	for nameIndex, n := range names {
		if nameIndex == 0 {
			name += n
		} else if nameIndex == firstUpperIndex-1 && len(names) > firstUpperIndex {
			name += " and " + n + "."
		} else if nameIndex == firstUpperIndex-1 {
			name += " and " + n
		} else if nameIndex == firstUpperIndex && strings.ToUpper(n) == n {
			name += " AND HELLO " + n
		} else if strings.ToUpper(n) == n {
			name += " AND " + n
		} else {
			name += ", " + n + ","
		}
	}

	switch {
	case name == "":
		return "Hello, my friend."
	case strings.ToUpper(name) == name:
		return "HELLO " + name + "!"
	case len(names) > firstUpperIndex:
		return "Hello, " + name + "!"
	default:
		return "Hello, " + name + "."
	}
}

func sortNames(names []string) (int, []string) {
	lowerNames := []string{}
	upperNames := []string{}
	for _, n := range names {
		if strings.ToUpper(n) == n {
			upperNames = append(upperNames, n)
		} else {
			lowerNames = append(lowerNames, n)
		}
	}
	return len(lowerNames), append(lowerNames, upperNames...)
}

func main() {
	fmt.Println(sortNames([]string{"Mike", "JACK", "Nick"}))
}
