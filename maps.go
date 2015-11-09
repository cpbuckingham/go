package main

import "fmt"

func main(){

  presAge:= make(map[string] int)
  presAge["Theodore Roosevelt"]=42
  fmt.Println(len(presAge))
  presAge["John F. Kennedy"]=43
  fmt.Println(len(presAge))
  delete(presAge, "John F. Kennedy")
  fmt.Println(len(presAge))
  delete(presAge, [0])
  fmt.Println(len(presAge))

}
