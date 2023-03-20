package main

import (
	"fmt"
)

/* 原型模式 https://golangbyexample.com/prototype-pattern-go/
这是一种创造性的设计模式，可让您创建对象的副本。
在此模式中，创建克隆对象的责任被委派给要克隆的实际对象。
要克隆的对象公开了一个克隆方法，该方法返回该对象的克隆副本。

何时使用：
1.当要克隆的对象的创建过程很复杂时，我们使用原型模式，即克隆可能涉及处理深副本，分层副本等的情况。此外，也可能有一些私人成员无法直接访问。
2.创建对象的副本，而不是从头开始创建新实例。这样可以防止在创建新对象 (例如数据库操作) 时涉及昂贵的操作。
3.当您想要创建一个新对象的副本时，但它仅作为接口提供给您。因此，您不能直接创建该对象的副本。

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
