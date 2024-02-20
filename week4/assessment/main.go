package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]Animal)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		if input == "" {
			continue
		}
		if input == "exit" {
			break
		}
		action, parameters := input[:strings.Index(input, " ")], input[strings.Index(input, " ")+1:]

		switch action {
		case "newanimal":
			name, animalType := parameters[:strings.Index(parameters, " ")], parameters[strings.Index(parameters, " ")+1:]
			switch animalType {
			case "cow":
				animals[name] = Cow{name}
			case "bird":
				animals[name] = Bird{name}
			case "snake":
				animals[name] = Snake{name}
			}
			fmt.Println("Created it!")
		case "query":
			name, info := parameters[:strings.Index(parameters, " ")], parameters[strings.Index(parameters, " ")+1:]
			switch info {
			case "eat":
				animals[name].Eat()
			case "move":
				animals[name].Move()
			case "speak":
				animals[name].Speak()
			}
		default:
			fmt.Println("Invalid command")
		}
	}
}
