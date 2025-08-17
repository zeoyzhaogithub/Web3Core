package main
import ( "errors"; "fmt")
// 4、函数

func div(a, b int)(int, error){
	if b == 0{
		return 0, errors.New("除数不能为0")
	}
	return a/b, nil
}

// 函数是第一类型，可作为参数或者返回值
func add(a, b int) func(){
	return func() {
		var c = a+b
		fmt.Println("a+b=",c) 
	}
}

// 用defer定义延迟调用，无论函数是否出错，他都确保结束前辈调用
func deferDemo(a,b int){
	defer func(){
		fmt.Println("defer....")
	}()
	fmt.Println(a/b)
}

// 5、数据
// 切片slice 可实现类似动态数组的功能

func sliceDemo(){
	x := make([]int, 0, 5)  // 创建一个长度为0，容量为5的切片
    fmt.Println("初始长度：",len(x), "初始容量：",cap(x))

	for i:=0; i<8; i++{
		// 追加数据 当超出容量限制时，会自动分配更大的存储空间
		x = append(x, i)  
	}
	fmt.Println("当前长度：",len(x), "当前容量：",cap(x))
	fmt.Println(x)
}

// 将字典map类型内置，可直接从运行时层面获得性能优化

// 所谓ok-idiom模式，是指在多返回值中用一个名为ok的布尔值来标示操作是否成功。
// 因为很多操作默认返回零值，所以须额外说明。
func mapDemo(){
	m := make(map[string]int)  // 创建字典类型对象
	m["name"] = 2
	m["age"] = 18
    
	// 使用ok-idiom 获取值，可知道key、value是否存在
	x, ok := m["b"]
	fmt.Println("x=",x,"ok=",ok)
}

// 结构体struct可匿名嵌入其他类型

type user struct{  // 结构体
	name string
	age int 
}

type manager struct{
	user  // 匿名嵌入
	title string
}

func structDemo(){
	m := manager{
		user:user{
			name:"tom",
			age:18,
		},
		title:"ceo",
	}
	fmt.Println(m)
}


func main(){
    a,b := 10,2
	c, err := div(a,b)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(a,"/",b,"=",c)
	}
	fmt.Println("----1----")
	add(a,b)()
	deferDemo(a,b)
	//deferDemo(a,0)
	sliceDemo()
	mapDemo()
	fmt.Println("---2-----")
	structDemo()
}
