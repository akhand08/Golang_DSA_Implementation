package singlylinkedlist

type Node struct {
	Data int
	next *Node
}

type SingleLinkedList struct {
	head *Node
}

// function to create a new Singly Linked List
func New() *SingleLinkedList {
	return &SingleLinkedList{}
}

// Insert new element at the end of the Singly Linked List

func (list *SingleLinkedList) Append(data int) {
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
