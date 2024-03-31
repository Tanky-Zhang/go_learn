package main

import "fmt"

// 循环使用
func main() {

	//基础使用
	var sum = 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

}
