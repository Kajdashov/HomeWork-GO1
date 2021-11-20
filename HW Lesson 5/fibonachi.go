package main

import "fmt"

func fibonachi(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonachi(n-1) + fibonachi(n-2)
}
func main() {
	fmt.Println(fibonachi(2))
}
