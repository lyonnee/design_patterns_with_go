package main

import (
	"fmt"
	"sync"
)

type IPhone struct {
	Version int
}

// 用来存储IPhone的实例
var iphone *IPhone

// once.Do可以保证函数只可能执行一次，防止高并发场景下多次执行实例化方法
var once sync.Once

func GetInstance() *IPhone {
	if iphone == nil {
		once.Do(func() {
			iphone = &IPhone{Version: 13}
			// 输出iphoen对象指针地址
			fmt.Printf("初始化iphone指针：%p \n", iphone)
		})
	}

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
