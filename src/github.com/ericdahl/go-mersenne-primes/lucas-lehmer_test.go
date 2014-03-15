package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestPrime(t *testing.T) {
    assert.True(t, isPrime(3))
    assert.True(t, isPrime(5))
    assert.True(t, isPrime(7))
    assert.True(t, isPrime(13))
}
