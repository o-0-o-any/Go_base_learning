package main

import "fmt"

/*
空接口定义
type 接口名 interface {
  // 接口内什么都不写
}
*/

// 空接口作为函数参数
func interfaceFunc(a interface{}) {
	fmt.Printf("传入的值类型:%T 值:%v\n", a, a)
}

func main() {
	// map类型中使用空接口
	// interface{}为空接口 any为interface{}的别名 在源码中有定义
	// 空接口作为map的键和值
	var a = make(map[string]interface{}, 10)
	// 也可以使用any
	var b = make(map[any]int, 10)

	// 空接口可以接收任何类型的值 任何类型都是空接口类型
	a["张三"] = 100
	a["李四"] = "不到一千"
	a["王五"] = map[string]int{"HelloWorld": 1}

	// b[map[string]int{"Hello.go": 5}] = 10   // 不合法 map的键必须是可比较类型 map类型不可比较
	// b[[]int{1, 2, 3}] = 3    // // 不合法 map的键必须是可比较类型 slice类型不可比较
	b["你好"] = 10
	b[[...]int{1, 3, 5, 7}] = 10

	fmt.Println(a) // map[张三:100 李四:不到一千 王五:map[HelloWorld:1]]
	fmt.Println(b) // map[你好:10 [1 3 5 7]:10]

	interfaceFunc(false)                           // 传入的值类型:bool 值:false
	interfaceFunc(a)                               // 传入的值类型:map[string]interface {} 值:map[张三:100 李四:不到一千 王五:map[HelloWorld:1]]
	interfaceFunc(b)                               // 传入的值类型:map[interface {}]int 值:map[你好:10 [1 3 5 7]:10]
	interfaceFunc("Hello World")                   // 传入的值类型:string 值:Hello World
	interfaceFunc(12)                              // 传入的值类型:int 值:12
	interfaceFunc([...]int{1, 2, 3, 7})            // 传入的值类型:int 值:12
	interfaceFunc([]string{"hello", "world", "!"}) // 传入的值类型:[]string 值:[hello world !]
}
