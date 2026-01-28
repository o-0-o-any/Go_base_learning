package main

import (
	"fmt"
)

// 引出接口的实例 类似面向对象编程中的多态
/*
接口的定义格式如下
type 接口类型名 interface {
  方法名1(参数列表1) 返回值列表1
  方法名2(参数列表2) 返回值列表2
  ...
}
*/
// 定义一个接口
type caller interface {
	call() // 只要实现了call方法的类型都是caller类型
}

// 定义三个结构体
type cat2 struct {
}
type dog struct {
}
type person struct {
}

// 分别给这三个结构体类型绑定一个同名的方法
func (c cat2) call() {
	fmt.Println("cat:喵喵喵")
}
func (d dog) call() {
	fmt.Println("dog:汪汪汪")
}
func (p person) call() {
	fmt.Println("person:啊啊啊")
}

// 根据传入参数的不同调用不同结构体的call方法 实现类似多态的功能
func call(x caller) {
	x.call()
}

func main() {
	// 分别声明一个结构体变量
	var c1 cat2
	var d1 dog
	var p1 person
	// 逐个调用对应的方法进行执行
	c1.call() // cat:喵喵喵
	d1.call() // dog:汪汪汪
	p1.call() // person:啊啊啊
	// 使用接口完成类似多态的功能 将三个类型的call当成一个caller来处理
	call(c1) // cat:喵喵喵
	call(d1) // dog:汪汪汪
	call(p1) // person:啊啊啊
}
