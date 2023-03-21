package main

import (
	"fmt"
)

/* 抽象工厂设计模式 https://golangbyexample.com/abstract-factory-design-pattern-go/
是一种创造性的设计模式，可让您创建一系列相关对象。它是工厂模式的抽象。

demo:
假设我们有两个工厂
	nike
	adidas
可以购买一个有鞋子和短裤的运动套件。最好是在大多数情况下，您都想购买类似工厂的全套运动套件，即耐克或阿迪达斯。
这就是抽象工厂进入画面的具体产品，你想要的是鞋和一个短袖，这些产品将由耐克和阿迪达斯的抽象工厂创造。
这两个工厂-耐克和阿迪达斯都实现了iSportsFactory接口。
有两个产品接口:
	iShoe-此接口由nikeShoe和adidasShoe具体产品实现。
	iShort-此接口由nikeShort和adidasShort具体产品实现。
*/
type iSportsFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	switch brand {
	case "adidas":
		return &adidas{}, nil
	case "nike":
		return &nike{}, nil
	default:
		return nil, fmt.Errorf("wrong brand type passed")
	}
}

type iShoe interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type shoe struct {
	logo string
	size int
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) setSize(size int) {
	s.size = size
}

func (s *shoe) getSize() int {
	return s.size
}

type iShort interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

type short struct {
	logo string
	size int
}

func (s *short) setLogo(logo string) {
	s.logo = logo
}

func (s *short) getLogo() string {
	return s.logo
}

func (s *short) setSize(size int) {
	s.size = size
}

func (s *short) getSize() int {
	return s.size
}

type adidas struct {
}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *adidas) makeShort() iShort {
	return &adidasShort{
		short: short{
			logo: "adidas",
			size: 4,
		},
	}
}

type adidasShoe struct {
	shoe
}

type adidasShort struct {
	short
}

type nike struct {
}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 15,
		},
	}
}

func (n *nike) makeShort() iShort {
	return &nikeShort{
		short: short{
			logo: "nike",
			size: 5,
		},
	}
}

type nikeShoe struct {
	shoe
}

type nikeShort struct {
	short
}

func main() {
	nikeFactory, _ := getSportsFactory("nike")
	adidasFactory, _ := getSportsFactory("adidas")
	nikeShoe := nikeFactory.makeShoe()
	nikeShort := nikeFactory.makeShort()
	adidasShoe := adidasFactory.makeShoe()
	adidasShort := adidasFactory.makeShort()
	printShoeDetails(nikeShoe)
	printShortDetails(nikeShort)
	printShoeDetails(adidasShoe)
	printShortDetails(adidasShort)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

func printShortDetails(s iShort) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

/*
Logo: nike
Size: 14
Logo: nike
Size: 14
Logo: adidas
Size: 14
Logo: adidas
Size: 14
*/
