package search
/*
 bool string
 int8 int16 int32 int64
 Go语言有指针uintptr
 byte是8位
 rune即char类型是32位
 float32 float64
 complex64指的是实部和虚部分别为32位
 complex128指的是实部和虚部分别为64位
 复数：3 + 4i 实部为X轴虚部为Y轴
 变量类型写在变量名后
 Go语言没有隐式类型转换，只有强制类型转换
 for，if后面的条件没有括号
 if条件里也可以定义变量
 没有while
 switch不需要break，可以直接switch多个条件
 Go语言的指针不能运算
 Go语言只有值传递一种方式
 */

func love(a, b int) (int, int) {
	a++
	b++
	return a, b
}