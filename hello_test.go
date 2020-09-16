package main

import "testing"

/**
1. 程序需要在一个名为 xxx_test.go 的文件中编写
2. 测试函数的命名必须以单词 Test 开始
3. 测试函数只接受一个参数 t *testing.T
*/
func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		// 告诉测试套件这个方法时辅助函数，当测试失败时所报告的行号酱紫啊函数调用中而不是在辅助函数内部
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("to a person", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Lauren", "French")
		want := "Bonjour, Lauren"
		assertCorrectMessage(t, got, want)
	})
}
