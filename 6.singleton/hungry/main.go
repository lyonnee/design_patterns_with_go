package main

import "fmt"

// 声明一个结构体
type IPhone struct {
	Version int
}

// 用来存储IPhone的实例
var iphone *IPhone

// init函数保证了GetInstance()被调用前完成对象初始化 或可以直接 var iphone  = &IPhone{}
func init() {
	iphone = &IPhone{Version: 7}
	// 输出iphoen对象指针地址
	fmt.Printf("初始化iphone指针：%p \n", iphone)
}

// 获取IPhone的实例
func GetInstance() *IPhone {
	return iphone
}

func main() {
	// 第一次获取实例
	iphone1 := GetInstance()
	// 验证获取到了实例
	fmt.Printf("这是IPhone%v \n", iphone1.Version)

	// 第二次获取实例
	iphone2 := GetInstance()

	//输出两次实例的指针地址，判断是否是同一个对象
	fmt.Printf("iphone1指针：%p \n", iphone1)
	fmt.Printf("iphone2指针：%p", iphone2)
}
