package main

import "fmt"

// Go 的 for 语句后面的三个构成部分外没有小括号， 大括号 { } 则是必须的。
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
