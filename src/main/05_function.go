package main

import "fmt"

// 函数的使用
func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sumAndDown(5, 4))

	//函数也是一种类型,可以赋值给一个变量
	a := sum
	fmt.Printf("a对应的类型是%T", a)

	//匿名函数
	result := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 30)
	fmt.Println(result)

}

// init函数，在最开始的时候调用
func init() {
	fmt.Println("hello golang")
}

// 两数相加函数
func sum(num1 int, num2 int) int {
	return num2 + num1
}

// 相加相减操作函数-两个返回值
func sumAndDown(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}
