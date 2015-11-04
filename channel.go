package main
import "fmt"
import "time"
import"strconv"

var pizzaNum=0
var pizzaName=""

func makeDough(stringChan chan string){
  pizzaNum++
  pizzaName = "pizza #" + strconv.Itoa(pizzaNum)
  fmt.Println("make dough and send for sauce")
  stringChan <- pizzaName
  time.Sleep(time.Millisecond*10)
}

func makeSauce(stringChan chan string){
  pizza:= <-stringChan
  fmt.Println("added sauce and send", pizza, "ready for toppings")
  stringChan <- pizzaName
  time.Sleep(time.Millisecond*10)
}

func maketoppings(stringChan chan string){
  pizza:= <-stringChan
  fmt.Println("added toppings", pizza, "and ship")
  time.Sleep(time.Millisecond*10)

}

func main (){
stringChan := make(shan string)
for i:=0, i<3, i++{
  go makeDough(stringChan)
  go makeSauce(stringChan)
  go maketoppings(stringChan)
  time.Sleep(time.Millisecond*5000)
}
}
