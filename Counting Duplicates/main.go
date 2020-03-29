package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(duplicate_count("Indivisibilities"))
}

func duplicate_count(s1 string) int {
	count := 0
	world := strings.ToUpper(s1)
	for _, char := range world {
		if strings.Count(world, string(char)) > 1 {
			count++
			world = strings.ReplaceAll(world, string(char), "")
		}
	}

	return count
}
