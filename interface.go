package main
import "fmt"
import "math"


func main(){
rect:=Rectangle{20,50}
circ:=Circle{4}

fmt.Println(getArea(rect))
fmt.Println(getArea(circ))
}

type Shape interface{
  area() float64
}

type Rectangle struct{
  height float64
  width float64
}

type Circle struct{
  radius float64
}

func (r Rectangle) area() float64{
  return r.height * r.width
}

func (c Circle) area() float64{
  return math.Pi *math.Pow(c.radius, 2)
  //math pi & math power
  }

func getArea(shape Shape) float64{
  return shape.area()
}
