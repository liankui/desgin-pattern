package main

import "fmt"

// Abstract Interface
type iAlpha interface {
	work()
	common()
}

// Abstract Concrete Type
type alpha struct {
	name string
}

func (a *alpha) common() {
	fmt.Println("common called")
}

// Implementing Type
type beta struct {
	alpha
}

func (b *beta) work() {
	fmt.Println("work called")
	fmt.Printf("name is %s\n", b.name)
	b.common()
}

func main() {
	a := alpha{
		name: "test",
	}
	b := &beta{
		alpha: a,
	}
	b.work()
}

/* Output:
work called
name is test
common called
*/
