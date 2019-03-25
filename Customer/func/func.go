package main

import "fmt"

func main() {
	fmt.Println(eval(3, 4, "*"))
	q, r := div(13, 3)
	fmt.Println(q, r)
	fmt.Println(sum(1, 2, 3, 4, 5))
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers{
		s += numbers[i]
	}
	return s
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operator:" + op)
		
	}
}
