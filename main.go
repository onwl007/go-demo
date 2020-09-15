package main

import (
	"fmt"
	"reflect"
)

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

	// 函数的反射
	f1 := fun1
	value := reflect.ValueOf(f1)
	fmt.Printf("Kind: %s, Type: %s\n", value.Kind(), value.Type())

	value2 := reflect.ValueOf(fun2)
	fmt.Printf("Kind: %s, Type: %s\n", value2.Kind(), value2.Type())

	// 通过反射调用函数
	value.Call(nil)

	value2.Call([]reflect.Value{reflect.ValueOf(100), reflect.ValueOf("hello")})

}

func fun1() {
	fmt.Println("我是函数 fun1()，无参的。。。")
}

func fun2(i int, s string) {
	fmt.Println("我是函数 fun2()，有参数的。。。", i, s)
}
