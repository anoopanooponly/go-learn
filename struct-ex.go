package main


import (
	"fmt"
	"math")


type Person struct {
	name string
	status bool
	age int
}

func newPerson(name string) *Person {
 
  p := Person{name : name}
  p.age = 53
  p.status = true
  return &p
}

type Rectangle struct {
	l, b int
}

func (r Rectangle) perim() int {
	return 2 * ( r.l + r.b)
}

type Circle struct {
	radius float32
}

func (c Circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func main(){

	p := Person{name : "Anoop", age : 36}
  p.age = 22
	fmt.Println(p)

	 p0 := &Person{name : "Anoop", age : 36}
	 p0.status =true
	 fmt.Println(p0)

	p1 := newPerson("chithra")
	p1.age = 30
	fmt.Println(p1)


	p2:= &p // both p2 and p will change. p2:= p - copy will be modified
	p2.age =10
	fmt.Println(p2)
	fmt.Println(p)

	//struct method

	var rect Rectangle = Rectangle{2,3}
	fmt.Println(rect.perim())

	//passing struct to function
	 c1 := Circle{}
	 itakenopointer(c1)
	 fmt.Println(c1.radius) 

	 c2 := &Circle{}
	 itakepointer(c2)
	 fmt.Println(c2.radius)
	 
	 c3 := new(Circle) //no memory taken here and returns a pointer , same as &type{}
	 itakepointer(c3)
	 fmt.Println(c3.radius) 
}

func itakenopointer(c Circle) {
  c.radius = 4
}

func itakepointer(c *Circle) {
	c.radius = 4
  }