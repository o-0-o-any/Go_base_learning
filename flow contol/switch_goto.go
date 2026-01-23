package main

import "fmt"

/*
switch-case基本格式
switch 标识符 {
case 数据1:
语句1
case 数据2:
语句2
case 数据3:
语句3
default:  // 相当于if判断语句中的else
语句
}
*/
func main() {
	// 例子 90-100分为A 80-89为B 70-79为C 60-69为D 小于60为E
	score := 79
	var ans string
	score /= 10
	switch score {
	case 10:
	case 9:
		ans = "A"
	case 8:
		ans = "B"
	case 7:
		ans = "C"
	case 6:
		ans = "D"
	default:
		ans = "E"

	}
	fmt.Println(ans)
	// 分支使用判断语句 switch后面不用加判断变量了
	score = 67
	switch {
	case score <= 100 && score >= 90:
		ans = "A"
	case score < 90 && score >= 80:
		ans = "B"
	case score < 80 && score >= 70:
		ans = "C"
	case score < 70 && score >= 60:
		ans = "D"
		fallthrough
	default:
		ans = "E"
	}
	fmt.Println(ans) // E fallthrough 执行满足case的下一个case 即ans="E"
	// goto
	label := true
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto l1
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
	label = false
l1:
	fmt.Print("我没有执行fmt.Println的换行 使用goto跳到了这里哦!label=false的赋值语句没有执行!label=", label)
	// 整体的结果:0 1 2 3 4 我没有执行fmt.Println的换行 使用goto跳到了这里哦!label=false的赋值语句没有执行!label=true  goto与l1之间的换行操作与赋值操作没有执行
}
