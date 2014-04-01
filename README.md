go-mersenne-primes
==================
[![Build Status](https://travis-ci.org/ericdahl/go-mersenne-primes.png?branch=master)](https://travis-ci.org/ericdahl/go-mersenne-primes)

Web service which computes mersenne primes (experimental)


#### Quick Start
```bash
$ git clone https://github.com/ericdahl/go-mersenne-primes.git
$ cd go-mersenne-primes
$ export GOPATH=`pwd`
$ go install github.com/ericdahl/go-mersenne-prime
$ bin/go-mersenne-primes
$ curl 'http://localhost:8080/7'
true
```

#### Testing
```bash
$ go test github.com/ericdahl/go-mersenne-primes 
ok      github.com/ericdahl/go-mersenne-primes  0.008s
```
