package main

import (
	"fmt"
)

type animal interface {
	eat()
}
type cat struct {
}
type chicken struct {
}

func (c cat) eat() {
	fmt.Println("猫吃老鼠")
}
func (c chicken) eat() {
	fmt.Println("鸡吃虫")
}

func main() {
	var a animal          // 定义一个接口类型的变量
	fmt.Printf("%T\n", a) // <nil>

	// 声明两个结构体变量
	var b = cat{}
	var c = chicken{}

	// 将b赋值给a
	a = b
	a.eat()               // main.cat
	fmt.Printf("%T\n", a) // main.cat

	// 再将c赋值给a
	a = c
	a.eat()               // 鸡吃虫
	fmt.Printf("%T\n", a) // main.chicken

}
