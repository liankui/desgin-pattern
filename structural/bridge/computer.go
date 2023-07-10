package main

import "fmt"

/* https://golangbyexample.com/bridge-design-pattern-in-go
桥接设计模式是一种结构设计模式，允许将抽象与其实现分开。听起来令人困惑？别担心，随着我们前进，情况会更加清晰。
这种模式建议将一个大类划分为两个独立的层次结构
	抽象-它是一个接口，抽象的子项被称为精炼抽象。抽象包含对实现的引用。
	实施-它也是一个接口，实施的儿童被称为具体实施
客户端正在引用抽象层次结构，而不必担心实现。让我们举个例子。假设你有两种类型的computer mac和windows。此外，让我们假设两种类型的打印机epson和hp。
计算机和打印机都需要以任何组合相互配合。客户端只会访问计算机，而不必担心打印的情况。我们没有为2*2的组合创建四个结构，而是创建两个层次结构
	抽象层次结构
	实施层次结构
见下图。这两个层次结构通过一个桥相互通信，其中Abstraction（此处为计算机）包含对Implementation（此处为打印机）的引用。
抽象和实现都可以继续独立发展，而不会相互影响。注意mac和windows如何嵌入对打印机的引用。我们可以在运行时更改Abstraction的实现
（即计算机的打印机），因为抽象是指通过接口实现。在调用mac.print()或windows.print()时，它会将请求发送到printer.printFile()。
这充当桥梁，并在两者之间提供松散的耦合。
*/

type computer interface {
	print()
	setPrinter(printer)
}

type mac struct {
	printer printer
}

func (m *mac) print() {
	fmt.Println("print request for mac")
	m.printer.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

type windows struct {
	printer printer
}

func (w *windows) print() {
	fmt.Println("print request for windows")
	w.printer.printFile()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}

type printer interface {
	printFile()
}

type epson struct{}

func (e *epson) printFile() {
	fmt.Println("printing by a epson printer")
}

type hp struct{}

func (h *hp) printFile() {
	fmt.Println("printing by a hp printer")
}

func main() {
	hpPrinter := &hp{}
	epsonPrinter := &epson{}
	macComputer := &mac{}
	winComputer := &windows{}

	macComputer.setPrinter(hpPrinter)
	macComputer.print()

	winComputer.setPrinter(epsonPrinter)
	winComputer.print()
}

/* Output:
print request for mac
printing by a hp printer
print request for windows
printing by a epson printer
*/
