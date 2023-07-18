package main

import "fmt"

/* https://golangbyexample.com/flyweight-design-pattern-golang/
享元设计模式是一种结构设计模式。当需要创建大量类似对象时，使用此模式。这些物体被称为飞量级物体，是不可变的。
让我们先看一个例子。在这个例子之后，飞重模式将变得清晰。
在《反恐精英》游戏中，恐怖分子和反恐精英有不同的服装类型。为了简单起见，让我们假设恐怖分子和反恐人员各有
一种服装类型。连衣裙对象嵌入在玩家对象中，如下所示
下面是玩家的结构，我们可以看到服装对象嵌入在玩家结构中
type player struct {
    dress      dress
    playerType string //Can be T or CT
    lat        int
    long       int
}
假设有5名恐怖分子和5名反恐人员，所以总共有10名玩家。现在关于着装有两种选择
1.10个玩家对象中的每一个都会创建一个不同的服装对象并嵌入它们。将创建总共10件服装对象
2.我们创造了两个服装对象
	单一恐怖分子服装对象：这将在5名恐怖分子之间共享
	单一反恐服装对象：这将在5名反恐中共享
正如你所能做到的那样，在方法1中，总共创建了10个服装对象，而在方法2中，只创建了2个服装对象。第二种方法是
我们在享元设计模式中遵循的方法。我们创建的两个服装对象被称为飞量级对象。飞重图案取出普通部件并创建飞重物体。
然后，这些飞量级对象（在这里穿着）可以在多个对象（此处的玩家）之间共享。这大大减少了服装对象的数量，好的是，
即使你创造了更多的玩家，仍然只有两个服装对象就足够了。

在飞量级模式中，我们将飞量级物体存储在地图中。每当创建共享flyweight对象的其他对象时，就会从地图上获取flyweight对象。

内在状态和外在状态
	内在状态-穿着内在状态，因为它可以在多个恐怖分子和反恐对象之间共享
	外部状态-玩家位置和玩家武器是一种外部状态，因为它对每个物体都不同。
何时使用：
	当对象具有一些可以共享的内在属性时。与上述示例中所述，着装是被取出并共享的内在属性。
	当需要创建大量可能导致内存问题的对象时，请使用flyweight。以防弄清楚所有共同或内在状态，并为此创建飞量级对象。
*/

const (
	TerroristDressType        = "tDress"
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]dress),
	}
)

type dressFactory struct {
	dressMap map[string]dress
}

func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}

func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	if v, ok := d.dressMap[dressType]; ok {
		return v, nil
	}

	switch dressType {
	case TerroristDressType:
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	case CounterTerroristDressType:
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	default:
		return nil, fmt.Errorf("wrong dress type passed")
	}
}

type dress interface {
	getColor() string
}

type terroristDress struct {
	color string
}

func newTerroristDress() *terroristDress {
	return &terroristDress{color: "red"}
}

func (t *terroristDress) getColor() string {
	return t.color
}

type counterTerroristDress struct {
	color string
}

func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{color: "green"}
}

func (t *counterTerroristDress) getColor() string {
	return t.color
}

type player struct {
	dress      dress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType, dressType string) *player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &player{dress: dress, playerType: playerType}
}

func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
