package main

import ("fmt";"time")

// 6、 方法 可以为当前包内的任意类型定义方法。

type X int

func (x *X) inc() {                  // 名称前的参数称作receiver，作用类似python self
    *x++
}

// 还可直接调用匿名字段的方法，这种方式可实现与继承类似的功能。
type user struct {
	name string
	age byte
}

func (u user) ToString() string{
	fmt.Println(fmt.Sprintf("%+v",u))
    return fmt.Sprintf("name:%s, age:%d", u.name, u.age)
}

// 为 user 实现 Print 方法（满足 Printer 接口）
func (u user) Print() {
    fmt.Printf("User: %s, Age: %d\n", u.name, u.age)
}

type manager struct {
	user
	title string
}

func structDemo() {
    m := manager{
        user: user{
            name: "tom",
            age: 18,
        },
        title: "CTO",
    }
    fmt.Println(m.ToString())
}

// 7、接口
// 接口采用了duck type方式，也就是说无须在实现类型上添加显式声明
type  Printer interface {
    Print()
}

func interfaceDemo() { 
	var u user
	u.name = "tom"
	u.age = 18

	var p Printer = u // 只要包含接口所需的全部方法，既表示实现了该接口
	p.Print()
	
}

// 8、并发
//  整个运行时完全并发化设计。
// 凡能看到的，几乎都在以goroutine方式运行。
// 这是一种比普通协程或线程更加高效的并发设计，能轻松创建和运行成千上万的并发任务。

func task(id int){
	for i :=0; i<5; i++{
		fmt.Printf("%d: %d\n",id, i)
		time.Sleep(time.Second)
	}
}

// 通道(channel)与goroutine搭配，
// 实现用通信代替内存共享的CSP模型。

// 消费者
func consumer(data chan int, done chan bool) {
    for x := range data {      // 接收数据，直到通道被关闭
        println("recv:", x)
    }
    done <- true               // 通知main，消费结束
}

// 生产者
func producer(data chan int) {
    for i := 0; i < 4; i++ {
        data <- i                // 发送数据
    }
    close(data)                  // 生产结束，关闭通道
}

func chanDemo() {
    done := make(chan bool)            // 用于接收消费结束信号
    data := make(chan int)             // 数据管道

    go consumer(data, done)            // 启动消费者
    go producer(data)                  // 启动生产者
    <-done                            // 阻塞，直到消费者发回结束信号
}
func main() {
    var x X
    x.inc()
    fmt.Println(x)
	structDemo()
	interfaceDemo()

	fmt.Println("---------hello world")
	// go task(1)
	// go task(2)
	// time.Sleep(time.Second * 6)

	chanDemo()
}
