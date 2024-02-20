package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	animals := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

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
		animal, action := input[:strings.Index(input, " ")], input[strings.Index(input, " ")+1:]
		if (animal != "cow" && animal != "bird" && animal != "snake") || (action != "eat" && action != "move" && action != "speak") {
			fmt.Println("Invalid input")
			continue
		}
		switch action {
		case "eat":
			animals[animal].Eat()
		case "move":
			animals[animal].Move()
		case "speak":
			animals[animal].Speak()
		}
	}
}
