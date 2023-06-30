package main

import "fmt"

func main() {
    var array [6]float64
  
    var numero float64
    fmt.Print("Informe um número: ")
    fmt.Scan(&numero)

    for i := 0; i < len(array); i++ {
        array[i] = numero
    }

    fmt.Println(array)
}
