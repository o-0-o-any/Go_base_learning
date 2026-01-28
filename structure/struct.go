package main

import "fmt"

/*
结构体定义格式
type 类型名 struct {
    变量名 变量类型
    变量名 变量类型
    ...
}
*/

// 定义一个结构体
type person struct {
	name     string
	age      int
	gender   string
	identity int
	hobby    []string
	OnSchool bool
}

func main() {
	// 声明一个结构体变量
	var p1 person
	// 使用结构体变量中的属性方式为 结构体变量名.属性名
	// 通过字段赋值 与前面学的简单/复杂数据类型赋值一样
	p1.name = "张三"
	p1.age = 18
	p1.gender = "male"
	p1.identity = 411425202609189823
	p1.hobby = make([]string, 0, 10)
	p1.hobby = []string{"玩dead cells", "学习Go", "学习Gin+Gorm", "学习Docker"}
	p1.OnSchool = true
	// 访问字段数据
	fmt.Println(p1.name, p1.age, p1.gender, p1.identity, p1.hobby, p1.OnSchool)
	/*
		匿名函数定义
		func (函数参数) 函数返回值 {
		    函数体
		} (传参)
		匿名结构体定义
		var 结构体实例名 struct{
		    字段名 字段类型
		    字段名 字段类型
		    ......
		}
	*/
	// 声明一个匿名结构体 临时使用
	var t1 struct {
		name   string
		age    int
		gender string
		score  map[string]int
	}
	// 给字段赋值
	t1.name = "李四"
	t1.age = 18
	t1.gender = "male"
	t1.score = make(map[string]int, 10)
	t1.score["数据结构"] = 80
	t1.score["Python程序设计"] = 80
	// 访问字段名
	fmt.Println(t1.name, t1.age, t1.gender, t1.score)
}
