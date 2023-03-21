package main

import (
	"fmt"
)

/* 工厂设计模式 https://golangbyexample.com/golang-factory-design-pattern/
是一种创造性的设计模式，此模式提供了一种隐藏正在创建的实例的创建逻辑的方法。
客户端仅与工厂结构交互，并告诉需要创建的实例类型。工厂类与相应的具体结构交互，并返回正确的实例。

我们有iGun接口，它定义了枪支应该拥有的所有方法:
1.实现iGun接口的gun结构
2.两支具体的枪ak47和maverick。两者都嵌入了iGun结构，因此也间接实现了所有的iGun方法，因此属于iGun类型。
3.我们有一个枪厂结构，可以制造ak47型或maverick的枪。
4.main.go充当客户端，而不是直接与ak47或maverick交互，而是依靠gunFactory创建ak47和maverick的实例。
*/
type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getPower() int {
	return g.power
}

func getGun(gunType string) (iGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "maverick":
		return newMaverick(), nil
	default:
		return nil, fmt.Errorf("wrong gun type passed")
	}
}

type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

type maverick struct {
	gun
}

func newMaverick() iGun {
	return &maverick{
		gun: gun{
			name:  "Maverick gun",
			power: 5,
		},
	}
}

func main() {
	ak47, _ := getGun("ak47")
	maverick, _ := getGun("maverick")
	printDetails(ak47)
	printDetails(maverick)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s\n", g.getName())
	fmt.Printf("Power: %d\n", g.getPower())
}

/*
Gun: AK47 gun
Power: 4
Gun: Maverick gun
Power: 5
*/
