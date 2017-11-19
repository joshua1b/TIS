package main

import (
	"fmt"
)

type solution struct{}

func main() {
	solution := solution{}
	a, b := solution.Read()
	result := solution.Process(a, b)
	solution.Print(result)
}

func (solution) Read() (int, int) {
	var a, b int
	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%d\n", &b)
	return a, b
}

func (solution) Process(a, b int) int {
	return a + b
}

func (solution) Print(a int) {
	fmt.Println(a)
}
