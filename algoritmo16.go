package main

import "fmt"

func main() {
    arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    var slice []int

    for i := 0; i < len(arr); i++ {
        if arr[i]%2 == 0 {
            slice = append(slice, arr[i])
        }
    }

    fmt.Println(slice)
}
