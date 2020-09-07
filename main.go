package main

import (
	"fmt"
	"math"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (e Employee) displaySalary() {
	fmt.Printf("Slary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func main() {
	/**
	方法：
		就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或是指针
	语法
		func (t Type) methodName(parma list)(return list) {}
	*/
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary()

	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())
}
