package main

import "fmt"

func main() {
	var n int16
	fmt.Print("Введите трехзначное число: ")
	fmt.Scan(&n)
	var hundreds int16 = n / 100
	var dozens int16 = (n - (hundreds * 100)) / 10
	var units int16 = n - (hundreds * 100) - (dozens * 10)
	fmt.Printf("Сотни: %d\n", hundreds)
	fmt.Printf("Десятки: %d\n", dozens)
	fmt.Printf("Единицы: %d\n", units)
}
