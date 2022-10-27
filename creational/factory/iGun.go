package factory

/*
https://golangbyexample.com/golang-factory-design-pattern/
工厂设计模式是一种创造性的设计模式，也是最常用的模式之一。此模式提供了一种隐藏正在创建的实例的创建逻辑的方法。
客户端仅与工厂结构交互，并告诉需要创建的实例类型。工厂类与相应的具体结构交互，并返回正确的实例。

* 我们有iGun接口，它定义了枪支应该拥有的所有方法
* 有实现iGun接口的gun结构。
* 两支具体的枪ak47和maverick。两者都嵌入了iGun结构，因此也间接实现了所有的iGun方法，因此属于iGun类型。
* 我们有一个枪厂结构，可以制造ak47型或maverick的枪。
* main.go充当客户端，而不是直接与ak47或maverick交互，而是依靠gunFactory创建ak47和maverick的实例。
*/
type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}
