package main

import "fmt"

func main() {
    var array [10]int

    var valor int
    fmt.Print("Informe um valor a ser buscado: ")
    fmt.Scan(&valor)

    for i := 0; i < len(array); i++ {
        array[i] = i * 2
    }

    encontrado := false
    for _, elemento := range array {
        if elemento == valor {
            encontrado = true
            break
        }
    }

    if encontrado {
        fmt.Printf("O valor %d está presente no array.\n", valor)
    } else {
        fmt.Printf("O valor %d não está presente no array.\n", valor)
    }
}
