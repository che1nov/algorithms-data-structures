package main

import "fmt"

func main() {
	// массив оценок студентов
	scores := []int{85, 90, 78, 92, 88}

	// Вставка 80 на позицию с индексом 2
	scores = append(scores[:2], append([]int{80}, scores[2:]...)...)

	fmt.Println(scores) // [85, 90, 80, 78, 92, 88]

	// Удаление элемента с индексом 4
	scores = append(scores[:4], scores[4+1:]...)

	fmt.Println(scores) // [85, 90, 80, 78, 88]

	// Поиск элемента 78
	found := false
	for i, v := range scores {
		if v == 78 {
			fmt.Printf("Элемент %d найден на индексе %d\n", v, i)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Элемент не найден!")
	}
}
