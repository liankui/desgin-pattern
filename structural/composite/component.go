package main

import "fmt"

/* https://golangbyexample.com/composite-design-pattern-golang/
复合设计模式是一个结构设计模式。当我们希望一组名为“复合”的对象以类似于单个对象的方式处理时，
就会使用组合设计模式。它属于结构设计模式，因为它允许您将对象组合成树形结构。树结构中的每个
单个对象都可以以相同的方式处理，无论它们是复杂还是原始。
让我们试着用操作系统文件系统的例子来理解它。在文件系统中，有两种类型的对象文件和文件夹。在
某些情况下，文件和文件夹的处理方式相同。随着我们前进，情况会更加清晰。

何时使用
1.当需要从客户的角度以相同的方式处理复合和单个对象时，复合设计模式是有意义的。
–在我们上面的文件系统示例中，假设需要执行特定关键字的搜索操作。现在，此搜索操作适用于文件和
文件夹。对于一个文件，它只会查看文件的内容，对于文件夹，它将浏览该文件夹中层次结构中的所有文件，以找到该关键字
2.当复合和单个对象形成树状结构时，使用此模式。
在我们的示例中，文件和文件夹确实形成了一个树形结构
组件-它是定义复合和叶对象的常见操作的接口
复合-它实现了组件接口，并嵌入了子组件数组
叶子-它是树中的原始对象。它还实现了组件接口
*/

type component interface {
	search(string)
}

type folder struct {
	components []component
	name       string
}

func (fo *folder) search(keyword string) {
	fmt.Printf("searching recursively for keyword %s in folder %s\n", keyword, fo.name)
	for _, c := range fo.components {
		c.search(keyword)
	}
}

func (fo *folder) add(c component) {
	fo.components = append(fo.components, c)
}

type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Printf("searching for keyword %s in file %s\n", keyword, f.name)
}

func (f *file) getName() string {
	return f.name
}

func main() {
	f1 := &file{name: "file1"}
	f2 := &file{name: "file2"}
	f3 := &file{name: "file3"}

	fo1 := &folder{name: "folder1"}
	fo1.add(f1)

	fo2 := &folder{name: "folder2"}
	fo2.add(f2)
	fo2.add(f3)
	fo2.add(fo1)
	fo2.search("rose")
}

/* Output:
searching recursively for keyword rose in folder folder2
searching for keyword rose in file file2
searching for keyword rose in file file3
searching recursively for keyword rose in folder folder1
searching for keyword rose in file file1
*/
