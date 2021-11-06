package main

import (
	"fmt"
	"math"
)

func main() {
	var S float64
	pi := math.Pi

	fmt.Print("Введите площадь окружности: ")
	fmt.Scanln(&S)

	var D = math.Sqrt(S/pi) * 2
	var l = pi * D

	fmt.Printf("Диаметр окружности: %f\n", D)
	fmt.Printf("Длина окружности: %f\n", l)

}

//add commit
