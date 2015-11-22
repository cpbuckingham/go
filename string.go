package main

import (  "fmt"
          "strings"
          "sort"
          "os"
          "log"
          "io/ioutil"
          "strconv"

func main(){
  sampString:="Hello World"
  fmt.Println(strings.Contains(sampString, "lo"))
  fmt.Println(strings.Index(sampString, "lo"))
  fmt.Println(strings.Count(sampString, "l"))
  fmt.Println(strings.Replace(sampString, "l", "x", 3))

}
