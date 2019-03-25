package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func main()  {
	fmt.Println("Hello World")
	variableZeroValue()
	variableTypeDeduction()
	fmt.Println(aa, ss, true)
	euler()
	triangle()
	consts()
	enums()
}

func enums()  {//枚举类型
	const(
		cpp = 0
		java = 1
		python = 2
		golang = 3
	)
	const(
		scala = iota // 表示自增值
		ruby
		javascript
	)
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
	fmt.Println(scala, ruby, javascript)
}
func consts()  {
	const filename  = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}
func triangle()  {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}
func euler()  {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	/*
	 欧拉公式的应用
	 */
	fmt.Printf("%.3f\n", cmplx.Exp(1i * math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
	fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
}
func variableZeroValue()  {
	var a, b = 3, 4
	var s = "abc"
	fmt.Printf("%d %q %d\n", a, s, b)
}

func variableTypeDeduction()  {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}