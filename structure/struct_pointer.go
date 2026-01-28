package main

import "fmt"

/*
定义结构体指针
var 结构体指针变量 = new(结构体名)
*/

// 定义一个结构体
type student struct {
	name   string
	age    int
	gender string
	score  map[string]int
}

// 结构体是值类型不是引用类型
// 跟C/C++一样 参数值传递 无法修改原本变量的值 仅仅是将调用时传入的变量的值传递过来了并不是变量本身
func addGender(s student) {
	s.gender = "female"
}

// 通过地址传递 传递的变量的地址 再次通过解引用对变量本身进行操作
func addGender2(s *student) {
	(*s).gender = "male"
}

func addGender3(s *student) {
	s.gender = "female"
}

func main() {
	var s1 student
	s1.name = "王五"
	s1.age = 18
	s1.score = make(map[string]int, 10)
	s1.score["大学英语"] = 60
	// 值传递
	addGender(s1)
	// 打印结构体内容
	fmt.Println(s1.name, s1.age, s1.score, s1.gender) // 王五 18 map[大学英语:60]  s1.gender字段是空的 通过值传递函数无法修改字段

	// 地址传递
	addGender2(&s1)
	fmt.Println(s1.name, s1.age, s1.score, s1.gender) // 王五 18 map[大学英语:60] male 通过地址传递的方式可以修改 与C/C++性质一样

	// 使用new关键字定义结构体指针
	var p2 = new(student) // 申请student类型的内存
	p2 = &s1              // 指针指向s1
	// 通过指针访问结构体变量的字段名 可以直接使用指针变量名+.+字段名 也可以先对指针解引用然后+.+字段名
	fmt.Println(p2.name, p2.age, p2.score, p2.gender) // 王五 18 map[大学英语:60] male
	// 将指针作为参数传递给函数 也是地址传递
	addGender3(p2)
	fmt.Println(p2.name, p2.age, p2.score, p2.gender) // 王五 18 map[大学英语:60] female

	// 结构体初始化
	// 结构体指针变量直接指向结构体变量
	var s2 student
	p3 := &s2
	p3.name = "Join"
	fmt.Printf("%T\n", p3) // *main.student
	fmt.Println(s2.name)   // Join
	// 使用new直接为变量分配内存得到结构体指针变量
	p4 := new(student)
	// 两种引用字段名的方式
	p4.name = "York"
	(*p4).age = 19
	fmt.Printf("%T\n", p4)       // *main.student
	fmt.Println(p4.name, p4.age) // York 19
	// key-value初始化
	var p5 = &student{
		name: "Yss",
		age:  18,
	}
	fmt.Printf("%T\n", p5)       // *main.student
	fmt.Println(p5.name, p5.age) // Yss 18
	// 省略掉key值 直接按照定义顺序给字段名赋值  必须全部赋值
	var p6 = &student{
		"Jie",
		18,
		"male",
		map[string]int{"数据结构": 90},
	}
	fmt.Printf("%T\n", p6)                            // *main.student
	fmt.Println(p6.name, p6.age, p6.gender, p6.score) // Jie 18 male map[数据结构:90]

	// 结构体占用一块连续的内存
	var s3 = student{
		"Jie",
		18,
		"male",
		map[string]int{"数据结构": 90},
	}
	fmt.Printf("%p %p %p %p\n", &s3.name, &s3.age, &s3.score, &s3.gender) // 0xc000104e40 0xc000104e50 0xc000104e68 0xc000104e58
}
