package main 

import (
    "fmt"
    "net/http"
    "html"
    "log"
)

func main() {
    http.HandleFunc("/", handleRequest)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    log.Printf("Processing [%q]", html.EscapeString(r.URL.Path))
    
    result := isPrime(7)    
    
    fmt.Fprintf(w, "%t", result);
}
