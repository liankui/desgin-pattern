package abstractFactory

import "fmt"

/*
https://golangbyexample.com/abstract-factory-design-pattern-go/
抽象工厂设计模式是一种创造性的设计模式，可让您创建一系列相关对象。它是工厂模式的抽象。最好用一个例子来解释。
假设我们有两个工厂
	nike
	adidas
想象一下，您需要购买一个有鞋子和短裤的运动套件。最好是在大多数情况下，您都想购买类似工厂的全套运动套件，
即耐克或阿迪达斯。这就是抽象工厂进入画面的具体产品，你想要的是鞋和一个短袖，这些产品将由耐克和阿迪达斯的抽象工厂创造。
这两个工厂-耐克和阿迪达斯都实现了iSportsFactory接口。
我们有两个产品接口:
	iShoe-此接口由nikeShoe和adidasShoe具体产品实现。
	iShort-此接口由nikeShort和adidasShort具体产品实现。
*/
type iSportsFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}
	if brand == "nike" {
		return &nike{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}
