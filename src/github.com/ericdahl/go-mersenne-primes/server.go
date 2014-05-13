package main 

import (
    "fmt"
    "net/http"
    "html"
    "log"
    "strings"
    "strconv"
)

func main() {
    http.HandleFunc("/", handleRequest)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    log.Printf("Processing [%q]", html.EscapeString(r.URL.Path))

    p, e := strconv.Atoi(strings.Trim(r.URL.Path, "/"))
    if e != nil {
        fmt.Println(e)
    }

    fmt.Fprintf(w, "%t", isPrime(p));
}
