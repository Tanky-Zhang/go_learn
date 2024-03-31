package main

import "fmt"

// 流程控制学习
func main() {

	//年龄大于18岁才能开车
	//var age int
	//fmt.Scanln(&age)
	//if age > 18 {
	//	fmt.Printf("年龄为%v,可以开车\n", age)
	//} else {
	//	fmt.Printf("年龄为%v,不可以开车\n", age)
	//}
	//
	////可以在if后赋值
	//if count := 20; count < 30 {
	//	fmt.Println("口罩库存不足")
	//}

	//case switch使用，fallthrough可以进行穿透
	var score = 83
	switch score / 10 {
	case 10:
		fmt.Println("您的成绩为A")
	case 9:
		fmt.Println("您的成绩为B")
		//穿透到下一个分支
		fallthrough
	default:
		fmt.Println("您的成绩为C")
	}
}
