package main

import (
	"fmt"
)
/*
5.2 数组
定义数组类型时，数组长度必须是非负整型常量表达式，长度是类型组成部分。
既元素类型相同，但长度不同的数组不属于同一类型
*/

func test_array(){
	var d1 [3]int    // 数组类型为[3]int,元素自动初始化为0
	var d2 [2]int
	fmt.Println(d1, d2)
	// if d1 == d2 {
	// 	fmt.Println("d1 == d2")
	// } else {
	// 	fmt.Println("d1 != d2")
	// }
	b := [4]int{2,5}// [2 5 0 0] ,未提供初始化值的元素自动初始化为0
	c := [...]int{1,2,3,4}  // [...]int表示长度未知的数组，编译器会自动计算长度
	d := [5]int{5,3:10}   // 可指定所以初始位置
	e := [...]int{10,3:100}  // 支持索引初始化，但注意数组长度与此有关
	fmt.Println(b,c,d,e)
}

// 对于结构等复合类型，可省略元素初始化类型标签
func test_array2() {
   type user struct {
       name string
       age int
   }

    d := [...]user{  // 可省略元素初始化类型标签
       {"Tom", 20},
       {"Jerry", 30},
    }
    fmt.Printf("%#v\n", d)

	// 定义多维数组时，仅第一位度允许使用...
	e := [2][3]int{
		{1,2,3},
		{4,5,6},
	}
	f := [...][2]int{
		{1,2},
		{3,4},
		{5,6},
	}
	fmt.Println(e,f)
	//内置函数len和cap都返回第一维数组的长度
	fmt.Println(len(e), cap(e))
	fmt.Println(len(f), cap(f))
}

// 5.2.2 指针
// 要分清楚指针数组和数组指针的区别。
// 指针数组是指元素为指针的数组，
// 数组指针是指指向数组的指针，获取数组变量的地址

func test_array3() { 
	x, y := 10,20   
	a := [...]*int{&x, &y}    // 元素为指针的指针数组
	p := &a   // 存储数组地址的数组指针
	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", p, p)

	// 数组指针可直接用来来操作元素
}

// 5.2.3 复制
// go数组是值类型，赋值和传参操作都会复制整个数组数据
func test(x [2]int){
	fmt.Printf("x: %p, %v\n", &x, x)
}

// 可以代用指针或切片，以此避免你数据复制
func test2(x *[2]int){
	fmt.Printf("x: %p, %v\n", x, *x)
	x[1] += 100
}

func test_copy() { 
	a := [2]int{1,2}
	var b [2]int
	b = a
	fmt.Printf("a: %p, %v\n", &a, a)
	fmt.Printf("b: %p, %v\n", &b, b)
	test(a)
	test2(&a)
	fmt.Printf("a: %p, %v\n", &a, a)
}


func main() { 
    test_array()
	test_array2()
	test_array3()
	test_copy()
}