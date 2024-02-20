package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() {
	fmt.Println(d.Name, "says Woof!")
}

type Cat struct {
	Name string
}

func (c *Cat) Speak() {
	if c != nil {
		fmt.Println(c.Name, "says Meow!")
	}
}

func main() {
	var mila *Cat

	mila = &Cat{"Mila"}

	speakers := [10]Speaker{
		&Dog{"Fido"},
		&Cat{"Milo"},
		&Dog{"Rex"},
		&Cat{"Luna"},
		&Dog{"Buddy"},
		&Cat{"Molly"},
		&Dog{"Max"},
		&Cat{"Lucy"},
		&Dog{"Charlie"},
		mila,
	}

	for _, speaker := range speakers {
		speaker.Speak()
	}
}

