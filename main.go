package main

import "LeetCodePlayground/LinkedList"

func main() {
	TestLinkedList()
}

func TestLinkedList() {
	LinkedList.PrintLinkedList(LinkedList.LinkedListBuilder([]int{0, 1, 4, 5, 6, 7, 8}))
}
