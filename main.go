package main

import (
	"fmt"
)

func main() {
	//变量
	// variable()
	//数组
	//arrayTest()
	//map
	//mapTest()
	//struct
	//structTest()
	//路由匹配
	in()
	get("get/testing","get-true")
	post("post/testing","post-true")
	matchRoute()
}

// --------------------------- 变量 -----------------------------------

func variable() () {
	var name = "张三"
	age := 33
	fmt.Printf("%s\n", name, age)

}

//---------------------------- 数组 map -------------------------------

func arrayTest() {

	//var arr [5] string
	//arr[1] = "张三ARRAY"

	//arr2:= [5] int{1,2}

	var slice [] string
	var arr [10] string
	slice = arr[:3]
	arr[1] = "一"
	arr[2] = "二"
	arr[3] = "三"

	fmt.Printf("%s\n", slice[1])

}

func mapTest() {
	//另一种 map 的声明方式
	numbers := make(map[string]int)
	numbers["one"] = 1  // 赋值
	numbers["ten"] = 10 // 赋值
	numbers["three"] = 3

	fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
}

//--------------------------- struct 类型 -----------------------------

type personal struct {
	age  int
	name string
	sex  string
	info otherInfo
}

type otherInfo struct {
	address string
	city    string
	hobby   string
}

func structTest() {
	var p personal
	p.age = 18
	p.name = "李四"
	p.sex = "男"
	p.info = otherInfo{address: "四川", city: "成都", hobby: "打游戏"}
	fmt.Println("输出: ", p.info.city) // 读取数据
}

//----------------------------- 路由测试 ------------------------

var (
	allRoutes map[string]routes
)

type routes [] route

type route struct {
	url    string
	method string
}

func in() {
	allRoutes = make(map[string]routes)
}

func matchRoute() {

	fmt.Println("输出: ", allRoutes["GET"][0].url)
}

func get(url string, method string) {
	allRoutes["GET"] = append(routes{route{url: url, method: method}})
}

func post(url string, method string) {
	allRoutes["POST"] = append(routes{route{url: url, method: method}})
}
