package main

import "fmt"

// Implementor 实现部分
type Implementor interface {
	action()
}

type A struct{}

func (a A) action() {
	fmt.Println("print A")
}

type B struct{}

func (b B) action() {
	fmt.Println("print B")
}

// Abstraction 抽象部分
type Abstraction interface {
	Action()
}

// RefinedAbstraction 具体抽象部分
type RefinedAbstraction struct {
	Imp Implementor
}

func (ref RefinedAbstraction) Action() {
	ref.Imp.action()
}

func main() {
	aRef := &RefinedAbstraction{A{}}
	aRef.Action()

	bRef := &RefinedAbstraction{B{}}
	bRef.Action()
}
