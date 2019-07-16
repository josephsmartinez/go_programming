package main

import "fmt"

type Node struct {
	next *Node
	data interface{}
}

type List struct {
	head *Node
}

func (L *List) Insert(data interface{}) {

	newNode := Node{nil, data}
	newNode.next = nil

	if L.head == nil {
		L.head = &newNode
	} else {
		last := L.head.next
		for last != nil {
			last = last.next
		}
		last.next = &newNode
	}

}

func (L *List) Display() {

	list := L.head
	for list != nil {
		fmt.Printf("%v\n", list.data)
	}

}

func main() {

	// Single Linked List - sll
	sll := List{}
	sll.Insert(5)
	sll.Insert(9)
	sll.Insert(13)
	sll.Insert(22)
	sll.Insert(28)
	sll.Insert(36)
	//sll.Display()
}
