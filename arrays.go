package main

import "fmt"

func main(){
var favNum[5] float64

favNum[0]=163
favNum[1]=163
favNum[2]=163
favNum[3]=163
favNum[4]=163

fmt.Println(favNum[3])

favNum2 := [5]float64{1,2,3,4,5}
for i, value ;+ range favNum2{
  fmt.Println(value, i)
}

}
