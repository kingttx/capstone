package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, HelloServerLogic(r.URL.Path[1:]))
}

func HelloServerLogic(name string) string  {
    return fmt.Sprintf("Howdy, %s!", name)
}
