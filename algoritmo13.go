package main

import "fmt"

func main() {
    numeros := [5]int{8, 5, 6, 9, 12}

    multiplosDeTres := []int{}

    for _, numero := range numeros {
        if numero%3 == 0 {
            multiplosDeTres = append(multiplosDeTres, numero)
        }
    }

    fmt.Println("Elementos múltiplos de 3:", multiplosDeTres)
}
