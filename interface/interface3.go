package main

import "fmt"

// 前面写的都是多个类型实现一个接口
// 这里学习一个类型实现多个接口  和 接口嵌套

// 接口嵌套 显示实现interface3接口的结构体 必须实现其中包含的接口
type interface3 interface {
	interface1
	interface2
}

type interface1 interface {
	testFunc3()
}
type interface2 interface {
	testFunc4()
}

type testStruct struct {
	name string
}

// testStruct实现了interface1接口
func (t testStruct) testFunc3() {
	fmt.Println("testStruct testFunc3")
}

// testStruct实现了interface2接口
func (t testStruct) testFunc4() {
	fmt.Println("testStruct testFunc4")
}

// testStruct实现了多个接口
