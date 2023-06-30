package main

import "fmt"

func main() {
    array := [4]float64{1.5, 2.0, 2.5, 3.0}

    produto := 1.0
    for _, value := range array {
        produto *= value
    }

    fmt.Println(produto)
}
