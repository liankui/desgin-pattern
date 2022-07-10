package prototype

import "fmt"

type folder struct {
	childrens []inode
	name      string
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.childrens {
		i.print(indentation + indentation) // 如果是file类型，打印"入参+i.name";如果是folder，递归
	}
}

func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + "_clone"}
	var tempChildrens []inode
	for _, i := range f.childrens {
		copy := i.clone() // 如果是file类型，返回&file{name: f.name + "_clone"};如果是folder，递归
		tempChildrens = append(tempChildrens, copy)
	}
	cloneFolder.childrens = tempChildrens
	return cloneFolder
}

/*
由于文件和文件夹结构都实现了打印和克隆功能，因此它们的类型为inode。另外，请注意文件和文件夹中的克隆函数。
两者中的克隆函数都返回相应文件或文件夹的副本。克隆时，我们为名称字段附加关键字 “_clone”。让我们编写主要函数来测试事物。
*/
