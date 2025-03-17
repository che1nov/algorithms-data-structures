package main

import "fmt"

// узел списка
type DNode struct {
	value int
	next  *DNode
	prev  *DNode
}

// двусвязный список
type DoublyLinkedList struct {
	head *DNode
	tail *DNode
}

// добавление элемента в начало
func (dll *DoublyLinkedList) AddToFront(value int) {
	newNode := &DNode{value: value}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}
}

// добавление элемента в конец
func (dll *DoublyLinkedList) AddToEnd(value int) {
	newNode := &DNode{value: value}
	if dll.tail == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
}

// удаление элемента с начала
func (dll *DoublyLinkedList) RemoveFromFront() {
	if dll.head == nil {
		return
	}
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.head = dll.head.next
		dll.head.prev = nil
	}
}

// удаление элемента с конца
func (dll *DoublyLinkedList) RemoveFromEnd() {
	if dll.tail == nil {
		return
	}
	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else {
		dll.tail = dll.tail.prev
		dll.tail.next = nil
	}
}

// печать элементов списка
func (dll *DoublyLinkedList) PrintList() {
	current := dll.head
	for current != nil {
		fmt.Printf("%d <-> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	dll := &DoublyLinkedList{}

	dll.AddToFront(10)
	dll.AddToFront(20)
	dll.AddToEnd(30)
	dll.AddToEnd(40)

	dll.PrintList() // 20 <-> 10 <-> 30 <-> 40 <-> nil

	dll.RemoveFromFront()
	dll.PrintList() // 10 <-> 30 <-> 40 <-> nil

	dll.RemoveFromEnd()
	dll.PrintList() // 10 <-> 30 <-> nil
}
