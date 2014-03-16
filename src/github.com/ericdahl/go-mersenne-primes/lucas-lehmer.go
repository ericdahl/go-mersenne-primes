package main

import (
    "math/big"
)

func isPrime(p int) bool {
    zero := big.NewInt(0)
    one := big.NewInt(1)
    two := big.NewInt(2)

    p2 := big.NewInt(int64(p))
    s2 := big.NewInt(4)
    m2 := big.NewInt(2)

   if 0 == two.Cmp(p2) {
       // special case for 2
       return true
   }

    m2.Exp(m2, p2, nil)
    m2.Sub(m2, one)

    for i := 0; i < p - 2; i++ {
        s2.Mul(s2, s2)
        s2.Sub(s2, two)
        s2.Mod(s2, m2)
    }
    return 0 == zero.Cmp(s2)
}
