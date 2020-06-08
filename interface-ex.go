package main


import (
	"fmt"
	"math")

type polygon interface {
	perim() float32
	area() float32
}

type Rectangle struct {
	l, b float32
}

func (r Rectangle) area() float32 {
  return r.l * r.b
}

func (r Rectangle) perim() float32 {
	return 2 * (r.l + r.b)
}



type Circle struct {
	radius float32
}

func (c Circle) area() float32 {
    return math.Pi * c.radius * c.radius
}
func (c Circle) perim() float32 {
    return 2 * math.Pi * c.radius
}

func measure(shape polygon) {
	fmt.Println(shape)
    fmt.Println(shape.area())
    fmt.Println(shape.perim())
}

func emtyinterfaceASArg(arg interface{}) {
	fmt.Println(arg)
}

func interAsReturnType() polygon {
  return Rectangle{4,4}
}


func main() {

	
	fmt.Println(interAsReturnType().area())

	emtyinterfaceASArg("Abc")
	emtyinterfaceASArg(123)
	r := Rectangle{2.3,3.0}
	measure(r)

	c := Circle{5.3}
	measure(c)

	// Array of interface objects

	slices := []polygon {Rectangle{1,2}, Circle{7}}
	
	for _, shape:= range slices {
		fmt.Println(shape.area())
	}

}