package main

import "fmt"

// 产品接口（原型接口）
type Product interface {
	// 用于演示的方法
	Use()
	// 克隆方法，返回自身的副本（不是指针）
	Clone() Product
}

// 产品管理
type Manager struct {
	// 原型池
	products map[string]Product
}

// 注册原型到原型池
func (m *Manager) Register(name string, product Product) {
	_, ok := m.products[name]
	if ok {
		return
	}

	m.products[name] = product
}

// Clone一个对象，返回其副本
func (m *Manager) Clone(name string) Product {
	v, ok := m.products[name]
	if !ok {
		return nil
	}

	// 调用被克隆对象自身的克隆方法
	return v.Clone()
}

// 演示产品，实现了原型接口
type IPhone struct {
	Name string
}

// 演示方法
func (p *IPhone) Use() {
	fmt.Printf("正在使用 %s 打电话\n", p.Name)
}

// 克隆的细节
func (p *IPhone) Clone() Product {
	return &IPhone{Name: p.Name + "_clone"}
}

func main() {
	iphone7 := IPhone{Name: "IPhone7"}
	iphone7.Use()
	// 输出指针地址做比较
	fmt.Printf("%s 的指针式 %p \n\n", iphone7.Name, &iphone7)

	manager := Manager{products: make(map[string]Product)}
	manager.Register("iphone7", &iphone7)

	clone_iphone7 := manager.Clone("iphone7")
	clone_iphone7.Use()
	// 输出指针地址做比较
	fmt.Printf("clone_iphone7 的指针式 %p \n", &clone_iphone7)
}
