package main
import (
    "fmt"
	"time"
	"log"
	"errors"
)

// 第一类对象(first-class object)指可在运行期间创建，可用作函数参数或返回值，可存入变量的实体。
//最常见的用法就是匿名函数
func hello(){
	fmt.Println("hello world")
}

func exex(f func()){
	f()
}

// 指针
func test() *int{  // 声明返回 int 指针的函数
	a := 0x100     // 在栈上创建局部变量 a，值为 256 (0x100)
	return &a      // 返回变量 a 的内存地址
}

/*
关键机制：逃逸分析（Escape Analysis）
为什么这段代码能正常工作？
在传统语言（如 C/C++）中，返回局部变量地址会导致悬垂指针，但在 Go 中：

编译器检测：Go 编译器进行逃逸分析

内存逃逸：发现局部变量 a 的地址被返回，需要在函数结束后继续存在

堆分配：自动将 a 分配到堆（heap）而非栈（stack）上

内存管理：由 Go 的垃圾回收器（GC）负责回收


内存分配过程
test() 执行前：          test() 执行后：
┌─────────────┐         ┌─────────────┐
│   栈空间     │         │   栈空间     │
│             │         │             │
│             │         │             │
└─────────────┘         └─────────────┘
                         ↑
                         │
                         └───┐
                             ↓
                         ┌─────────────┐
                         │   堆空间     │
                         │   a = 256   │
                         └─────────────┘

使用场景
1、工厂模式创建对象
2、构建复杂的数据结构
3、避免大对象赋值
4、状态共享

实际使用示例
配置对象创建
构建器模式

性能考虑
优点：
避免复制：大对象传递指针更高效
状态共享：多个部分可访问同一对象
内存安全：由GC自动管理，无手动内存管理负担

缺点：
堆分配开销：比栈分配稍慢
GC压力：增加垃圾回收器的工作量
缓存不友好：指针跳转可能影响CPU缓存效率

使用场景：

创建需要共享或修改的对象
避免大对象的复制开销
实现工厂模式和构建器模式
需要跨函数边界共享状态时
*/

// 4.2 参数
// 不支持有默认值的可选参数，不支持命名实参
// 调用时，必须按签名顺序传递指定类型和数量的参数
// 在参数列表中，相邻的同类型参数可合并

func test1(p *int) {
	// fmt.Println("----",*p)
    go func() {      // 延长p的生命周期,启动新协程
        fmt.Println("----2",*p)
    }()
}
// 使用场景
// 异步任务处理（日志、清理、通知）
// 并发计算（并行处理多个任务）
// 事件监听（信号处理、定时任务）
// 资源释放（在后台安全关闭资源）

// 如果函数参数过多，可将其重构为一个复合结构类型，
// 也算变相实现可选参数和命名实参功能
type serverOption struct {
	address string
	port int
	path string
	timeout time.Duration
	log     *log.Logger
}

// 将过多的参数独立成option struct ，
// 既便于扩展参数集，也方便通过newOption函数设置默认配置。


func newOption() *serverOption {
	return &serverOption{    // 默认值
		address: "127.0.0.1",
		port: 8080,
		path: "/",
		timeout: 10 * time.Second,
		log: nil,
	}
}

func server(option *serverOption){}

// 4.2.2 变参
// 本质是一个切片，智能接收一到多个同类型参数，必须放到列表尾部

func testVarArgs(s string, args ...int) {
	fmt.Println(s)
    fmt.Printf("%T,%v\n",args, args)
}

// 将切片作为变参，必须进行展开操作。如果是数组，先将其转换为切片

// 4.3 返回值
func div(x,y int) (int, error) {
	if y == 0 {
		return 0, errors.New("除数不能为0")
	}
	return x/y, nil
}

// 4.3.2 命名返回值
func div2(x,y int) (z int, err error) {
	if y == 0 {
		err = errors.New("除数不能为0")
		return
	}
	z = x/y
	return
}

func main() {
    f := hello // 将函数赋值给变量
	exex(f)

	var a *int = test()    // 调用 test()，接收返回的指针
	fmt.Println(a, *a)     // 打印指针地址和指针指向的值

	// 参数
	x := 100
	p := &x
	test1(p)

	opt := newOption()
	opt.port = 8081   // 命名参数设置值
	server(opt)

	testVarArgs("--hello world",1,2,3,4)

	// 返回值
	x, err := div(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}
}