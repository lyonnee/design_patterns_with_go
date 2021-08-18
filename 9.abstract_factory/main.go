package main

import "fmt"

// 手机接口
type Phone interface {
	Call()
}

// 平板电脑接口
type Pad interface {
	PlayGame()
}

// 电子工厂接口，生成电子产品，包括 手机、平板等
type ElectronicsFactory interface {
	CreatePhone() Phone
	CreatePad() Pad
}

// 苹果工厂，具体工厂
type AppleFactory struct{}

func (f *AppleFactory) CreatePhone() Phone {
	return &IPhone{}
}

func (f *AppleFactory) CreatePad() Pad {
	return &IPad{}
}

// 苹果手机
type IPhone struct{}

func (p *IPhone) Call() {
	fmt.Println("正在使用 IPhone 打电话")
}

// 苹果平板电脑
type IPad struct{}

func (p *IPad) PlayGame() {
	fmt.Println("正在使用 IPad 打游戏")
}

// 华为工厂，具体工厂
type HuaweiFactory struct{}

func (f *HuaweiFactory) CreatePhone() Phone {
	return &Mate40{}
}

func (f *HuaweiFactory) CreatePad() Pad {
	return &MatePad{}
}

// 华为手机
type Mate40 struct{}

func (p *Mate40) Call() {
	fmt.Println("正在使用 Mate40 打电话")
}

// 华为平板电脑
type MatePad struct{}

func (p *MatePad) PlayGame() {
	fmt.Println("正在使用 MatePad 打游戏")
}

func main() {
	// 苹果工厂
	appleFactory := &AppleFactory{}
	phone1 := appleFactory.CreatePhone()
	phone1.Call()

	pad1 := appleFactory.CreatePad()
	pad1.PlayGame()

	// 华为工厂
	huaweiFactory := &HuaweiFactory{}
	phone2 := huaweiFactory.CreatePhone()
	phone2.Call()

	pad2 := huaweiFactory.CreatePad()
	pad2.PlayGame()
}
