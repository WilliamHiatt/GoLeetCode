package LinkedList

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type ReverseLL struct {
	Head       *ListNode
	ReturnNode *ListNode
}

func LinkedListBuilder(arr []int) *ListNode {
	headNode := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	prevNode := headNode

	for _, val := range arr[1:] {
		currNode := &ListNode{
			Val:  val,
			Next: nil,
		}

		prevNode.Next = currNode
		prevNode = currNode
	}

	return headNode
}

func PrintLinkedList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}
