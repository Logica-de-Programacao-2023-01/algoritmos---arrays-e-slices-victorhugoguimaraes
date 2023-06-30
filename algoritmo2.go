package main

import "fmt"

func main() {
    slice := []int{1, 2, 3, 4, 5}

    index := 2

    newSlice := make([]int, len(slice)-1)

    copy(newSlice, slice[:index])
    copy(newSlice[index:], slice[index+1:])

    fmt.Println(newSlice)
}
