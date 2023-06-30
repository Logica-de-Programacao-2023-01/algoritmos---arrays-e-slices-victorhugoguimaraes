package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    fmt.Print("Informe o tamanho do slice: ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    size, _ := strconv.Atoi(scanner.Text())

    slice := make([]int, size)

    fmt.Printf("Informe %d valores para o slice, separados por vírgula: ", size)
    scanner.Scan()
    input := scanner.Text()
    values := strings.Split(input, ",")
    for i, value := range values {
        slice[i], _ = strconv.Atoi(strings.TrimSpace(value))
    }

    fmt.Printf("Slice: %v\n", slice)
    soma := 0
    for _, value := range slice {
        soma += value
    }
    fmt.Printf("Soma dos valores: %d\n", soma)
}
