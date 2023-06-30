package main

import "fmt"

func main() {
    n := 5
    arr1 := [5]int{1, 2, 3, 4, 5}
    arr2 := [5]int{6, 7, 8, 9, 10}
    arr3 := [5]int{}

    for i := 0; i < n; i++ {
        arr3[i] = arr1[i] + arr2[i]
    }

    fmt.Println(arr3)
}
