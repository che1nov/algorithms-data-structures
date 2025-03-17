package main

import (
	"errors"
	"fmt"
)

// Queue определяем как срез целых чисел
type Queue []int

// Enqueue добавляет элемент в конец очереди
func (q *Queue) Enqueue(v int) {
	*q = append(*q, v)
}

// Dequeue удаляет элемент из начала очереди
func (q *Queue) Dequeue() (int, error) {
	if len(*q) == 0 {
		return 0, errors.New("очередь пуста")
	}
	res := (*q)[0]
	*q = (*q)[1:]
	return res, nil
}

// Peek возвращает первый элемент без его удаления
func (q *Queue) Peek() (int, error) {
	if len(*q) == 0 {
		return 0, errors.New("очередь пуста")
	}
	return (*q)[0], nil
}

// IsEmpty проверяет, пуста ли очередь
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func main() {
	var queue Queue

	queue.Enqueue(1)
	queue.Enqueue(10)
	queue.Enqueue(22)

	peek, _ := queue.Peek()
	fmt.Printf("Первый элемент очереди: %v\n", peek)
	fmt.Printf("Содержимое очереди: %v\n", queue)

	dequeued, _ := queue.Dequeue()
	fmt.Printf("Удалённый элемент из очереди: %v\n", dequeued)
	fmt.Printf("Очередь после применения метода Dequeue(): %+v\n", queue)
}
