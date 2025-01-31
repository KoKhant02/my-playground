package main

import (
	algorithm "Sorting/Algo"
	"fmt"
)

func main() {

	//INSERTION SORTING

	insertionArr := []int{4, 2, 1, 3}

	head := algorithm.CreateLinkedList(insertionArr)
	sortedHead := algorithm.InsertionSortList(head)

	fmt.Println("\nInsertion Sort List:")
	algorithm.PrintList(sortedHead)

	//REVERSE SORTING
	fmt.Println("\nReverse List:")
	fmt.Println(algorithm.Reverse(123))
	fmt.Println(algorithm.Reverse(-123))
	fmt.Println(algorithm.Reverse(120))
	fmt.Println(algorithm.Reverse(2147483647))

	//SELECTION SORTING
	selectionArr := []int{64, 25, 12, 22, 11}
	algorithm.SelectionSort(selectionArr)
	fmt.Println("Sorted array:", selectionArr)

	//QUICK SORTING
	QuickArr := []int{89, 23, 65, 2, 10}
	sortedArr := algorithm.QuickSort(QuickArr)
	fmt.Println("Sorted array:", sortedArr)
}
