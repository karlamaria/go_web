package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
}

func main() {
    // Tells the http package to handle all requests to the web root ("/")
    http.HandleFunc("/", helloWorld)
    // Listen on a port to accept connections from the internet
    http.ListenAndServe(":9999", nil)
}