package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	1. 用于 goroutine，传递消息的
	2. 通道，每个都有相关联的数据类型
	3. 使用通道传递数据 <-
		chan <- data 发送数据到通道。向通道中写数据
		data <- chan 从通道中获取数据。从通道中读数据
	4. 阻塞
		发送和读取数据都是阻塞的
	5. 本身 channel 就是同步的，意味着同一时间，只能有一条 goroutine 来操作
	通道是 goroutine 之间的连接，所以通道的发送和接收必须处在不同的 goroutine 中
	*/
	var a chan int
	if a == nil {
		fmt.Println("channel 是 nil，不能使用，需要先创建通道。。")
		a = make(chan int)
		fmt.Printf("数据类型是：%T", a)
	}

	ch1 := make(chan int)
	fmt.Printf("%T,%p\n", ch1, ch1)

	test1(ch1)

	var ch2 chan bool
	fmt.Println(ch2)
	fmt.Printf("%T\n", ch2)
	ch2 = make(chan bool)
	fmt.Println(ch2)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子 goroutine 中，i: ", i)
		}
		ch2 <- true
		fmt.Println("结束。。。")
	}()

	data := <-ch2
	fmt.Println("data --> ", data)
	fmt.Println("main..over..")

	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)

	ch3 := make(chan int)
	go sendData(ch3)
	// for 循环的 for range 形式可用于从通道接收值，直到它关闭为止
	for v := range ch3 {
		fmt.Println("读取数据：", v)
	}
	fmt.Println("main...over...")
}

func test1(ch chan int) {
	fmt.Printf("%T,%p\n", ch, ch)
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func sendData(ch1 chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		ch1 <- i
	}
	close(ch1)
}
