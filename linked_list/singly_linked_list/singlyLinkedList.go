package singlylinkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type SinglyLinkedList struct {
	Head *Node
}

// function to create a new Singly Linked List
func New() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// Insert new element at the end of the Singly Linked List
func (list *SinglyLinkedList) Append(data int) {
	newNode := &Node{Data: data, Next: nil}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	current := list.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode

}

// Insert new element at the beginning of the Singly Linked List
func (list *SinglyLinkedList) Prepend(data int) {
	newNode := &Node{Data: data, Next: nil}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	newNode.Next = list.Head
	list.Head = newNode
}

// Add new element on the given index
func (list *SinglyLinkedList) AddByIndex(data int, index int) error {

	if index < 0 {
		return fmt.Errorf("Index out of Range")
	}

	counter := 0
	currentNode := list.Head

	for counter < index && currentNode != nil {
		currentNode = currentNode.Next
		counter += 1
	}

	newNode := &Node{Data: data, Next: nil}

	if currentNode == nil && index == 0 {
		list.Head = newNode
	} else if currentNode == nil && counter < index-1 {
		return fmt.Errorf("Index out of Range")
	} else if index == 0 {

		newNext := list.Head
		list.Head = newNode
		newNode.Next = newNext

	} else {
		oldNext := currentNode.Next
		currentNode.Next = newNode
		newNode.Next = oldNext
	}

	return nil

}

func (list *SinglyLinkedList) Size() int {
	counter := 0
	currentNode := list.Head

	for currentNode != nil {
		counter++
		currentNode = currentNode.Next
	}

	return counter
}
