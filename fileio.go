package main

import (  "fmt"
          "strings"
          "sort"
          "os"
          "log"
          "io/ioutil"
          "strconv"

func main(){
  file, err:= os.Create("samp.txt")
  if err != nil{
    log.Fatal(err)

  }
  file.WriteString("this is the random text")

  file.Close()

  stream, err := ioutil.ReadFile("samp.txt")
  if err != nil{
    log.Fatal(err)

  }
  readString : string(stream)

  fmt.Println(readString)

}
