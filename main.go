package main

import "fmt"

func main() {
	/**
	Go 中字符串是一个字节的切片，是 Unicode 兼容的，并且是 UTF-8 编码的
	*/
	name := "Hello World"
	fmt.Println(name)

	for i := 0; i < len(name); i++ {
		fmt.Printf("%d ", name[i])
	}
	fmt.Println()
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c ", name[i])
	}
}
