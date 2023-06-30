package main

import "fmt"

func main() {
    numeros := [7]int{5, 7, 2, 9, 1, 8, 4}

    var numero int
    fmt.Print("Informe o número a ser adicionado: ")
    fmt.Scanln(&numero)

    numeros[0] += numero
    numeros[len(numeros)-1] += numero

    fmt.Println("Array resultante:", numeros)
}
