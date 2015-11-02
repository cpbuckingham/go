package main

import "fmt"

func main(){
yourAge :=18

if yourAge >= 16{
  fmt.Println("you can drive")

  // stops once a case is met
} else if yourAge >=18 {
  fmt.Println("you can vote")
} else{
  fmt.Println("you can have fun")
}

}
