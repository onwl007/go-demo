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

	p1 := Person{"王二狗", 30, "男"}
	DoFiledAndMethod(p1)

}

// 通过接口来获取任意参数
func DoFiledAndMethod(input interface{}) {
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is: ", getType.Name())
	fmt.Println("get Kind is: ", getType.Kind())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is: ", getValue)

	// 获取方法字段
	// 1. 先获取 interface 的 reflect.Type，然后通过 NumField 进行遍历
	// 2. 再通过 reflect.Type 的 Field 获取其 Field
	// 3. 最后通过 Field 的 Interface() 得到对应的 value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("字段名称：%s，字段类型：%s，字段数值：%v \n", field.Name, field.Type, value)
	}

	// 通过反射，操作方法
	// 1. 先获取 interface 的 reflect.Type，然后通过 NumMethod 进行遍历
	// 2. 再通过 reflect.Type 的 Method 获取其 Method
	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法名称：%s，方法类型：%v \n", method.Name, method.Type)
	}
}
