package main

import (
	"fmt"
	"sync"
)

type Fruit interface {
	getPrice() int
}

func GetFruit(name string) Fruit {
	switch name {
	case "apple":
		return NewApple()
	case "pear":
		return NewPear()
	default:
		return nil
	}
}

type Apple struct{}

func NewApple() *Apple {
	return &Apple{}
}

func (a *Apple) getPrice() int {
	return 5
}

type Pear struct{}

func NewPear() *Pear {
	return &Pear{}
}

func (p *Pear) getPrice() int {
	return 2
}

func main() {
	_fruit := GetFruit("apple")
	if _fruit != nil {
		price := _fruit.getPrice()
		fmt.Println("price=", price)
	}

	f := fruits["apple"]
	if f != nil {
		price := f.getPrice()
		fmt.Println("price=", price)
	}

	Start(f)
	fmt.Println("price=", fruit.Fruit.getPrice())
}

// register 逻辑
var fruits map[string]Fruit

func RegisterFruit(name string, fruit Fruit) {
	if fruits == nil {
		fruits = make(map[string]Fruit)
	}

	fruits[name] = fruit
}

func init() {
	RegisterFruit("apple", NewApple())
	RegisterFruit("pear", NewPear())
}

// once 逻辑
var (
	fruit     *defaultFruit
	fruitOnce sync.Once
)

type defaultFruit struct {
	Fruit Fruit
}

func newDefaultFruit() *defaultFruit {
	return &defaultFruit{}
}

func Start(fru Fruit) {
	fruitOnce.Do(func() {
		fruit = newDefaultFruit()
	})

	fruit.Fruit = fru
}
