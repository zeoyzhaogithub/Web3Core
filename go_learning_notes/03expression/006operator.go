package main

import (
	"fmt";
)

// 三、表达式
// 1. 保留字符
// 2. 运算符

// 自增 
// 自增、自减不再是运算符。只能作为独立语句，不能用于表达式

// 指针
// 不能将内存地址与指针混为一谈。
// 内存地址是内存中每个字节单元的唯一编号，而指针则是一个实体。
// 指针会分配内存空间，相当于一个专门用来保存地址的整型变量。

// 取址运算符“&”用于获取对象地址。
// 指针运算符“*”用于间接引用目标对象。
// 二级指针**T，如包含包名则写成*package.T。


func testPointer() {
    x := 10

    var p *int = &x      // 获取地址，保存到指针变量
	
    *p += 20             // 用指针间接引用，并更新对象

    fmt.Println(p, *p)       // 输出指针所存储的地址，以及目标对象
}

// 3. 初始化
// 对复合类型（数组、切片、字典、结构体）变量初始化时，
// 有一些语法限制。初始化表达式必须含类型标签。
// 左花括号必须在类型尾部，不能另起一行。
// 多个成员初始值以逗号分隔。
// 允许多行，但每行须以逗号或右花括号结束


func testStruct(){
	type data struct {
        x   int
        s   string
    }

    var a data = data{1, "abc"}

    b := data{
        1,
        "abc",
    }
	fmt.Println(a, b)
}

// 4. 流程控制
// if...else
// switch...case
// for
// 可用for...range完成数据迭代，
//   支持字符串、数组、数组指针、切片、字典、通道类型，
// 返回索引、键值数据。

// range会复制目标数据。
// 受直接影响的是数组，可改用数组指针或切片类型。
func testSlice(){
    data := [3]int{10, 20, 30}

    for i, x := range data {   // 从data复制品中取值
        if i == 0 {
            data[0] += 100
            data[1] += 200
            data[2] += 300
        }
        fmt.Printf("x: %d, data: %d\n", x, data[i])
    }
    println("----------------")
    for i, x := range data[:] {   // 仅复制slice，不包括底层array
        if i == 0 {
            data[0] += 100
            data[1] += 200
            data[2] += 300
        }

        fmt.Printf("x: %d, data: %d\n", x, data[i])
    }
}

/*
// 输出：
x: 10, data: 110
x: 20, data: 220             // range返回的依旧是复制值
x: 30, data: 330

x: 110, data: 210            // 当i == 0修改data时，x已经取值，所以是110
x: 420, data: 420            // 复制的仅是slice自身，底层array依旧是原对象
x: 630, data: 630
*/

// 相关数据类型中，字符串、切片基本结构是个很小的结构体，
// 而字典、通道本身是指针封装，复制成本都很小，无须专门优化。

// goto, continue, break
func testGoto() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto END
		}
		fmt.Println(i)
	}
    END:
	   fmt.Println("END")


    // start:                     // 错误: label start defined and not used
    for i := 0; i < 3; i++ {
        println(i)
        if i > 1 {
            goto exit
        }
    }
    exit:
       println("exit.")
}
func main() { 
	testPointer()
	testStruct()
	testSlice()
	testGoto()
}