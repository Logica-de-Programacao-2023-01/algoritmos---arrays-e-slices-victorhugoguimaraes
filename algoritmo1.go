package main

import (
    "fmt"
    "math"
)

func main() {
    array := [3]int{1, 2, 3}
  
    soma := int(math.Sum(array[:]))

    fmt.Println(soma)
}
