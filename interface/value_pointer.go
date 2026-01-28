package main

import "fmt"

// 使用值接收者和指针接收者实现接口的区别
type tester interface {
	testFunc1()
	testFunc2()
}

type testStruct1 struct {
	name string
}
type testStruct2 struct {
	name string
}

// 使用值接收者实现接口的所有方法
func (ts testStruct1) testFunc1() {
	fmt.Println("testStruct1.testFunc1")
}
func (ts testStruct1) testFunc2() {
	fmt.Println("testStruct1.testFunc2")
}

// 使用指针接收者实现接口的所有方法
func (ts *testStruct2) testFunc1() {
	fmt.Println("testStruct2.testFunc1")
}
func (ts *testStruct2) testFunc2() {
	fmt.Println("testStruct2.testFunc2")
}

func main() {
	// 定义一个接口类型的变量
	var tr1 tester
	// 定义结构体1变量 使用值接收者实现接口的所有方法
	tt1 := testStruct1{ // testStruct1类型
		name: "test_name1",
	}
	tt2 := &testStruct1{ // *testStruct1类型
		name: "test_name2",
	}
	tr1 = tt1
	fmt.Println(tr1) // {test_name1}
	tr1 = tt2
	fmt.Println(tr1) // {test_name2}

	// 定义结构体2变量 使用指针接收者实现接口的所有方法
	tt3 := testStruct2{ // testStruct2类型
		name: "test_name1",
	}
	fmt.Println(tt3)
	tt4 := &testStruct1{ // *testStruct2类型
		name: "test_name2",
	}
	/*  非法 实现接口方法的是testStruct2的指针类型 这里传入的是类型 只能取地址
	tr1 = tt3
	fmt.Println(tr1)
	*/
	tr1 = tt4
	fmt.Println(tr1)
}
