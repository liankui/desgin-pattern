package main

import "fmt"

// 目标接口
type target interface {
	action()
}

type A struct{}

func (a A) action() {
	fmt.Println("action A")
}

// B 被适配者
type B struct{}

func (b B) specificAction() {
	fmt.Println("action B")
}

// BAdapter 适配器，B未实现insertA接口，使用BAdapter包一层B，实现insertA接口
type BAdapter struct {
	b *B
}

func (ba BAdapter) action() {
	ba.b.specificAction()
}

type Client struct{}

func (c Client) Action(en target) {
	en.action()
}

func main() {
	c := &Client{}

	a := &A{}
	c.Action(a)

	b := &B{}
	ba := &BAdapter{b: b}
	c.Action(ba)
}
