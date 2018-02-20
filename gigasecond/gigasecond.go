package main

import "fmt"

func main() {
	var year float64
	var giga float64
	year = 60 * 60 * 24 * 7 * 365
	giga = 1000000000
	gigasecond := float64(giga / year)
	fmt.Printf("%f\n", gigasecond)
}
