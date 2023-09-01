package main

import "net/http"

type VMServer struct {
}


func (v VMServer) CreateVM(w http.ResponseWriter, r *http.Request) {
}


func main() {
  s := VMServer{}
  h := Handler(s)

  http.ListenAndServe(":7000", h)
}
