package algorithm

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func InsertionSortList(head *ListNode) *ListNode {

	// Create a dummy node that will serve as the start of the sorted part of the list
	dummy := &ListNode{}
	dummy.Next = head

	// Iterate through the original list
	current := head

	for current != nil && current.Next != nil {

		prev := dummy
		// Find the previous node of the sorted part where current.Next needs to be inserted
		for prev.Next != nil && prev.Next.Val < current.Next.Val {
			prev = prev.Next
		}

		// If the current node needs to be inserted between prev and prev.Next
		if prev.Next != current.Next {
			// Remove current.Next from its original position
			temp := current.Next
			current.Next = temp.Next // Skip over the node we're removing
			temp.Next = prev.Next    // Insert it into the sorted part
			prev.Next = temp         // Update the sorted list to include the inserted node
		} else {
			// Move to the next node if no insertion was necessary
			current = current.Next
		}

	}

	return dummy.Next
}

func CreateLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	dummy := &ListNode{}
	curr := dummy
	for _, val := range arr {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return dummy.Next
}

func PrintList(head *ListNode) {
	if head == nil {
		fmt.Println("Empty list")
		return
	}
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}
