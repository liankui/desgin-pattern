package main

import (
	"fmt"
)

/*
	https://golangbyexample.com/prototype-pattern-go/

原型模式（Prototype Pattern）是一种创建型设计模式，它允许对象通过复制现有对象（称为原型）来创建新对象，而不是通过类构造函数实例化。
这种模式的核心思想是“克隆”现有的实例，以避免创建过程中的开销，特别是当对象的创建过程比较复杂时。

原型模式的主要作用是：
1.提高对象创建的效率：避免复杂对象的重复创建，直接复制现有实例。
2.动态定制对象：允许在运行时动态地创建对象，不必依赖于类来生成对象实例。
3.减少子类数目：通过克隆机制减少需要定义的子类数目。

适用于以下场景：
1.对象初始化复杂或耗时：如对象创建需要读取配置、数据库操作等。
2.需要大量相似对象：例如需要大量配置相似但部分属性不同的对象。
3.避免创建类时的性能瓶颈：例如在游戏开发中，复制一个已经配置好的角色原型。
4.与复合对象结合使用：如组合模式中，复制复杂的对象树。

demo：
	os文件系统有文件和文件夹，文件夹本身包含文件和文件夹。
	每个文件和文件夹都可以用inode接口表示。inode接口也有clone()函数。
*/

type inode interface {
	print(string)
	clone() inode
}

type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *file) clone() inode {
	return &file{name: f.name + "_clone"}
}

type folder struct {
	children []inode
	name     string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation) // 如果是file类型，打印"入参+i.name"；如果是folder，递归
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildren []inode
	for _, i := range f.children {
		copy := i.clone() // 如果是file类型，返回&file{name: f.name + "_clone"}；如果是folder，递归
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}
	folder1 := &folder{
		children: []inode{file1},
		name:     "Folder1",
	}
	folder2 := &folder{
		children: []inode{folder1, file2, file3},
		name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("_")
	cloneFolder := folder2.clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("_")
}

/*
Printing hierarchy for Folder2
  Folder2
    Folder1
        File1
    File2
    File3

Printing hierarchy for clone Folder
  Folder2_clone
    Folder1_clone
        File1_clone
    File2_clone
    File3_clone
*/
