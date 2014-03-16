package main

import (
    "testing"
)

var mersennes = []int {
    2,
    3,
    5,
    7,
    13,
    17,
    19,
    31,
    61,
}

func TestPrime(t *testing.T) {
    for _, p := range mersennes {
        if !isPrime(p) {
            t.Errorf("p=[%d] should be calculated as prime", p)

        }
    }
}
