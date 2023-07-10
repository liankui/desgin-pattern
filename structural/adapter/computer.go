package main

import "fmt"

/* https://golangbyexample.com/adapter-design-pattern-go/
适配器设计模式是结构设计模式。最好用一个例子来理解这个模式。假设你有两台笔记本电脑
	MacBook Pro
	Windows笔记本电脑
MacBook Pro有一个方形的USB端口，Windows有一个圆形的USB端口。作为客户，您有一根方形的USB电缆，因此它只能插入Mac笔记本电脑中。所以你看到了这里的问题:
我们有一个类（客户端）期望对象的一些功能（这里的方形USB端口），但我们有另一个名为 adaptee（这里的Windows Laptop）
的对象，它提供相同的功能，但通过不同的接口（圆形端口）
这就是适配器模式进入图片的地方。我们创建了一个被称为适配器的类，它将
	坚持客户端期望的相同接口（此处为方形USB端口）
	以被申请人期望的形式将客户的请求翻译给被申请人。基本上，在我们的示例中，它充当一个适配器，在方形端口中接受USB，然后插入Windows笔记本电脑的圆形端口。
何时使用
	当对象根据客户端的要求实现不同的接口时，使用此设计模式。
*/

type computer interface {
	insertInSquarePort()
}

type mac struct{}

func (m *mac) insertInSquarePort() {
	fmt.Println("insert square port into mac machine")
}

type windowsAdapter struct {
	windowsMachine *windows
}

func (wa *windowsAdapter) insertInSquarePort() {
	wa.windowsMachine.insertInCirclePort()
}

type windows struct{}

func (w *windows) insertInCirclePort() {
	fmt.Println("insert circle port into windows machine")
}

type client struct{}

func (c *client) insertSquareUsbComputer(com computer) {
	com.insertInSquarePort()
}

func main() {
	c := &client{}
	m := &mac{}
	c.insertSquareUsbComputer(m)

	w := &windows{}
	wa := &windowsAdapter{windowsMachine: w}
	c.insertSquareUsbComputer(wa)
}

/* Output:
insert square port into mac machine
insert circle port into windows machine
*/
