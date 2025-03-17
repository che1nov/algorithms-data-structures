package main

import "fmt"

// Node определение узла бинарного дерево
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert функция для вставки нового значения в бинарное дерево
func Insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{Value: value}
	}
	if value < node.Value {
		node.Left = Insert(node.Left, value)
	} else {
		node.Right = Insert(node.Right, value)
	}
	return node
}

// InOrder функция для обхода дерева в симметричном порядке
func InOrder(node *Node) {
	if node != nil {
		InOrder(node.Left)
		fmt.Print(node.Value, " ")
		InOrder(node.Right)
	}
}

func main() {
	var root *Node
	values := []int{5, 3, 8, 1, 4}

	for _, value := range values {
		root = Insert(root, value)
	}

	fmt.Print("Симметричный обход дерева: ")
	InOrder(root) // Вывод: 1 3 4 5 8
}
