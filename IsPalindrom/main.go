package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrom("anana"))
}

func isPalindrom(s1 string) bool {
	res1 := strings.Split(
		strings.ReplaceAll(
			strings.ToUpper(s1), " ", ""), "")
	length := len(res1)
	count := 0
	for i := 0; i < length/2; i++ {
		if res1[i] == res1[length-i-1] {
			count++
		}
	}
	return length/2 == count
}
