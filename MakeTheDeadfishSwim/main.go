package main

import "fmt"

func main() {

	fmt.Println(Parse("uddodiilddzdsdddssirzdsyisbd"))
}

func Parse(data string) []int {
	result := make([]int, 0)
	sum := 0
	for _, char := range data {
		switch char {
		case 'i':
			sum++
		case 'd':
			sum--
		case 's':
			if sum < 0 {
				sum *= -1
			}
			sum *= sum

		case 'o':
			result = append(result, sum)
		}
	}

	return result

}
