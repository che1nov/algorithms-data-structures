package main

import (
	"errors"
	"fmt"
)

// Stack определяем тип stack как срез целых чисел
type Stack []int

// Push добавляет элемент в стек
func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

// Pop извлекает верхний элемент из стека
func (s *Stack) Pop() (int, error) {
	if len(*s) == 0 {
		return 0, errors.New("стек пустой")
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res, nil
}

// Peek возвращает верхний элемент без его удаления
func (s *Stack) Peek() (int, error) {
	if len(*s) == 0 {
		return 0, errors.New("стек пустой")
	}
	return (*s)[len(*s)-1], nil
}

// IsEmpty проверяет, пуст ли стек
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func main() {
	var stack Stack

	stack.Push(1)
	stack.Push(10)
	stack.Push(22)

	peek, _ := stack.Peek()
	fmt.Printf("Верхняя часть стека: %v\n", peek)
	fmt.Printf("Содержимое стека: %v\n", stack)

	popped, _ := stack.Pop()
	fmt.Printf("Верхний элемент стека: %v\n", popped)
	fmt.Printf("Стек после применения метода Pop(): %+v\n", stack)

}
