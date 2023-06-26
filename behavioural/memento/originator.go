package main

import "fmt"

/* https://golangbyexample.com/memento-design-pattern-go/
纪念品设计模式是一种行为设计模式。它允许我们为对象保存检查点，从而允许对象恢复到其之前的状态。基本上，它有助于对对象进行撤销重做操作。

以下是Memento设计模式的设计组件。
	发起者：它是实际对象，其状态被保存为纪念品。
	纪念品：这是保存发起人状态的对象
	管理员：这是保存多个纪念品的对象。给定一个索引，它返回相应的纪念品。
发起者定义了两种方法。savememento()和restorememento()
	savememento()-在此方法中，发起人将其内部状态保存为memento对象。
	restorememento()-此方法将输入作为memento对象。发起人将自己恢复到通行证纪念品中。因此，恢复了以前的状态。
*/

type originator struct {
	state string
}

func (o *originator) createMemento() *memento {
	return &memento{state: o.state}
}

func (o *originator) restoreMemento(m *memento) {
	o.state = m.state
}

type memento struct {
	state string
}

func (m *memento) getSavedState() string {
	return m.state
}

// caretaker contains the mementoArray which holds all the memento.
type caretaker struct {
	mementoArray []*memento
}

func newCaretaker() *caretaker {
	return &caretaker{
		mementoArray: make([]*memento, 0),
	}
}

func (c *caretaker) addMemento(m *memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *caretaker) getMemento(index int) *memento {
	return c.mementoArray[index]
}

func main() {
	c := newCaretaker()

	ori := &originator{state: "a"}
	fmt.Printf("originator current state: %s\n", ori.state)
	c.addMemento(ori.createMemento())

	ori.state = "b"
	fmt.Printf("originator current state: %s\n", ori.state)
	c.addMemento(ori.createMemento())

	ori.state = "c"
	fmt.Printf("originator current state: %s\n", ori.state)
	c.addMemento(ori.createMemento())

	ori.restoreMemento(c.getMemento(0))
	fmt.Printf("restore to state: %s\n", ori.state)

	ori.restoreMemento(c.getMemento(1))
	fmt.Printf("restore to state: %s\n", ori.state)
}

/* Output:
originator current state: a
originator current state: b
originator current state: c
restore to state: a
restore to state: b
*/
