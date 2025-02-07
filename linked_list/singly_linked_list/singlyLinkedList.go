package singlylinkedlist

import "fmt"

type Node struct {
	Data int
	next *Node
}

type SinglyLinkedList struct {
	head *Node
}

// function to create a new Singly Linked List
func New() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// Insert new element at the end of the Singly Linked List
func (list *SinglyLinkedList) Append(data int) {
	newNode := &Node{Data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode

}

// Insert new element at the beginning of the Singly Linked List
func (list *SinglyLinkedList) Prepend(data int) {
	newNode := &Node{Data: data, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	newNode.next = list.head
	list.head = newNode
}

// Add new element on the given index
func (list *SinglyLinkedList) AddByIndex(data int, index int) error {

	if index < 0 {
		return fmt.Errorf("Index out of Range")
	}

	counter := 0
	currentNode := list.head

	for counter < index && currentNode != nil {
		currentNode = currentNode.next
		counter += 1
	}

	newNode := &Node{Data: data, next: nil}

	if currentNode == nil && index == 0 {
		list.head = newNode
	} else if currentNode == nil && counter < index-1 {
		return fmt.Errorf("Index out of Range")
	} else if index == 0 {

		newNext := list.head
		list.head = newNode
		newNode.next = newNext

	} else {
		oldNext := currentNode.next
		currentNode.next = newNode
		newNode.next = oldNext
	}

	return nil

}
