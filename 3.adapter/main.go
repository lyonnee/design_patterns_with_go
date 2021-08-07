package main

import "fmt"

// 电器设备接口
type Appliances interface {
	GetName() string
}

// 10A插头电器接口，所有10A插头的电器设备都要实现该接口
type Appliances10A interface {
	// 组合电器设备接口
	Appliances
	// 插头是10A的，实现该接口
	Plug10A()
}

// 16A插头电器接口，所有16A插头的电器设备都要实现该接口
type Appliances16A interface {
	Appliances
	// 插头是16A的，实现该接口
	Plug16A()
}

// 电视机
type TV struct {
	// 电器名称，比如： 空调，电视，冰箱
	Name string
}

func (tv *TV) GetName() string {
	return tv.Name
}

// 电视机是10A的插头
func (tv *TV) Plug10A() {}

// 空调
type HVAC struct {
	Name string
}

func (hvac *HVAC) GetName() string {
	return hvac.Name
}

// 空调是16A的插头
func (hvac *HVAC) Plug16A() {}

// 标准插线板
type PlugBoard struct {
	// 电器集合
	Appliances []Appliances
}

// 插入 10A插头的电器设备
func (p *PlugBoard) Insert10A(app Appliances10A) {
	p.Appliances = append(p.Appliances, app)
}

// 插线板通电，那么在插线板上的电器都通上电了
func (p *PlugBoard) TransferCurrent() {
	for _, app := range p.Appliances {
		fmt.Printf("%s 已通电 \n", app.GetName())
	}
}

// 适配器
type Adapter struct {
	// 正面（10A）插入的插线板
	PB PlugBoard
	// 反面（16A）接入的电器设备
	APP Appliances16A
}

// 适配器能够插入插线板，且能通电，所以本身也是一个电器设备
func (a *Adapter) GetName() string {
	return a.APP.GetName()
}

// 正面是 10A的插头
func (a *Adapter) Plug10A() {}

// 背面有 16A的插孔，所以可以插入一个16A插头的设备
func (a *Adapter) Insert16A(app Appliances16A) {
	a.APP = app
	a.PB.Insert10A(a)
}

func main() {
	// 电视机
	tv := TV{Name: "电视机"}
	// 空调
	hvac := HVAC{Name: "空调"}
	//插线板
	pb := PlugBoard{}

	// 插线板上插入 电视机插头
	pb.Insert10A(&tv)

	// 适配器
	adapter := Adapter{}

	//插线板上插入 适配器插头
	pb.Insert10A(&adapter)

	// 适配器上插入 空调插头
	adapter.Insert16A(&hvac)

	// 给插线板通电
	pb.TransferCurrent()
}
