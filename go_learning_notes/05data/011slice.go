package main

import (
	"fmt"
	"unsafe"
	"reflect"
	"errors"
)

/*
5.3 切片slice
切片并非动态数组或数组指针。它内部通过指针引用底层数组，
设定相关属性将数据读写操作限定在指定区域内
切片是只读对象
*/

func slice_test(){
    x := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(x[1:4])
	m := make([]int,5,10)  // 创建切片 长度为5，容量为10
	fmt.Println(m)
	fmt.Println(len(m), cap(m))
	s2 := make([]int, 5)    // 创建切片 长度为5 省略cap，和len相同
	fmt.Println(s2, len(s2), cap(s2))
	s3 := make([]int, 10, 20)
	fmt.Println("s3---",s3, len(s3), cap(s3))

	var a []int // 定义一个[]int类型变量，并未执行初始化操作
	b := []int{}  // 用初始化表达式完成了全部创建过程。
    //fmt.Printf("a---",a==nil, len(a), b == nil)

	// a:&reflect.SliceHeader{Data:0x0, Len:0, Cap:0}
	fmt.Printf("a:%#v\n",(*reflect.SliceHeader)(unsafe.Pointer(&a)))
	// b:&reflect.SliceHeader{Data:0x103caec0, Len:0, Cap:0}
	fmt.Printf("b:%#v\n",(*reflect.SliceHeader)(unsafe.Pointer(&b)))
    
	// 变量b的指针被赋值，尽管它指向runtime.zerobase，但它依然完成了初始化操作 
	// 另外， a==nil，仅表示它是个未初始化的切片对象，切片本身依然会分配所需内存。
	// 可以直接对nil切片执行slice[:]操作，同样返回nil

	fmt.Printf("a size:%d\n", unsafe.Sizeof(a))
	fmt.Printf("b size:%d\n", unsafe.Sizeof(b))
}

// 可以获取元素地址，但是不能像数组那样直接用指针访问元素内容
func slice_test2() { 
    s := []int{1,2,3,4,5,6,7,8,9,10}
	p := &s  // 取header地址
	p0 := &s[0]  // 取元素0的地址
	p1 := &s[1]
	fmt.Printf("%T, %v\n", p, p)
	fmt.Println(p,p0,p1)

	(*p)[0] += 100 // *[]int不支持索引操作，必须先返回[]int对象
	*p1 += 100   // 直接用元素指针操作
	fmt.Println(s)
}

// 如果元素类型也是切片，那么就可实现类似交错数组（jagged array）功能
func slice_test3() { 
    s := [][]int{
		[]int{1,2,3},
		[]int{4,5,6},
		[]int{7,8,9},
	}
	fmt.Println(s)
	for i := range s {
		for j := range s[i] {
			s[i][j] *= 10
		}
	}
	fmt.Println(s)
	s[2] = append(s[2], 10,300)
	fmt.Println(s[2])
	fmt.Println(s)
}

// make函数允许在运行期动态指定数组长度，绕开了数组类型必须使用编译期常量的限制 
// 并非所有时候都适合用切片代替数组，小数组在栈上拷贝的消耗并不比make代价大

// 5.3.2 reslice 
// 将切片视作[cap]slice数据源，据此创建新的切片对象。不能超出cap，但不受len限制 
// 新建切片对象依旧指向原底层数组，既修改对所有关联切片可见
func slice_test4() { 
	d := []int{1,2,3,4,5,6,7,8,9,10}
	s1 := d[2:7]
	s2 := s1[2:5]
	for i := range s2 {
		s2[i] += 100
	}
	fmt.Println(d)
	fmt.Println(s1)
	fmt.Println(s2)
}

// 利用reslice操作，可实现一个栈式数据结构
func reslice_stack() { 
	// 栈最大荣浏览量5
	stack := make([]int, 0, 5)

	// 入栈
	push := func(x int) error{
		n := len(stack)
		if n == cap(stack) {
			return errors.New("stack is full")
		}
		stack = stack[:n+1]
		stack[n] = x
		return nil
	}
	// 出栈
	pop := func() (int, error) {
		n := len(stack)
		if n == 0 {
			return 0, errors.New("stack is empty")
		}
		x := stack[n-1]
		stack = stack[:n-1]
		return x, nil
	}
	// 入栈元素
	for i := 0; i < 7; i++ {
		fmt.Printf("push %d:%v, %v\n", i, push(i), stack)
	}
	// 出栈元素
	for i := 0; i < 7; i++ {
		x, err := pop()
		fmt.Printf("pop %d:%v, %v, %v\n", i, x, err, stack)
	}
}

