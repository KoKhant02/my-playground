package main

import (
	algorithm "Sorting/Algo"
	"fmt"
)

func generateWorstCaseArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = n - i - 1 // Reverse order array (worst-case for insertion sort)
	}
	return arr
}

func main() {
	// seed := time.Now().UnixNano() // the random number generator
	// rand.Seed(uint64(seed))
	// arr := make([]int, 1000)
	// for i := 0; i < 1000; i++ {
	// 	arr[i] = rand.Intn(1000) // Generate random numbers between 0 and 1000
	// }

	// arr := []int{4, 2, 1, 3}
	// 2147483647
	n := 1000000000
	arr := generateWorstCaseArray(n)

	head := algorithm.CreateLinkedList(arr)

	fmt.Println("Original List:")
	algorithm.PrintList(head)

	sortedHead := algorithm.InsertionSortList(head)

	fmt.Println("Sorted List:")
	algorithm.PrintList(sortedHead)
}
