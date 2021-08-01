package main

import (
	"fmt"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}
func main() {
	numb1 := []int{1, 2, 3, 4, 5}
	numb2 := []int{1, 2, 3, 4, 5}
	fmt.Print("Result :", sum(numb1+numb2))
}
