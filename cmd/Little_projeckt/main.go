package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Привет, мир!") // Вывод текста

	// Пример работы с переменными, циклами и условиями
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	fmt.Printf("Сумма чисел: %d\n", sum) // Сумма чисел: 15

	// Пример использования горутин (конкурентность)
	go func() {
		fmt.Println("Это выполняется в отдельной горутине!")
	}()

	time.Sleep(1 * time.Second) // Чтобы main не завершился раньше горутин
}
