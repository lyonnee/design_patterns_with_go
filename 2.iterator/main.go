package main

import "fmt"

// 迭代器接口
type Iterator interface {
	// 判断是否还有值
	HasNext() bool
	// 取下一个值
	Next() interface{}
}

// 集合接口
type Aggregate interface {
	// 返回一个迭代器
	GetIterator() Iterator
}

// 书
type Book struct {
	// 名字
	Name string
	// 作者
	Author string
}

// 书架，一个存放书籍的集合
type BookShelf struct {
	books []Book
}

// 书架实现集合接口
func (bs *BookShelf) GetIterator() Iterator {
	// 每次都会创建一个新的迭代器，从 index=0 位置开始遍历
	return &BookShelfIterator{bookShelf: bs}
}

// 书架迭代器
type BookShelfIterator struct {
	// 当前遍历位置
	index int
	// 所指向书架
	bookShelf *BookShelf
}

// 书架迭代器实现迭代器接口
func (i *BookShelfIterator) HasNext() bool {
	return i.index < len(i.bookShelf.books)
}

// 书架迭代器实现迭代器接口
func (i *BookShelfIterator) Next() interface{} {
	book := i.bookShelf.books[i.index]
	i.index++

	return book
}

func main() {
	// 声明一个书架对象，并添加一些书籍
	bookShelf := &BookShelf{
		books: []Book{
			{Name: "以太坊技术详解与实战", Author: "闫莺，郑恺，郭众鑫"},
			{Name: "大话代码架构", Author: "田伟，郎小娇"},
			{Name: "GO语言公链开发实战", Author: "郑东旭，杨明珠，潘盈瑜，翟萌"},
			{Name: "区块链原理、设计与应用", Author: "杨保华，陈昌"},
			{Name: "精通区块链智能合约开发", Author: "熊丽兵"},
			{Name: "C程序设计", Author: "谭浩强"},
		},
	}

	// 获取书架的迭代器
	iterator := bookShelf.GetIterator()

	// 遍历，直到 没有值（HasNext == flase）
	for iterator.HasNext() {
		book := iterator.Next().(Book)
		fmt.Printf("书名：%s, 作者：%s \n", book.Name, book.Author)
	}
}
