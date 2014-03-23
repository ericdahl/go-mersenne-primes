package main

import (
    "testing"
    "log"
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

func TestPrimality(t *testing.T) {
    p := 0
    for i:= 2; i <= mersennes[len(mersennes) - 1]; i++ {
        log.Printf("isPrime(%d) -> [%t]\n", i, isPrime(i))
        if (mersennes[p] == i) {
            if !isPrime(i) {
                t.Errorf("[%d] should be calculated as prime", i)
            }
            p++
        } else {
            if isPrime(i) {
                t.Errorf("[%d] should be calculated as composite", i)
            }
        }
    }
}

func TestLarge(t *testing.T) {
    if !isPrime(607) {
        t.Errorf("[%d] should be calculated as prime", 607)
    }
}

func benchmarkLarge(n int, b *testing.B) {
    for i := 0; i < b.N; i++ {
        isPrime(n)
    }
}

func Benchmark107(b *testing.B) { benchmarkLarge(107, b); }
func Benchmark1279(b *testing.B) { benchmarkLarge(1279, b); }
func Benchmark11213(b *testing.B) { benchmarkLarge(11213, b); }
