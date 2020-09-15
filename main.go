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
	*/

	// 反射操作：通过反射，可以获取一个接口类型变量的类型和数值
	var x float64 = 3.4
	fmt.Println("type: ", reflect.TypeOf(x))
	fmt.Println("value: ", reflect.ValueOf(x))

	fmt.Println("---------------")

	// 根据反射的值，来获取对应的类型和数值
	v := reflect.ValueOf(x)
	fmt.Println("kind is float64: ", v.Kind() == reflect.Float64)
	fmt.Println("type: ", v.Type())
	fmt.Println("value: ", v.Float())

	var num float64 = 1.2345
	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)
	// 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接 panic
	// Golang 对类型要求非常严格，类型一定要完全符合
	// 如下两个，一个是 *float64，一个是 float64，如果弄混，则会 panic
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)
	fmt.Println(convertPointer)
	fmt.Println(convertValue)
}
