// Задание 1. Конвейер

// Цели задания:
// -Научиться работать с каналами и горутинами.
// -Понять, как должно происходить общение между потоками.

// Что нужно сделать:
// -Реализуйте паттерн-конвейер:
// Программа принимает числа из стандартного ввода в бесконечном цикле и передаёт число в горутину.
// -Квадрат: горутина высчитывает квадрат этого числа и передаёт в следующую горутину.
// -Произведение: следующая горутина умножает квадрат числа на 2.
// -При вводе «стоп» выполнение программы останавливается.

// Советы и рекомендации:
// Воспользуйтесь небуферизированными каналами и waitgroup.

// Что оценивается:
// Ввод : 3
// Квадрат : 9
// Произведение : 18

package main

import (
	"fmt"
)

func main() {

	var number int

	for {
		entryChan := make(chan int)

		fmt.Print("Введите число: ")
		fmt.Scan(&number)
		fmt.Println("Ввод:", number)

		go multiply(entryChan)
		go square(entryChan)
		entryChan <- number

	}
}

func square(squareChan chan int) {
	num := <-squareChan
	num *= num
	fmt.Println("Квадрат:", num)
	squareChan <- num
}

func multiply(multiplyChan chan int) {
	num := <-multiplyChan
	num *= 2
	fmt.Println("Произведение:", num)
	multiplyChan <- num
}
