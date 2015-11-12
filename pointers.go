package main
import "fmt"


func main(){
x:=0
changeXVal(&x)
fmt.Println("x= ", x)
}
//passed an integer called x
//*means change as a referrence
//builds on top of a variable
func changeXValNow(x *int){
  *x=2
}
