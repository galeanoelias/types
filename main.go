package main

import "fmt"

type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

// ALIAS declaration
type myAlias = course

// TYPES definitions
type newCourse course

type newBool bool

func (b newBool) String() string {
	if b {
		return "VERDADERO"
	}
	return "FALSO"
}

func main() {
	a := course{name: "Go"}
	b := myAlias{name: "Go"}
	c := newCourse{name: "Go"}
	b.Print()
	fmt.Printf("El tipo es: %T\n", a)
	fmt.Printf("El tipo es: %T\n", c)
	var d newBool = false
	fmt.Println(d.String())
}