package main

import "fmt"

func main() {
    matriz := [2][3]int{{1, 2, 3}, {4, 5, 6}}

    var linha, coluna int
    fmt.Print("Informe o índice de linha: ")
    fmt.Scanln(&linha)
    fmt.Print("Informe o índice de coluna: ")
    fmt.Scanln(&coluna)

    fmt.Println("Valor na posição [", linha, "][", coluna, "]:", matriz[linha][coluna])
}
