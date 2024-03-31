package main

import (
	"fmt"
	"unsafe"
)

// QUAN_JU 定义全局变量
var QUAN_JU = 1

// 全局赋值多个
var (
	r1 = 2
	r2 = "hhhh"
)

// 这是第一个文件，打印的使用
func main() {

	//1、打印的使用
	fmt.Println("hello")
	println("hello world!!!")

	fmt.Println("-----------------------------------------------------")

	//2、变量的使用
	//Go语言的类型： bool、string、int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、uintptr、byte // uint8 的别名、rune // int32 的别名 代表一个 Unicode 码
	//float32、float64
	//complex64、complex128
	var name string = "zhangsan"
	println(name)

	//不赋值就默认是0
	var c int
	fmt.Println("c=", c)

	//可以进行类型推断
	var d = 1.2
	fmt.Println("d=", d)

	//可以省略var
	e := 8
	fmt.Println("e=", e)

	//一次赋值多个
	var n1, n2, n3 int
	println("n1=", n1, "n2=", n2, "n3=", n3)

	//不用var
	n5, n6, n7 := 100, "lisi", 4
	println("n5=", n5, "n6=", n6, "n7=", n7)

	var a = 100
	var b = 200
	//两个数字交换，就是这么简单
	a, b = b, a
	fmt.Println(a, b)

	//查看一个变量的类型和占用的字节大小
	var m int = 10
	fmt.Printf("m的数据类型是%T,占用的字节数是%d", m, unsafe.Sizeof(m))

}

//func GetData() (int,int){
//	return 100,200
//}
