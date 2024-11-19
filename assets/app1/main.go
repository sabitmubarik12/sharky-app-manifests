package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "There’s no such thing as a bad day when you’re fishing.")
  })
  http.ListenAndServe(":8080", nil)
}
