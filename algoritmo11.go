package main

import "fmt"

func main() {
    numeros := []int{4, 9, 2, 1, 5, 7, 3, 8, 6, 0}

    minimo := numeros[0]
    maximo := numeros[0]

    for _, numero := range numeros {
        if numero < minimo {
            minimo = numero
        }
        if numero > maximo {
            maximo = numero
        }
    }

    fmt.Println("Valor mínimo:", minimo)
    fmt.Println("Valor máximo:", maximo)
}
