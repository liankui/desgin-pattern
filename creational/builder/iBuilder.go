package main

import (
	"fmt"
)

/* https://golangbyexample.com/builder-pattern-golang/
构建器模式是创造型的设计模式。
功能
1.逐步构建一个复杂对象。
2.将对象的构建过程与表现形式分离。
3.相同的构建过程能够创建对象的不同表现形式。
*/

type Product struct {
	PartA string
	PartB string
}

type Builder interface {
	BuildPartA()
	BuildPartB()
	GetProduct() Product
}

type ConcreteBuilder struct {
	product Product
}

func (c *ConcreteBuilder) BuildPartA() {
	c.product.PartA = "partA"
}

func (c *ConcreteBuilder) BuildPartB() {
	c.product.PartB = "partB"
}

func (c *ConcreteBuilder) GetProduct() Product {
	return c.product
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
}

func main() {
	builder := &ConcreteBuilder{}
	director := &Director{builder: builder}

	director.Construct()
	prod := builder.GetProduct()
	fmt.Println(prod)
	// output:
	// {partA partB}
}
