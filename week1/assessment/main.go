package main

import "fmt"

func main() {
	integers := [10]int{}

	fmt.Println("Enter 10 integers: ")
	for i := 0; i < 10; i++ {
		fmt.Scan(&integers[i])
	}

	BubbleSort(integers[:])
	fmt.Println("Sorted integers: ", integers)
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

func Swap(arr []int, i int) {
	arr[i], arr[i+1] = arr[i+1], arr[i]
}
