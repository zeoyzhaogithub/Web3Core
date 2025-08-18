package main

import (
	"fmt";"unsafe"
)

// 二、类型
// 1、变量

// 简短模式(short variable declaration)有些限制：
// 定义变量，同时显式初始化。
// 不能提供数据类型。
// 只能用在函数内部。


// 2. 命名
// 区分大小写。
// 使用驼峰(camel case)拼写格式。
// 局部变量优先使用短名。
// 不要使用保留关键字。
// 不建议使用与预定义常量、类型、内置函数相同的名字。
// 专有名词通常会全部大写，例如escapeHTML。
// 尽管Go支持用汉字等Unicode字符命名，但从编程习惯上来说，这并不是好选择。

// 3. 常量
// 可在函数代码块中定义常量，
// 不曾使用的常量不会引发编译错误。
func testConst(){
    const x,y = 1,2
    const (
	   i,f = 1,2
	   z = 3
    )
	fmt.Println(x,y,i,f,z)

	// 如果显式指定类型，必须确保常量左右值类型一致，需要时可做显式转换。
	// 右值不能超出常量类型取值范围，否则会引发溢出错误。
	const (
        h, l int    = 99, -999
        b    byte   = byte(h)         // x被指定为int类型，须显式转换为byte类型
        //n           = uint8(l)        // 错误: constant -999 overflows uint8
    )

	// 常量值也可以是某些编译器能计算出结果的表达式，
	// 如unsafe.Sizeof、len、cap等。
	const (
		ptrSize = unsafe.Sizeof(uintptr(0))
		strSize = len("hello")
	)
	fmt.Println(ptrSize, strSize)

	// 在常量组中如不指定类型和初始化值，则与上一行非空常量右值（表达式文本）相同。
	const (
		d uint16 = 120
		e           // 与上一行的d类型、右值相同
		s = "hello"
		t
	)
	fmt.Printf("%T, %v\n",d, d)
	fmt.Printf("%T, %v\n",t, t)
}

// 枚举
// Go并没有明确意义上的enum定义，
// 不过可借助iota标识符实现一组自增常量值来实现枚举类型。

func testEnum(){
	const (
		a = iota // iota初始值为0
		b
		c
	)
	fmt.Println(a, b, c)
	const (
		_ = iota
		KB = 1 << (10 * iota)  // 1 << (10 * 1)
		MB                     // 1 << (10 * 2)
		
		GB                     // 1 << (10 * 3)
	)
	fmt.Println(KB, MB, GB)

	// 自增作用范围为常量组。
	// 可在多常量定义中使用多个iota，它们各自单独计数，只须确保组中每行常量的列数量相同即可。
	const (
        _, _ = iota, iota * 10    // 0，0 * 10
        d, e                      // 1, 1 * 10
        f, g                      // 2, 2 * 10
    )
	fmt.Println(d, e, f, g)

	// 如中断iota自增，则必须显式恢复。
	// 且后续自增值按行序递增，而非C enum那般按上一取值递增。
	const (
        h   = iota            // 0
        l                     // 1
        j   = 100             // 100
        k                     // 100（与上一行常量右值表达式相同）
        n   = iota            // 4（恢复iota自增，计数包括j、k）
        m                     // 5
    )
	fmt.Println(h, l, j, k, n, m)
}

// 常量除“只读”外，和变量究竟有什么不同？
// 1. 不同于变量在运行期分配存储内存（非优化状态）​，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用。
// 2. 数字常量不会分配存储空间，无须像变量那样通过内存寻址来取值，因此无法获取地址。

func main() {
    testConst()
	testEnum()
}