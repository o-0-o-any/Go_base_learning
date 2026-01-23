package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// 字符串声明 用双引号包裹
	var s1 string
	s1 = "Hello World"
	// var s1 string = "Hello World"  // 声明与赋值写在一起
	// 短声明字符串
	s2 := "Hello World!"
	// 声明字符变量 单独的字母 数字符号 类型是int32类型
	c1 := '我'
	var c2 int32 = 'h'
	// 打印数据类型和数据
	fmt.Printf("%T %T %T %T\n", s1, s2, c1, c2) // string string int32 int 32
	fmt.Printf("%s %s %c %c\n", s1, s2, c1, c2)
	fmt.Println(c1, c2) // 这种打印出的内容是Unicode编码值 25105 104
	// 字节 1字节=8bit即8个二进制位  1个字符如'A'=1个字节  应该utf-8编码的汉子'我'=3个字节

	// 多行字符串
	s3 := `
第一行
第二行
第三行
	`
	// 打印多行字符串
	fmt.Println(s3)
	// 字符串相关操作
	// 打印字符串长度
	fmt.Println(len(s1), len(s2), len(s3))
	// 拼接字符串
	s4 := "hello"
	s5 := "world"
	s6 := fmt.Sprintln(s4, s5) // 或s6 := fmt.Sprintf("%s %s", s4, s5)
	// 等同于 s6 = s4 + " " +s5
	fmt.Println(s6)
	// 分割字符串  得到n个对象存在一个切片中 strings包中的Split方法 传入两个参数 第一个是待分割的对象 第二个是分割哪个 两个参数都是字符串类型 与Python中的Split方法类似
	s7 := strings.Split(s3, "\n")
	fmt.Println(s7) // [ 第一行 第二行 第三行  ]
	// 包含  传入两个参数 第一个是判断的对象 第二个是判断第一个是否含有的字符串 结果是bool值
	fmt.Println(strings.Contains(s3, "第二行")) // true
	fmt.Println(strings.Contains(s3, "第二列")) // false
	// 前缀后缀判断  判断字符串是否由指定字符串为前缀或后缀 也是包含两个参数 第一个参数是判读字符串对象 第二个参数是判断的字符串内容 HasPrefix与HasSuffix分别是判断前缀与后缀
	fmt.Println(strings.HasPrefix(s1, "Hello")) // true
	fmt.Println(strings.HasSuffix(s1, "World")) // true
	// 子串出现的位置 结果是子串出现的位置下标 从0开始计数 也是包含两个参数第一个参数是主字符串 第二个参数是子串 Index与LastIndex分别是得到子串首次在字符串中出现的位置与子串最后一次在字符串出现的位置
	fmt.Println(strings.Index(s2, "l"), strings.LastIndex(s2, "l")) // 2 9
	// 拼接 也是包含两个参数 第一个是切片(将字符串分割之后得到切片 其中包含n个字符串) 第二个是拼接内容 如用+拼接切片中的字符串
	fmt.Println(strings.Join(s7, "+")) // +第一行+第二行+第三行+
	// 遍历字符串得到单个字符
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%c ", s1[i])
	}
	fmt.Println()
	for _, c := range s1 { // 使用range类似Python中的enumerate生成可迭代对象 这里使用匿名变量_省略掉序号0...
		fmt.Printf("%c ", c)
	}
	fmt.Println()
	// 修改字符串  先将字符串转换成byte或rune类型
	s8 := "big" // 非法 cannot assign to s8[0] (neither addressable nor a map index expression)
	// s8[0] = 'p'
	bytes8 := []byte(s8)        // 强制类型转换
	bytes8[0] = 'p'             // 修改单个字符 就是字符类型 用单引号包裹
	fmt.Println(string(bytes8)) // pig
	s9 := "物理"
	runes9 := []rune(s9)
	runes9[0] = '地'
	fmt.Println(string(runes9))   // 地理
	t1 := "红"                     // string
	t2 := '红'                     // int32
	fmt.Printf("%T %T\n", t1, t2) // string int32
	// 类型转换  使用math.Sqrt()函数计算直角三角形斜边长 这个函数接受的参数类型是float64 需要强制类型转换
	a, b := 3, 4
	var c int = 0
	c = int(math.Sqrt(float64(a*a + b*b))) // 先转换成float64类型传入math.Sqrt()函数 然后将计算结果转换成int类型赋值给c
	fmt.Printf("%T %d\n", c, c)            // int 5
}
