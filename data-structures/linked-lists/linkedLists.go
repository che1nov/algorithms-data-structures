package main

import "fmt"

// узел списка
type Node struct {
	value int
	next  *Node
}

// связанный список
type LinkedList struct {
	head *Node
}

// Добавление элемента в начало списка
func (ll *LinkedList) AddToFront(value int) {
	newNode := &Node{value: value}
	newNode.next = ll.head
	ll.head = newNode
}

// добавление элемента в конец списка
func (ll *LinkedList) AddToEnd(value int) {
	newNode := &Node{value: value}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// ддаление первого элемента
func (ll *LinkedList) RemoveFromFront() {
	if ll.head != nil {
		ll.head = ll.head.next
	}
}

// печать элементов списка
func (ll *LinkedList) PrintList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}
func main() {
	ll := &LinkedList{}

	ll.AddToFront(10)
	ll.AddToFront(20)
	ll.AddToEnd(30)

	ll.PrintList() // 20 -> 10 -> 30 -> nil

	ll.RemoveFromFront()

	ll.PrintList() // 10 -> 30 -> nil
}
