package main

import "fmt"

// Builder 接口，定义生成实例的方法
type ComputerBuilder interface {
	// 添加中央处理器
	AddCPU()
	// 添加内存
	AddMemory()
	// 添加主板
	AddMainboard()
	// 添加硬盘
	AddHardDisk()
	// 添加显卡
	AddGPU()

	// 获取电脑主机实例
	GetComputer() Computer
}

// 电脑结构， 主要描述电脑有哪些重要组件，没有方法
type Computer struct {
	CPU       string
	Memory    string
	Mainboard string
	HardDisk  string
	GPU       string
}

// 办公型电脑主机
type OffceModelBuilder struct {
	Computer
}

// 实现ComputerBuilder的方法
func (om *OffceModelBuilder) AddCPU() {
	om.CPU = "Intel-i3"
}

func (om *OffceModelBuilder) AddMemory() {
	om.Memory = "8G"
}

func (om *OffceModelBuilder) AddMainboard() {
	om.Mainboard = "B460"
}

func (om *OffceModelBuilder) AddHardDisk() {
	om.HardDisk = "256G HDD"
}

func (om *OffceModelBuilder) AddGPU() {
	om.GPU = "Intel HD630"
}

func (om *OffceModelBuilder) GetComputer() Computer {
	return Computer{
		CPU:       om.CPU,
		Memory:    om.Memory,
		Mainboard: om.Mainboard,
		HardDisk:  om.HardDisk,
		GPU:       om.GPU,
	}
}

// 游戏型电脑主机
type GameModelBuilder struct {
	Computer
}

// 实现ComputerBuilder的方法
func (gm *GameModelBuilder) AddCPU() {
	gm.CPU = "Intel-i9"
}

func (gm *GameModelBuilder) AddMemory() {
	gm.Memory = "32G"
}

func (gm *GameModelBuilder) AddMainboard() {
	gm.Mainboard = "X460"
}

func (gm *GameModelBuilder) AddHardDisk() {
	gm.HardDisk = "512G SSD，1T HDD"
}

func (gm *GameModelBuilder) AddGPU() {
	gm.GPU = "NVIDIA GTX3090"
}

func (gm *GameModelBuilder) GetComputer() Computer {
	return Computer{
		CPU:       gm.CPU,
		Memory:    gm.Memory,
		Mainboard: gm.Mainboard,
		HardDisk:  gm.HardDisk,
		GPU:       gm.GPU,
	}
}

// 指挥者
type Director struct {
	builder ComputerBuilder
}

// 建造方法，定义步骤及顺序（类似于Template Method
func (d *Director) Build() Computer {
	d.builder.AddCPU()
	d.builder.AddMemory()
	d.builder.AddMainboard()
	d.builder.AddHardDisk()
	d.builder.AddGPU()

	return d.builder.GetComputer()
}

func main() {
	// 办公型电脑建造者
	offceModelBuilder := &OffceModelBuilder{}

	// 实例化一个指挥者，并传入具体建造者
	director := Director{offceModelBuilder}
	computer1 := director.Build()

	fmt.Printf("办公型主机配置：%v \n", computer1)

	// 游戏型电脑建造者
	gameModelBuilder := &GameModelBuilder{}

	// 这里直接赋值，也证明指挥者 Director不直接依赖具体建造者
	director.builder = gameModelBuilder
	computer2 := director.Build()

	fmt.Printf("游戏型主机配置：%v \n", computer2)
}