// 5.3.3 append 向切片尾部添加数剧，返回新的切片对象
func append_test(){
    s := make([]int, 0, 5)
    s1 := append(s, 10)
	s2 := append(s1, 20, 30)

	fmt.Println(s, len(s), cap(s))   // 不修改原slice属性
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
}

// 数据被追加到原底层数组。如果超出cap限制，则为新切片对象重新分配数组
func append_test2(){
    s := make([]int, 0, 100) // 创建长度为0，容量为100的整型切片
	// s1 与 s 共享同一个底层数组
    s1 := s[:2:4] // 从s创建新切片，长度=2，容量=4
	s2 := append(s1, 1,2,3,4,20, 30)  // 超出s1 cap限制，分配新底层数组
	fmt.Printf("s1:%p:%v\n", &s1[0], s1)
	fmt.Printf("s2:%p:%v\n", &s2[0], s2)
	fmt.Printf("s data:%v\n", s[:10])
	fmt.Printf("s1 cap:%d, s2 cap:%d\n", cap(s1), cap(s2))

	// 初始状态
// 	底层数组 (100个int容量)
// ┌───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬─── ... ───┐
// │   │   │   │   │   │   │   │   │   │   │            │
// └───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴─── ... ───┘
// ↑
// s: len=0, cap=100

// s1 := s[:2:4] 后：
// s1指向同一个数组，但视野受限
// ┌───┬───┬───┬───┼───┼───┼─── ... ───┐
// │ s1[0] │ s1[1] │   │   │   │       │  ← s1看到的范围
// └───┴───┴───┴───┼───┼───┼─── ... ───┘
// s1: len=2, cap=4

// 追加操作后：
// 由于超出 s1 的容量，Go 分配新的底层数组：
// 新底层数组 (更大容量)
// ┌───┬───┬───┬───┬───┬───┬───┬───┬───┬───┐
// │s1[0]│s1[1]│ 1 │ 2 │ 3 │ 4 │20 │30 │   │   │ ... 
// └───┴───┴───┴───┴───┴───┴───┴───┴───┴───┘
// ↑
// s2: len=8, cap=?

// 原底层数组 (保持不变)
// ┌───┬───┬───┬───┬───┬───┬─── ... ───┐
// │   │   │   │   │   │   │            │
// └───┴───┴───┴───┴───┴───┴─── ... ───┘
// ↑
// s和s1仍然指向这里

/*
扩容机制详解
Go 的切片扩容算法：
计算所需容量：newLen = oldLen + numElements

如果所需容量 ≤ 2×旧容量：新容量 = 2×旧容量

如果所需容量 > 2×旧容量：新容量 = 所需容量

内存对齐：根据元素大小进行内存对齐调整

对于此例：
旧容量 = 4
所需容量 = 2 + 6 = 8
2×旧容量 = 8
所以新容量 = 8
*/
}

// 5.3.4 copy
// 在两个切片对象间复制数据，允许指向同一底层数组，允许目标区间重叠。
// 最终所复制长度以较短的切片长度为准

func copy_test() { 
	s1 := []int{1,2,3,4,5,6,7,8,9,10}
	s2 := s1[5:8]
	fmt.Println(s2)
	fmt.Println(s1[4:])
	n := copy(s2, s1[4:])  // 在同一底层数组间的不同区间复制，
	fmt.Println(s2, n)
	/*
	初始状态：
s1: [1,2,3,4,5,6,7,8,9,10]
s2: s1[5:8] = [6,7,8] (长度=3)

copy(s2, s1[4:]) 执行：
src: s1[4:] = [5,6,7,8,9,10] (长度=6)
dst: s2 = [6,7,8] (长度=3)

复制 min(3,6)=3 个元素：
将 src[0:3] = [5,6,7] 复制到 dst[0:3]

结果：
s2 变为 [5,6,7]
s1 也相应改变（因为共享底层数组）：
s1: [1,2,3,4,5,5,6,7,9,10]
	*/

	// 在不同数组间复制
	s3 := make([]int, 5)
	n2 := copy(s3, s1)
	fmt.Println(s3, n2)

// func copy(dst, src []T) int
// 作用：将 src 切片的数据复制到 dst 切片
// 返回值：实际复制的元素个数
// 规则：复制 min(len(dst), len(src)) 个元素

// 应用 1. 缓冲区管理 2. 数据截断处理 3. 滑动窗口协议 4. 内存池模式
}

func main() { 
	slice_test()
	slice_test2()
	slice_test3()
	slice_test4()
	fmt.Println("------88888------------")
	reslice_stack()
	append_test()
	fmt.Println("------88888----999--------")
	append_test2()
	copy_test()
}