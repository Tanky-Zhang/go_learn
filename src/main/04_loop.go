package main

import (
	"fmt"
)

// 循环使用
func main() {

	//基础使用
	var sum = 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	//遍历字符串，这里是对字节进行遍历，不支持中文的遍历，后边可以使用切片
	var str string = "hello Golang"
	for i := 0; i < len(str); i++ {
		//如果不实用%c则会使用uncode码值
		fmt.Printf("%c\n", str[i])
	}

	//使用for-range进行遍历
	var str2 string = "你好 Golang"
	for k, value := range str2 {
		fmt.Printf("索引为%d，值是%c\n", k, value)
	}

	//break使用，和Java没有区别，但是有个标签的功能可以结束到外层循环
label1:
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Println(i, j)
			if i == j {
				break label1
			}
		}
	}

	//goto的使用
	var name = "lisi"
	fmt.Println(name)
	if len(name) > 4 {
		goto label2
	}
	fmt.Println("未跳转")
label2:
	fmt.Println("跳转位置")

}
