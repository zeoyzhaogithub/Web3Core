package main

import (
	"fmt"
	"errors"
	"log"
	"runtime/debug" // 调试阶段，输出完整调用堆栈信息
)
// 4.4 匿名函数

// 匿名函数可以在函数内部定义，可以直接调用，保存到变量
// 普通函数和匿名函数都可以作为结构体字段，或经过通道传递

func testStruct(){
	type callc struct {
		mul func(int, int) int    // 匿名函数
	}

	c := callc{
		mul: func(a, b int) int {
			return a * b
		},
	}
	fmt.Println(c.mul(2, 3))
}

// 不曾使用的匿名函数会被编译器当做错误
func testChannel(){
	ch := make(chan func(int, int) int,2)
	
	ch <- func(a, b int) int {
		return a + b
	}

	fmt.Println((<-ch)(1, 2))
}

// 4.4.2 闭包closure 
// 在其词法上下文中引用了自由变量的函数，或者说是函数和其引用的环境的组合体
// 闭包是函数和引用环境的组合体


// 4.5 defer延迟调用 
// 语句defer向当前函数注册稍后执行的函数调用 这些调用被称作延迟调用，
// 因为他们直到当前函数执行结束前才被执行，常用于资源释放、解除锁定，以及错误处理

func testDefer(){ 
	x, y := 1, 2
	defer func(a int){
		fmt.Println("defer x, y = ",a, y)  // y为闭包引用
	}(x) // 注册时复制调用参数
	x += 100    // 对x的修改不会影响延迟调用
	y += 1000
	fmt.Println("x, y = ",x, y)
}

// 4.6 错误处理
var errDivByZero = errors.New("除数不能为0")

// 4.6.2 panic,recover
// panic会立即中断当前函数的流程执行延迟调用,
// 在执行延迟调用的函数中recover可捕获并返回panic提交的错误对象

func testPanic(){
	defer func(){
		if err := recover(); err != nil { // 捕获错误
			log.Fatalln(err)
			// debug.PrintStack()
		}
	}()
	panic("I am panic")  // 抛出错误
	log.Println("I am log")
}

// 连续调用panic，仅最后一个会被recover捕获

func main() {
	// 匿名函数
	// var f func(int, int) int
	// f = func(a, b int) int {
	// 	return a + b
	// }
	// fmt.Println(f(1, 2))
	s := func(s string){
		fmt.Println(s)
	}
	s("abc")
	testChannel()

	testDefer()
	testPanic()
}