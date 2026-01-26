package main

import "fmt"

/*
递归语法
func func() {
    func()  // 调用自己
}

func main() {
    func()
}
*/

// 求阶乘
func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

// 斐波那契数列
func Fibonacci_sequence(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fibonacci_sequence(n-1) + Fibonacci_sequence(n-2)
	}
}

// 上阶梯问题 一次可以走一阶或两阶 n阶台阶有几种走法
func step(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return step(n-1) + step(n-2)
	}
}

func main() {
	fmt.Print("你要求几的阶乘?")
	var n1 int
	fmt.Scan(&n1)                               // 5
	fmt.Printf("%d的阶乘为%d\n", n1, factorial(n1)) // 5的阶乘为120

	fmt.Print("你要得到斐波那契数列的几个数?")
	var n2 int
	fmt.Scan(&n2) // 10
	ans := make([]int, 0, n2)
	for i := 1; i <= n2; i++ {
		ans = append(ans, Fibonacci_sequence(i))
	}
	fmt.Printf("前%d个斐波那契数列中的数为:%v\n", n2, ans) // 前10个斐波那契数列中的数为:[1 1 2 3 5 8 13 21 34 55]

	fmt.Print("共有几阶台阶?")
	var n3 int
	fmt.Scan(&n3)                                    // 10
	fmt.Printf("一次过一阶或两阶 %d阶台阶有%d种方法", n3, step(n3)) // 一次过一阶或两阶 10阶台阶有89种方法
}
