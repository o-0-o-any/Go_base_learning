package main

import "fmt"

/*
方法的定义格式
func (接收者变量 接收者类型) 方法名(参数列表) (返回值列表) {
    函数体
}
本质就是将特定的函数绑定到特定的结构体中，这个特定的方法可以对结构体中的字段进行操作
*/

// 定义结构体
type dog struct {
	color    string
	isHealth bool
}

// 构造函数
func newDog(color string, isHealth bool) dog {
	return dog{
		color:    color,
		isHealth: isHealth,
	}
}

// 定义方法 方法是作用于特定结构体类型的函数
// 接收者表示的是调用该方法的具体类型变量 多用类型名首字母小写表示
func (d dog) printInfo() {
	if d.isHealth {
		fmt.Printf("%s is healthy\n", d.color)
	} else {
		fmt.Printf("%s is NOT healthy\n", d.color)
	}
}

// dog结构体类型的方法 治疗dog类型变量:修改isHealth为true
// 值接收者
func (d dog) quackTreat() { // 庸医
	if !d.isHealth {
		d.isHealth = true
	}
}

// 指针接收者
func (d *dog) miracleDoctorTreat() {
	if !d.isHealth {
		d.isHealth = true
	}
}

// 自定义数据类型
type INT int

// 给自定义数据类型添加方法
func (i INT) sayHello() {
	fmt.Println("Hello INT!", i)
}

func main() {
	// 构造一个结构体变量
	d1 := newDog("red", false)
	// 结构体变量.方法名 调用方法
	d1.printInfo() // red is NOT health

	d2 := newDog("blue", false)
	d2.printInfo() // red is NOT healthy
	// 调用方法 值接收者与指针接收者方法调用方式一样
	// 值接收者方法
	d2.quackTreat()
	d2.printInfo() // red is NOT healthy
	// 指针接收者方法
	d2.miracleDoctorTreat()
	d2.printInfo() // blue is healthy

	// 定义自定义数据类型变量
	var a INT = 10
	// 调用方法
	a.sayHello() // Hello INT! 10
}
