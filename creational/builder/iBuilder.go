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

模块
1.	Product（产品）：这是正在构建的复杂对象（包含 PartA、PartB 部分）。
2.	Builder 接口：定义了构建产品各部分的方法。
3.	ConcreteBuilder（具体建造者）：实现了 Builder 接口，负责构建和组装产品的各个部分。
4.	Director（指导者）：使用 Builder 来构建产品，封装了构建过程。
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
