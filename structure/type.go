package main

import "fmt"

/*
使用type关键字 定义新的类型
type MyInt int
MyInt就是一种新的类型 它具有int的特性
*/

type MyInt int // 自定义数据类型

/*
给类型起别名
type 类型别名 = 类型
type YourInt = int
*/
type YourInt = int // 给类型起别名

func main() {

	var a MyInt = 100
	fmt.Printf("%T %d\n", a, a) // main.MyInt 100   main.MyInt意思就是main包里的MyInt类型 是一种新的类型

	var b YourInt = 100         // 用类型别名声明变量  本质还是原本的数据类型
	fmt.Printf("%T %d\n", b, b) // int 100
}
