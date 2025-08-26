package main

import (
	"fmt"
	"reflect"
	"unsafe"
	"unicode/utf8"
)

// 5.1 字符串 for遍历字符串，分byte和runne
func string_test(){
	s := "hello world"
	s1 := s[:3]
	s2 := s[1:4]
	s3 := s[2:]

	fmt.Println(s1, s2, s3)
	// 提示：reflect.SliceHeader和string头结构相同
	// unsafe.Pointer用于指针类型转换
	fmt.Printf("%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&s)))
	fmt.Printf("%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&s1)))
}
func string_for() {
	s := "西湖晴雨"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d:[%c]\n", i,s[i])
	}
	for i, c := range s{
		fmt.Printf("%d:[%c]\n", i,c)
	}
}
// 5.1.2 转换 
// 要修改字符串，需要先将其转换为可变类型([] rune,[]byte),
// 待完成后再转换回来。并且必须重新分配内存

func pp(format string, ptr interface{}) {
    p := reflect.ValueOf(ptr).Pointer()
    h := (*uintptr)(unsafe.Pointer(p))
    fmt.Printf(format, *h)
}

func reflect_test() {
    s := "hello, world!"
    pp("s: %x\n", &s)

    bs := []byte(s)
    s2 := string(bs)

    pp("string to []byte, bs: %x\n", &bs)
    pp("[]byte to string, s2: %x\n", &s2)

    rs := []rune(s)
    s3 := string(rs)

    pp("string to []rune, rs: %x\n", &rs)
    pp("[]rune to string, s3: %x\n", &s3)
}

// 5.1.3 unicode
/*
类型rune专门用来存储unicode码点（code poinnt）,它是int32的别名。
使用单引号的字面量，其默认类型就是rune
*/

func unicode_test() {
	var r rune = '中'
	pp("r: %x\n", &r)
	s := '中'
	fmt.Printf("s: %T\n", s)
	// 除[]rune外，还可直接在rune、byte、string之间转换
	s1 := string(r)    // rune -> string
	b := byte(r)    // string -> byte
	s2 := string(b)  // byte -> string
	r2 := rune(b)   // byte -> rune
	fmt.Println(r,"--",s1, b, s2, r2)
	// 字符串存储的字节数组，不一定就是合法的utf-8文本
	ss := "中.国"
	aa := string(ss[0:1]+ss[3:4]) // 截取并拼接一个”不合法“的字符串
	fmt.Println(aa, utf8.ValidString(aa))
	// RuneCountInString代替len返回准确的unicode字符数量
	fmt.Println("字符长度",len(ss),utf8.RuneCountInString(ss))
}

func main() {
	string_test()
	// string_for()
	reflect_test()
	unicode_test()
}