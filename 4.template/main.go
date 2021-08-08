package main

import "fmt"

// 手机接口
type Phone interface {
	Unlock()
	Transfer()
	Verification()
}

// 组合手机接口
type IPhone struct {
	Phone
}

// 因为Golang嵌入接口的特殊性，需要传入一个实现了嵌入接口的对象
func NewIPhone(phone Phone) *IPhone {
	return &IPhone{phone}
}

// 支付，模版方法
func (p *IPhone) WechatPay() {
	p.Phone.Unlock()
	p.Phone.Transfer()
}

// 只有TouchID的IPhone7
type IPhone7 struct {
	IPhone
}

// 使用TouchID解锁
func (p *IPhone7) Unlock() {
	fmt.Println("TouchID验证通过，解锁手机成功！")
}

func (p *IPhone7) Transfer() {
	fmt.Println("发起微信转账，请完成支付验证")
	p.Verification()
}

// 使用TouchID验证
func (p *IPhone7) Verification() {
	fmt.Println("使用TouchID完成验证，转账成功！")
}

// 只有FaceID的IPhone7
type IPhone12 struct {
	IPhone
}

// 使用FaceID解锁
func (p *IPhone12) Unlock() {
	fmt.Println("FaceID验证通过，解锁手机成功！")
}

func (p *IPhone12) Transfer() {
	fmt.Println("发起微信转账，请完成支付验证")
	p.Verification()
}

// 使用FaceID验证
func (p *IPhone12) Verification() {
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
