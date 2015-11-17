package main

import (  "fmt"
          "net/http"

func main(){
http.HandleFunc("/", handler)
http.HandleFunc("/earth", handler2)
http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "hello world\n")

func handler2(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "hello earth\n")
}
