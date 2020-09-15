package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("hello, ", msg)
}

func (p Person) PrintInfo() {
	fmt.Printf("姓名：%s，年龄：%d，性别：%s\n", p.Name, p.Age, p.Sex)
}

func (p Person) Test(i, j int, s string) {
	fmt.Println(i, j, s)
}

func main() {
	/**
	反射
		指计算机程序在运行时可以访问、检测和修改它本身状态或行为的一种能力。
		反射就是程序在运行的时候能够观察并且修改自己的行为

	静态类型
		每个变量的类型在编译时都是确定的
	动态类型
		运行时给这个变量赋值时，这个值的类型（如果值为 nil 的时候没有动态类型）。
		一个变量的动态类型在运行时可能改变，这主要依赖于它的赋值（前提是这个变量是接口类型）

	其实反射的操作步骤非常的简单，就是通过实体对象获取反射对象 (Value, Type)，
	然后操作相应的方法即可

	反射 API 的分类总结
		1. 从实例到 Value
			通过实例获取 Value 对象
		2. 从实例到 Type
			通过实例获取反射对象的 Type
		3. 从 Type 到 Value
		4. 从 Value 到 Type
		5. 从 Value 到实例
		6. 从 Value 的指针到值
		7. Type 指针和值的相互转换
		8. Value 值的可修改性

	如何通过反射来进行方法的调用？
	本来可以用结构体对象，方法名称直接调用的
	但如果要通过反射，首先要将方法注册，也就是 MethodByName，然后通过反射调用 Call
	*/

	p2 := Person{"Ruby", 30, "男"}
	// 1. 要通过反射来调用起对应的方法，必须要先通过 reflect.ValueOf(interface)来获取到 reflect.Value,得到反射类型对象后才能做下一步处理
	getValue := reflect.ValueOf(p2)

	// 2. 一定要指定参数为正确的方法名，先看看没有参数的调用方法
	methodValue1 := getValue.MethodByName("PrintInfo")
	fmt.Printf("Kind: %s, Type: %s\n", methodValue1.Kind(), methodValue1.Type())
	methodValue1.Call(nil) // 没有参数，直接写 nil

	args1 := make([]reflect.Value, 0) // 或者创建一个空的切片也可以
	methodValue1.Call(args1)

	// 有参数的方法调用
	methodValue2 := getValue.MethodByName("Say")
	fmt.Printf("Kind: %s, Type: %s\n", methodValue2.Kind(), methodValue2.Type())
	args2 := []reflect.Value{reflect.ValueOf("反射机制")}
	methodValue2.Call(args2)

	methodValue3 := getValue.MethodByName("Test")
	fmt.Printf("Kind: %s, Type: %s\n", methodValue3.Kind(), methodValue3.Type())
	args3 := []reflect.Value{reflect.ValueOf(100), reflect.ValueOf(200), reflect.ValueOf("Hello")}
	methodValue3.Call(args3)

}
