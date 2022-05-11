// Diberikan head dari singly linked list, balikkan (reverse) list, dan kembalikan daftar yang sudah dibalik.
//
// Contoh 1
// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]
//
// Contoh 2
// Input: head = [1,2]
// Output: [2,1]
//
// Contoh 3
// Input: head = []
// Output: []

package main

import "fmt"

// Definisi untuk singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node1 := new(ListNode)
	node1.Val = 1

	node2 := new(ListNode)
	node2.Val = 2
	node1.Next = node2

	node3 := new(ListNode)
	node3.Val = 3
	node2.Next = node3

	node4 := new(ListNode)
	node4.Val = 4
	node3.Next = node4

	node5 := new(ListNode)
	node5.Val = 5
	node4.Next = node5

	reverseList(node1)
}

func reverseList(head *ListNode) *ListNode {
	var front *ListNode
	curr, next := head, head
	for curr != nil {
		//fmt.Println(curr)
		next = curr.Next
		fmt.Println(front, curr, next)

		curr.Next = front
		fmt.Println(front)
		//front, mid = curr, next
		front = curr
		curr = next
	}
	return front
}
