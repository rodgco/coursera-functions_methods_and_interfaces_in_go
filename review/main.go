// main.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food, move, speak string
}

type Bird struct {
	food, move, speak string
}

type Snake struct {
	food, move, speak string
}

func main() {
	animalMap := make(map[string]Animal)
	for {
		printInstructions()
		request_p1, request_p2, request_p3 := getUserInput()
		if request_p1 == "X" {
			break
		}
		fmt.Println("You have chosen", request_p1, request_p2, request_p3)
		if request_p1 == "newanimal" {
			newAnimal, err := decideAnimalInit(request_p3)
			// Error handling for init animals
			if err != nil {
				fmt.Println("Try again, read the instructions carefully!")
				continue
			}
			// Append aninmal to slice
			animalMap[request_p2] = newAnimal
			fmt.Println("Created it!")
		} else if request_p1 == "query" {
			printAnimalInfo(animalMap, request_p2, request_p3)
		} else {
			fmt.Println("I can only create and query animals in the format explained above.")
		}
	}

}

func printInstructions() {
	fmt.Println("============================================================")
	fmt.Println("Would you like to create a new animal or query an existing animal?")
	fmt.Println("============================================================")
	fmt.Println("For new animal enter 'newanimal <animal_name> <animal_type>'.")
	fmt.Println("You can select the <animal_name> but the <animal_type> must be 'cow', 'bird' or 'snake'.")
	fmt.Println("============================================================")
	fmt.Println("For a query enter 'query <animal_name> <animal_information>'")
	fmt.Println("You can select the animal name but the <animal_information> must be 'eat', 'move' or 'speak'.")
	fmt.Println("============================================================")
}

func getUserInput() (string, string, string) {

	fmt.Print("> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	if text == "X" {
		return text, text, text
	}

	parts := strings.Fields(text)
	request_p1 := strings.ToLower(parts[0])
	request_p2 := strings.ToLower(parts[1])
	request_p3 := strings.ToLower(parts[2])

	return request_p1, request_p2, request_p3
}

func decideAnimalInit(animal string) (Animal, error) {
	switch animal {
	case "cow":
		return initCow(), nil
	case "bird":
		return initBird(), nil
	case "snake":
		return initSnake(), nil
	default:
		return nil, fmt.Errorf("unknown animal type: %s", animal)
	}
}

func printAnimalInfo(animalMap map[string]Animal, name string, action string) {
	animal, ok := animalMap[name]
	name = strings.ToUpper(name)
	if !ok {
		fmt.Println("Animal not found: ", name)
		return
	}

	switch aCh := animal.(type) {
	case Cow:
		if action == "eat" {
			fmt.Println(name, "eats:", aCh.food)
		} else if action == "move" {
			fmt.Println(name, "moves:", aCh.move)
		} else if action == "speak" {
			fmt.Println(name, "moves:", aCh.speak)
		} else {
			fmt.Println("WRONG ACTION!")
		}
	case Bird:
		if action == "eat" {
			fmt.Println(name, "eats:", aCh.food)
		} else if action == "move" {
			fmt.Println(name, "moves:", aCh.move)
		} else if action == "speak" {
			fmt.Println(name, "moves:", aCh.speak)
		} else {
			fmt.Println("WRONG ACTION!")
		}
	case Snake:
		if action == "eat" {
			fmt.Println(name, "eats:", aCh.food)
		} else if action == "move" {
			fmt.Println(name, "moves:", aCh.move)
		} else if action == "speak" {
			fmt.Println(name, "moves:", aCh.speak)
		} else {
			fmt.Println("WRONG ACTION!")
		}
	default:
		fmt.Println("Unknown animal type")
	}
}

// COW INTERFACE IMPLEMENTATION
func (ani Cow) Eat() {
	fmt.Println(ani.food)
}

func (ani Cow) Move() {
	fmt.Println(ani.food)
}

func (ani Cow) Speak() {
	fmt.Println(ani.food)
}

func initCow() Cow {
	var cow Cow
	cow.food = "grass"
	cow.speak = "moo"
	cow.move = "walk"
	return cow
}

// BIRD INTERFACE IMPLEMENTATION
func (ani Bird) Eat() {
	fmt.Println(ani.food)
}

func (ani Bird) Move() {
	fmt.Println(ani.food)
}

func (ani Bird) Speak() {
	fmt.Println(ani.food)
}

func initBird() Bird {
	var bird Bird
	bird.food = "worms"
	bird.speak = "peep"
	bird.move = "fly"
	return bird
}

// SNAKE INTERFACE IMPLEMENTATION
func (ani Snake) Eat() {
	fmt.Println(ani.food)
}

func (ani Snake) Move() {
	fmt.Println(ani.food)
}

func (ani Snake) Speak() {
	fmt.Println(ani.food)
}

func initSnake() Snake {
	var snake Snake
	snake.food = "mice"
	snake.speak = "slither"
	snake.move = "hiss"
	return snake
}
