package main

import (
    "math"
)

func isPrime(p int) bool {
    var (
        s int = 4
        m int = int(math.Pow(2, float64(p)) - 1)
    )

    for i := 0; i < p - 2; i++ {
        s = ((s * s) - 2) % m
    }
    return s == 0 
}
