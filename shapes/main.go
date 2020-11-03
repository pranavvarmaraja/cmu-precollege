package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	x1, y1                  float64 //lower left corner
	width, height, rotation float64
}

type Circle struct {
	x1, y1 float64 //center
	radius float64
}

/*
//This gives compiler error -- can't name two functions the same thing

func Area(r Rectangle) float64 {
	return r.width * r.height
}

func Area(c Circle) float64 {
	return math.Pi * c.radius * c.radius
}
*/

//solution: write a "method" for each of the shapes

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

//exercise: write Translate() methods that shift a circle and a rectangle
//by a units in the x direction and b units in the y direction (a and b are input parameters).

/*
func (r Rectangle) Translate(a, b float64) Rectangle {
	r.x1 += a
	r.y1 += b
	return r
}

func (c Circle) Translate(a, b float64) Circle {
	c.x1 += a
	c.y1 += b
	return c
}
*/

//solution to duplicating objects is to use a pointer!

func (c *Circle) Translate(a, b float64) {
	c.x1 += a
	c.y1 += b
}

func (r *Rectangle) Translate(a, b float64) {
	r.x1 += a
	r.y1 += b
}

// exercise: write a circle and rectangle method using pointers called Scale()
// that "dilates" the size of each shape by an input factor f.

func (c *Circle) Scale(f float64) {
	c.radius *= f
}

func (r *Rectangle) Scale(f float64) {
	r.width *= f
	r.height *= f
}

func ChangeFirst(list []int) {
	list[0] = 1
}

func main() {
	fmt.Println("Methods and pointers.")

	var r Rectangle
	r.width = 3.0
	r.height = 4.0

	var c Circle
	c.radius = 5.0

	//we access methods in the same way we access fields
	fmt.Println(r.Area())

	fmt.Println(c.Area())

	//let's move our rectangle and circle.
	//r = r.Translate(-3.1, 2.7)
	//c = c.Translate(4.78, 0.03)

	fmt.Println(r.x1, r.y1)
	fmt.Println(c.x1, c.y1)
	// nothing happened!

	//A pointer is a variable that stores the address in memory of some other variable.
	//In a sense, a pointer is like a bookmark, or an index.
	//If you know where something is, you can change it.
	var b int = -14
	var a *int // * means "pointer to"
	a = &b     // &  means "location of"
	fmt.Println("b is", b)
	fmt.Println("The location of b is", a)

	//if we have a pointer to an object, we can change it.
	var pointerToC *Circle

	fmt.Println(pointerToC) // at this point, pointer has value "nil"

	//let's point it at our circle
	pointerToC = &c

	fmt.Println(pointerToC)

	//next big idea: if I have a pointer, I can CHANGE the fields of the thing it points to
	(*pointerToC).x1 = -1.7 // *p means "variable that p points to"
	//this is equivalent to c.x1 = -1.7

	//Go is nice because I don't need the * dereference
	pointerToC.y1 = 1478.3 // this is fine

	fmt.Println(c.x1) //should see -1.7

	fmt.Println(c.x1, c.y1)

	pointerToC.Translate(-1.0, 1.0)

	fmt.Println(c.x1, c.y1)

	var r2 Rectangle
	//Go doesn't mind whether you give it a pointer or the object itself
	r2.width = 2.0
	r2.height = 3.0
	r2.Scale(2.0)
	fmt.Println(r2.width, r2.height)

	list := make([]int, 5)
	ChangeFirst(list)
	fmt.Println(list)
	//Go quiz: what gets printed?

	//what happened? ChangeFirst changed our array!
	//Why? Because slices are *pointers* to arrays

	//To be more precise, a slice is a pair -- it's a pointer to a position in an array
	//along with the length of the slice

}
