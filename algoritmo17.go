package main

import "fmt"

func main() {
    arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    var sum int

    for i := 0; i < len(arr); i++ {
        if i%2 == 0 {
            sum += arr[i]
        }
    }

    fmt.Println(sum)
}
