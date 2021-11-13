package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	calculator()
}

func calculator() {
	fmt.Println("Hello calculator!")

	fmt.Println("Введите оператор: ")
	var operator string // "+", "-", "*", "/"

	if _, err := fmt.Scan(&operator); err != nil {
		fmt.Println("Ошибка!", err.Error())
		os.Exit(1)
	}

	fmt.Println("Введите операнд1 и операнд2: ")
	var operand1, operand2 float64
	
	if _, err := fmt.Scan(&operand1, &operand2); err != nil {
		fmt.Println("Ошибка! ", err.Error())
		os.Exit(1)
	} else if {fmt.Scan(operand1)}

	var result float64
	switch operator {
	case "Exit":
		os.Exit(1)
	case "+":
		result = operand1 + operand2
	case "Корень":
		result = math.Sqrt(operand1)
	case "^":
		result = operand1 * operand1
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			fmt.Println("Делить на ноль нельзя")
			os.Exit(1)
		}

		result = operand1 / operand2 // todo:

	default:
		fmt.Println("Не поддерживаемая операция")
		os.Exit(1)
	}
	fmt.Println(result)

}
