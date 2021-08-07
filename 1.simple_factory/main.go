package main

import "fmt"

type IPhone interface {
	LoginQQ()
}

type IPhone7 struct{}

func (p *IPhone7) LoginQQ() {
	fmt.Println("正在使用 IPhone7 登陆QQ")
}

type IPhone13 struct{}

func (p *IPhone13) LoginQQ() {
	fmt.Println("正在使用 IPhone13 登陆QQ")
}

type IPhone110 struct{}

func (p *IPhone110) LoginQQ() {
	fmt.Println("正在使用 IPhone110 登陆QQ")
}

//IPhone生产工厂，根据要求版本号生成新的IPhone
func NewIphone(version int) IPhone {
	switch version {
	case 7:
		return &IPhone7{}
	case 13:
		return &IPhone13{}
	case 110:
		return &IPhone110{}
	default:
		//该工厂不支持该版本号的IPhone生产，所以返回空
		return nil
	}
}

func main() {
	iphone := NewIphone(110)
	if iphone == nil {
		fmt.Println("创建新IPhone失败")
		return
	}

	iphone.LoginQQ()
}
