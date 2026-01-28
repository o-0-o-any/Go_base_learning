package main

import "fmt"

// 构造函数

// 定义一个结构体
type person2 struct {
	name string
	age  int
}

// 构造函数 返回一个结构体变量的函数 通常函数名用new开头
// 返回值决定返回的是结构体变量还是结构体指针变量
// 如果是*person2就是结构体指针变量 person2就是结构体变量
// 当结构体比较大时 尽量使用结构体指针变量 减少程序的开销
func newPerson(name string, age int) person2 {
	return person2{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("张三", 18)
	p2 := newPerson("李四", 19)
	fmt.Println(p1, p2) // {张三 18} {李四 19}
}
