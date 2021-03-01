package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var a, b string
	var res float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	newA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		log.Fatalf("Ошибка проверки числа %s: %v", a, err)
	}

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	newB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		log.Fatalf("Ошибка проверки числа %s: %v", b, err)
	}

	fmt.Print("Введите арифметическую операцию (+, -, *, /, **): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = newA + newB
	case "-":
		res = newA - newB
	case "*":
		res = newA * newB
	case "/":
		res = newA / newB
	case "**":
		bb := int(newB)
		firstA := newA
		for i := 0; i < bb-1; i++ {
			newA = newA * firstA
		}
		res = newA
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
}
