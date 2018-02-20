package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(test[0:6])

	test = append(test, 5)
	printSlice(test[:])
}
