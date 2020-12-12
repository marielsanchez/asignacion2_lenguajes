package main

import "fmt"

//aqui el main

func main() {

	slice := GenerateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	insertionsort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice)
}
