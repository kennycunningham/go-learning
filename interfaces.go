package main

import "fmt"

type Person struct {
	name string
	cell int
}

type I interface {
	Call()
	Text()
}

func (bob Person) Text() {
	fmt.Println(bob)
}

func (bob Person) Call() {
	fmt.Println(bob)
}

func main() {
	var bob I = Person{"Bob Barker", 1234567890}
	bob.Call()
	bob.Text()
}
