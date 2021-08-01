package main

import (
	"fmt"
	"sort"
)

func main() 
{
	num1 := []int{2, 4, 6, 8, 10, 12}
	num2 := []int{12, 10, 8, 6, 4, 2}
	fmt.Println("The addition of two slices", Add(num1, num2))

}
func Add(num1 []int, num2 []int) []int {
	c := make([]int, len(num1))
	sort.Sort(sort.Reverse(sort.IntSlice(num2)))
	for i, _ := range num1 {
		c[i] = num1[i] + num2[i]

	}
	return c

}
