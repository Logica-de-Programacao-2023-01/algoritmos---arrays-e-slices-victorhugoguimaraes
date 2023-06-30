package main

import "fmt"

func main() {
    slice := []string{"a", "b", "c", "d", "a", "e", "f", "a"}

    var valor string
    fmt.Print("Informe um valor: ")
    fmt.Scan(&valor)

    novoSlice := []string{}
    for _, elemento := range slice {
        if elemento != valor {
            novoSlice = append(novoSlice, elemento)
        }
    }
    fmt.Println(novoSlice)
}
