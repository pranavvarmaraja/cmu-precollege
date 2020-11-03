package main

import "fmt"

type Contact struct {
	firstName, lastName string //shorthand for declaring two fields of same type
	phone               []int
	email               string
	zipCode             [5]int
	// etc.
	//Twitter username (string)
}

/*
// we have seen this before when we declared our own types
type StringIndex map[string][]int

// or cellular automata game boards
type GameBoard [][]int

//this allowed us to do things like
func PlayCellularAutomaton(board GameBoard, numGens int) []GameBoard {

}

//instead of
func PlayCellularAutomaton(board [][]int, numGens int) [][][]int {

}
*/

func PrintContact(c Contact) {
	fmt.Println("First name is:", c.firstName)
	fmt.Println("Zip code is:", c.zipCode)
	//etc.
}

func main() {
	// let's create a contact
	//x := 3 means var x int = 3
	var me Contact

	//what are me's fields right now?

	fmt.Println(me)

	//Go starts a variable with default values for each field
	//e.g., an int is 0, a float is 0.0, a string is "",  a slice is nil, an array
	//of length k has k default values

	//how can we set the fields of me? We use "object.field = blah"
	me.firstName = "Phillip"
	me.lastName = "Compeau"
	me.zipCode = [5]int{1, 5, 2, 1, 3}
	me.email = "pcompeau@andrew.cmu.edu"
	// any slices need to be set or made
	me.phone = make([]int, 7)
	//etc.
	me.phone = []int{3, 3, 0, 8, 0, 0, 4}
	fmt.Println(me)

	PrintContact(me)
}
