package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct{}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct{}

func (iphone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

type Student struct{}

func main() {
	/**
	接口
		Go 中，接口是一组方法签名。当类型为接口中的所有方法提供定义时，被称为实现接口
		接口定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了该接口

		1：interface 可以被任意的对象实现
		2：一个对象可以实现任意多个 interface
		3：任意的类型都实现了空 interface，也就是包含 0 个 method 的 interface
	*/
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	// 接口断言
	var i1 interface{} = new(Student)
	s := i1.(Student) // 不安全，如果断言失败，则会直接 panic

	fmt.Println(s)

	var i2 interface{} = new(Student)
	s, ok := i2.(Student) // 安全，断言失败，也不会 panic，只是 ok 的值为 false
	if ok {
		fmt.Println(s)
	}
}
