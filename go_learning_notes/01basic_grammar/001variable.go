package main
import "fmt"
// 基础语法

// go语言定义的变量或者import包如果没有用到，编译就不会通过

// 1、匿名变量
func getDate() (int, int){
	return 2025,8
}

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

// 3、表达式
// go语言有三种流程控制语句
// if
func testIf(){
	var x int = 100
	if x > 0 {
		fmt.Println("x >0 ")
	} else if x < 0 {
		fmt.Println("x < 0")
	}else {
		fmt.Println("x = 0")
	}
}

// switch
func testSwitch(){
	var x int = 100
	switch x {
	case 0:
		fmt.Println("x = 0")
	case 1:
		fmt.Println("x = 1")
	default:
		fmt.Println("x > 1")
	}
}

// for
func testFor(){
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
}
func testFor2(){
	i := 0
	// 相当于while x<3
	for i < 3 {
		fmt.Println(i)
		i++
	}
}

func testFor3(){
	x := 3
	// 相当于while true
	for {
		fmt.Println("hello world")
		x--
		if x == 0 {
			break
		}
	}
}

// 迭代遍历，for...range除元素外，还可返回索引
func testFor4(){
	x := [] int{1,2,3}
	for index, value := range x {
		fmt.Println(index,":", value)
	}
}



func main(){
   a, _ := getDate()   // 舍弃第二个返回值
   _, b := getDate()   // 舍弃第一个返回值
   // 如果不使用变量，则运行代码时会输出declared and not used: a
   fmt.Println("a=",a,"b=",b)
   fmt.Println("b----------------")
   testIf()
   testSwitch()
   testFor()
   testFor2()
   testFor3()
   fmt.Println("b----------------")
   testFor4()
}
