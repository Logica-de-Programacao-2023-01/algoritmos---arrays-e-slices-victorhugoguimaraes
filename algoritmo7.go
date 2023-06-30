package main

import "fmt"

func main() {
    slice := make([]int, 5)

    var numero int
    fmt.Print("Informe um número inteiro: ")
    fmt.Scan(&numero)

    encontrado := false
    for _, elemento := range slice {
        if elemento == numero {
            encontrado = true
            break
        }
    }

    if !encontrado {
        slice = append(slice, numero)
    }
    fmt.Println(slice)
}
