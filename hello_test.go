package main

import "testing"

/**
1. 程序需要在一个名为 xxx_test.go 的文件中编写
2. 测试函数的命名必须以单词 Test 开始
3. 测试函数只接受一个参数 t *testing.T
*/
func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
