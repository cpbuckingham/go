package main

import (  "fmt"
          "strings"
          "sort"
          "os"
          "log"
          "io/ioutil"
          "strconv"

func main(){

  randInt:=5
  randFloat:=10.5
  randString:="100"
  randString2:="250.5"

fmt.Println(float64(randInt))
fmt.Println(int(randFloat))

newInt, _ :=strconv.ParseInt(rantString, 0, 64)
fmt.Println(newInt)

newFloat, _ :=strconv.ParseFloat(rantString2, 64)
fmt.Println(newFloat)

}
