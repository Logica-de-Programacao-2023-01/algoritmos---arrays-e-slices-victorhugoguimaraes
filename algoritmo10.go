package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	min := 9999999999999
	max := -999999999999

	for _, x := range slice {
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}
	fmt.Println("O maior valor é", max)
	fmt.Println("O menor valor é", min)
