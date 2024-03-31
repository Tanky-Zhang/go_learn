package main

import "fmt"

func main() {

	//字符串使用
	var address string = "北京-长城"
	fmt.Println(address)

	address = "hello"
	fmt.Println(address)

	str2 := "abc\nabc"
	fmt.Println(str2)
	//使用反引号可以输出转义符
	str2 = `abc\nabc`
	fmt.Println(str2)

	var name = "hello " + "world!"
	fmt.Println(name)

	fmt.Println("-----------------------------------------------------")

	//基本数据类型的转换都要显示转换
	var a = 10
	var b float32 = float32(a)
	var c int64 = int64(a)
	fmt.Printf("a=%v,b=%v,c=%v", a, b, c)

	fmt.Println("-----------------------------------------------------")

}
