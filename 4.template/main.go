package main

import "fmt"

// 手机接口，所有方法都是私有的，就像现实生活中我们不会通过输入命令调用系统底层接口。这些都是APP封装好的。
type Phone interface {
	unlock()
	transfer()
	verification()
}

// 微信接口
type Wechat interface {
	// 微信有微信支付的功能，是对外公开的方法，会调用系统底层的接口
	WechatPay()
}

// 组合手机接口
type IPhone struct {
	Wechat
	Phone
}

// 因为Golang嵌入接口的特殊性，需要传入实现了嵌入接口的对象
func NewIPhone(phone Phone) *IPhone {
	return &IPhone{Phone: phone}
}

// IPhone类型（可以理解为父类）需要实现模版方法，支付
func (p *IPhone) WechatPay() {
	// 定义处理流程的框架，子类实现具体处理
	p.Phone.unlock()
	p.Phone.transfer()
}

// 只有TouchID的IPhone7
type IPhone7 struct {
	IPhone
}

// 使用TouchID解锁
func (p *IPhone7) unlock() {
	fmt.Println("TouchID验证通过，解锁手机成功！")
}

func (p *IPhone7) transfer() {
	fmt.Println("发起微信转账，请完成支付验证")
	p.verification()
}

// 使用TouchID验证
func (p *IPhone7) verification() {
	fmt.Println("使用TouchID完成验证，转账成功！")
}

// 只有FaceID的IPhone12
type IPhone12 struct {
	IPhone
}

// 使用FaceID解锁
func (p *IPhone12) unlock() {
	fmt.Println("FaceID验证通过，解锁手机成功！")
}

func (p *IPhone12) transfer() {
	fmt.Println("发起微信转账，请完成支付验证")
	p.verification()
}

// 使用FaceID验证
func (p *IPhone12) verification() {
	fmt.Println("使用FaceID完成验证，转账成功！")
}

func main() {
	fmt.Println("使用IPhone7微信支付")
	iphone7 := NewIPhone(&IPhone7{})
	iphone7.WechatPay()

	fmt.Println("\n使用IPhone12微信支付")
	iphone12 := NewIPhone(&IPhone12{})
	iphone12.WechatPay()
}
