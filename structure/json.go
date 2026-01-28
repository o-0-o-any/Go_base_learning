package main

import (
	"encoding/json"
	"fmt"
)

/* 结构体和JSON
1、把Go语言中的结构体变量 转换成 JSON格式的字符串
2、把JSON格式的字符串 转换成 Go语言中的结构体变量
*/

// 定义一个结构体
type test2 struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	// 定义一个结构体变量并初始化
	t1 := test2{
		Name:   "李四",
		Age:    18,
		Gender: "male",
	}
	// 序列化 将结构体变量转换成JSON格式的字符串
	// 需要注意的是 json.Marshal仅会将结构体中首字母大小的字段导出
	content, err := json.Marshal(t1)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(string(content)) // {"Name":"李四","Age":18,"Gender":"male"}

	// 反序列化
	str := `{"name":"李四", "age":18, "gender":"male"}`
	var t2 test2
	json.Unmarshal([]byte(str), &t2) // 传指针是为了能在json.Unmarshal内部修改p2的值
	fmt.Println(t2)                  // {李四 18 male}
}
