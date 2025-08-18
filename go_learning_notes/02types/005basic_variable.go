package main
import (
	"fmt"; "math"; "strconv";
)
// 2、数据类型
/*
在Go语言中，数据类型分为以下两种
基本数据类型（原生数据类型）​：
    整型、分两大类
	  有符号整型：int8、int16、int32、int64、int
	  无符号整型：uint8、uint16、uint32、uint64、uint
	浮点型、
	复数型、
	布尔型、
	字符串、
	字符（byte、rune）​
复合数据类型（派生数据类型）​：
    数组（array）​、切片（slice）​、映射（map）​、
	函数（function）​、结构体（struct）​、通道（channel）​、
	接口（interface）​、指针（pointer）​
*/

func testMath(){
	a,b,c := 100,0144,0x64
	fmt.Println(a,b,c)
	fmt.Printf("0b%b, %#o, %#x\n",a,a,a)
	fmt.Println(math.MinInt8,math.MaxInt8)
}

// 标准库strconv可在不同进制（字符串）间转换。
func testStrconv(){
	aa,err := strconv.Atoi("100")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(aa)
	
    a, _ := strconv.ParseInt("1100100", 2, 32)
    b, _ := strconv.ParseInt("0144", 8, 32)
    c, _ := strconv.ParseInt("64", 16, 32)

    println(a, b, c)

    println("0b" + strconv.FormatInt(a, 2))
    println("0" + strconv.FormatInt(a, 8))
    println("0x" + strconv.FormatInt(a, 16))

}

// 5. 引用类型 
// 所谓引用类型(reference type)特指slice、map、channel这三种预定义类型。
// 相比数字、数组等类型，引用类型拥有更复杂的存储结构。
// 除分配内存外，它们还须初始化一系列属性，诸如指针、长度，甚至包括哈希分布、数据队列等。
// 
// 内置函数new按指定类型长度分配零值内存，返回指针，并不关心类型内部构造和初始化方式。
// 而引用类型则必须使用make函数创建，编译器会将make转换为目标类型专用的创建函数（或指令）​，
// 以确保完成全部内存分配和相关属性初始化。


func mkslice() []int {
    s := make([]int, 0, 10)
    s = append(s, 100)
    return s
}

func mkmap() map[string]int {
    m := make(map[string]int)
    m["a"] = 1
    return m
}

// 6. 类型转换

// 7. 自定义类型
// 使用关键字type定义用户自定义类型，
// 包括基于现有基础类型创建，或者是结构体、函数类型等。
// 多个type定义可合并成组，可在函数或代码块内定义局部类型。

func testType() {
    type (                              // 组
        user struct {                   // 结构体
            name string
            age  uint8
        }

        event func(string) bool         // 函数类型
    )

    u := user{"Tom", 20}
    fmt.Println(u)

    var f event = func(s string) bool {
        println(s)
        return s != ""
    }

    mm := f("abc")
	println(mm)
}

// 8. 未命名类型 
// 具有相同声明的未命名类型被视作同一类型。
// 具有相同基类型的指针。
// 具有相同元素类型和长度的数组(array)。
// 具有相同元素类型的切片(slice)。
// 具有相同键值类型的字典(map)。
// 具有相同数据类型及操作方向的通道(channel)。
// 具有相同字段序列（字段名、字段类型、标签，以及字段顺序）的结构体(struct)。
// 具有相同签名（参数和返回值列表，不包括参数名）的函数(func)。
// 具有相同方法集（方法名、方法签名，不包括顺序）的接口(interface)。
// 相关类型会在后续章节做详细说明，此处无须了解更多细节。
// 容易被忽视的是struct tag，它也属于类型组成部分，而不仅仅是元数据描述。

// 未命名类型转换规则：
// 所属类型相同。
// 基础类型相同，且其中一个是未命名类型。
// 数据类型相同，将双向通道赋值给单向通道，且其中一个为未命名类型。
// 将默认值nil赋值给切片、字典、通道、指针、函数或接口。
// 对象实现了目标接口。
func testType2() {
    type data [2]int
    var d data = [2]int{1, 2}    // 基础类型相同，右值为未命名类型

    fmt.Println(d)

    a := make(chan int, 2)
    var b chan<- int = a  // 双向通道转换为单向通道，其中b为未命名类型

    b <- 2
	fmt.Println(<-a)
	fmt.Println(a)
}
func main() { 
	testMath()
	testStrconv()

	m := mkmap()
    println(m["a"])
    println(m)

    s := mkslice()
    println(s[0])

	testType()
	testType2()

}