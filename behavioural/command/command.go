package main

import (
	"fmt"
)

/* https://golangbyexample.com/command-design-pattern-in-golang/
命令设计模式是一种行为设计模式。它建议将请求封装为独立对象。创建的对象具有有关请求的所有信息，因此可以独立执行它。
命令设计模式中使用的基本组件是:
	Receiver-它是包含业务逻辑的类。命令对象仅延迟其对接收器的请求。
	命令-嵌入接收器并绑定接收器的特定动作。
	Invoker-它嵌入命令，并通过调用命令的execute方法来实现命令。
	客户端-它使用适当的接收器创建命令，将接收器绕过命令的构造函数。之后，它还将生成的命令与调用程序相关联。

让我们了解一种情况，之后将清楚为什么命令模式有用。想象一下电视的情况。电视可以通过任何一个打开
	远程开启按钮
	打开电视上的按钮。
这两个触发点都做同样的事情，即打开电视。为了在电视上，我们可以用接收器作为电视来实现ON命令对象。当在这个on command对象
上调用execute() 方法时，它又调用TV.ON () 函数。所以在这种情况下:
	接收器是电视
	命令是嵌入电视的ON命令对象
	Invoker是电视上的遥控器ON按钮或ON按钮。都嵌入ON命令对象
请注意，我们已经将打开电视的请求包装为可以由多个invokers调用的on command对象。此ON command对象嵌入了接收器 (此处为TV)，并且可以独立执行。

在上面的例子中创建一个单独的命令对象有什么好处。
1.它将UI逻辑与底层业务逻辑解耦
2.无需为每个invokers创建不同的处理程序。
3.命令对象包含它需要执行的所有信息。因此，它也可以用于延迟执行。
*/

type command interface {
	execute()
}

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type device interface {
	on()
	off()
}

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

func main() {
	tv := &tv{}
	onCommand := &onCommand{
		device: tv,
	}
	offCommand := &offCommand{
		device: tv,
	}
	onButton := &button{
		command: onCommand,
	}
	onButton.press()
	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}

/*
Turning tv on
Turning tv off
*/
