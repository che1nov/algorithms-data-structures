package main

import "fmt"

// Определение узла дерева поиска
type TreeNode struct {
	Key   int
	Left  *TreeNode
	Right *TreeNode
}

// Функция для вставки узла в дерево поиска
func (t *TreeNode) Insert(key int) {
	if key < t.Key {
		if t.Left == nil {
			t.Left = &TreeNode{Key: key}
		} else {
			t.Left.Insert(key)
		}
	} else {
		if t.Right == nil {
			t.Right = &TreeNode{Key: key}
		} else {
			t.Right.Insert(key)
		}
	}
}

// Функция для обхода дерева в порядке возрастания ключей
func (t *TreeNode) InOrder() {
	if t != nil {
		t.Left.InOrder()
		fmt.Print(t.Key, " ")
		t.Right.InOrder()
	}
}

func main() {
	root := &TreeNode{Key: 10}
	keys := []int{5, 15, 3, 7, 12, 18}

	for _, key := range keys {
		root.Insert(key)
	}

	fmt.Print("Ключи в порядке возрастания: ")
	root.InOrder() // Вывод: 3 5 7 10 12 15 18
}
