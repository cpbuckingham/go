package main

import (  "fmt"
          "strings"
          "sort"
          "os"
          "log"
          "io/ioutil"
          "strconv"

func main(){
  csvString:= "1,2,3,4,5,6"
  fmt.Println(Strings.Split(csvString, ","))
  listOfLetters :=[]string{"c","a","b"}
  sort.Strings(listOfLetters)
  fmt.Println(listOfLetters)
  listOfNums:= strings.Join([]string{"3","2","1"}, ", ")
  fmt.Println(listOfNums)
