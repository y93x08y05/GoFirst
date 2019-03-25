package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	fmt.Println(grade(0))
	fmt.Println(grade(101))
	//if contents, err := ioutil.ReadFile(filename); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}
	fmt.Println(bounded(10))
}

func grade(score int) string {
	g := ""
	switch  {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g
}

func eval(a, b int, op string) int {
	var result int
	switch op {//switch会自动break，除非使用fallthrough
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)//panic相当于报错让程序停下来
	}
	return result
}
func bounded(v int) int {
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
}
