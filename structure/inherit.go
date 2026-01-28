package main

import "fmt"

// 结构体模拟实现其他语言中的继承
type animal struct {
	name string
}

// 给animal结构体实现一个跑的方法
func (a animal) run() {
	fmt.Printf("%s会跑!", a.name)
}

// 猫结构体
type cat struct {
	animal // 模拟继承了animal中的字段与方法
}

// 给cat结构体实现一个叫的方法
func (c cat) call() {
	fmt.Println("喵喵喵!")
}

func main() {
	// 声明一个cat结构体变量
	c1 := cat{
		animal: animal{
			name: "小喵",
		},
	}
	c1.call() // 喵喵喵!
	c1.run()  // 小喵会跑! 调用模拟继承的方法
}
