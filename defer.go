package main
import "fmt"


func main(){
defer printTwo()
printOne()
}

func printOne(){fmt.Print(1)}
func printTwo(){fmt.Print(2)}
