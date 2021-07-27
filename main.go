package main

import "fmt"
import "net/http"
//import "time"

func main() {
  // Create serve mux
  mux := http.NewServeMux()

  // Handle all request
  mux.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
    if request.URL.Path != "/" {
      http.NotFound(response, request)
      return
    }
    response.Header().Set("Content-Type", "text/plain; charset=utf-8")
    response.WriteHeader(http.StatusOK)
    response.Write([]byte("Hi, I'm David Gaspar\n"))
  })
  fmt.Printf("mux: %+v\n", mux)
  http.ListenAndServe(":7080", mux)
}
