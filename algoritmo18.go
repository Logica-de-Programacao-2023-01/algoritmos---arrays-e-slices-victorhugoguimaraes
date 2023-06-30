package main

import "fmt"

func isPrime(num int) bool {
    if num <= 1 {
        return false
    }

    for i := 2; i <= num/2; i++ {
        if num%i == 0 {
            return false
        }
    }

    return true
}

func main() {
    var n int
    fmt.Print("Digite um número inteiro positivo: ")
    fmt.Scanln(&n)

    var primes []int
    i := 2
    for len(primes) < n {
        if isPrime(i) {
            primes = append(primes, i)
        }
        i++
    }

    fmt.Println(primes)
}
