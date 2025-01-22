package main

import (
	algorithm "Sorting/Algo"
	"fmt"
)

func main() {

	//INSERTION SORTING

	arr := []int{4, 2, 1, 3}

	head := algorithm.CreateLinkedList(arr)
	sortedHead := algorithm.InsertionSortList(head)

	fmt.Println("\nInsertion Sort List:")
	algorithm.PrintList(sortedHead)

	//REVERSE SORTING
	fmt.Println("\nReverse List:")
	fmt.Println(algorithm.Reverse(123))
	fmt.Println(algorithm.Reverse(-123))
	fmt.Println(algorithm.Reverse(120))
	fmt.Println(algorithm.Reverse(2147483647))
}
