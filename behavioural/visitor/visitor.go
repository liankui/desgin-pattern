package main

import "fmt"

/* https://golangbyexample.com/visitor-design-pattern-go/
访客设计模式是一种行为设计模式，允许您在不实际修改结构的情况下向结构添加行为。
让我们用一个例子来了解访客模式。假设你是一个lib的维护者，该lib具有不同的形状结构，例如
	正方形
	椭圆形
	三角形
上述每个形状结构都实现了一个通用的接口形状。你的公司里有很多团队正在使用你的lib。
现在，假设团队中的一个人希望你在形状结构中再添加一个行为（getArea()）。使用访客模式可以很好的解决上述问题。
通过定义一个访客界面（visitor的interface）
type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForRectangle(*rectangle)
}
功能 visitForSquare（正方形）、visitForCircle（圆形）、visitForTriangle（三角形）允许我们分别为正方形、圆形和三角形添加功能。
现在想到的问题是，为什么我们不能在访客界面中有一个单一的方法访问（形状）。我们没有它的原因是因为GO和其他一些语言支持方法超载。因此，每个结构都有不同的方法。
我们用下面的签名在形状接口中添加一个接受方法，每个形状结构都需要定义此方法。
func accept(v visitor)
我们刚刚提到我们不想修改现有的形状结构。但是，在使用访客模式时，我们确实必须修改形状结构，但这种修改只会进行一次。如果添加任何额外的行为，
如getNumSides()，getMiddleCoordinates()将使用相同的accept(v visitor)函数，而无需进一步更改形状结构。基本上，形状结构只需要
修改一次，所有未来其他行为的请求都将使用相同的接受函数进行处理。让我们看看怎么做：
正方形结构将实现如下接受方法：
func (obj *squre) accept(v visitor){
    v.visitForSquare(obj)
}
同样，圆和三角形也将定义上述接受函数。
现在，请求getArea()行为的团队可以简单地定义访问者界面的具体实现，并在该具体实现中编写区域计算逻辑。
*/

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForRectangle(*rectangle)
}

type shape interface {
	getType() string
	accept(visitor)
}

type square struct {
	side int
}

func (s *square) getType() string {
	return "square"
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

type circle struct {
	radius int
}

func (c *circle) getType() string {
	return "circle"
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

type rectangle struct {
	l int
	b int
}

func (r *rectangle) getType() string {
	return "rectangle"
}

func (r *rectangle) accept(v visitor) {
	v.visitForRectangle(r)
}

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("calculator area for square")
}

func (a *areaCalculator) visitForCircle(c *circle) {
	fmt.Println("calculator area for circle")
}

func (a *areaCalculator) visitForRectangle(r *rectangle) {
	fmt.Println("calculator area for rectangle")
}

type middleCoordinates struct {
	x int
	y int
}

func (m *middleCoordinates) visitForSquare(s *square) {
	fmt.Println("calculator middle point coordinates for square")
}

func (m *middleCoordinates) visitForCircle(c *circle) {
	fmt.Println("calculator middle point coordinates for circle")

}

func (m *middleCoordinates) visitForRectangle(r *rectangle) {
	fmt.Println("calculator middle point coordinates for rectangle")

}

func main() {
	s := &square{side: 2}
	c := &circle{radius: 3}
	r := &rectangle{l: 2, b: 3}

	ac := &areaCalculator{}
	s.accept(ac)
	c.accept(ac)
	r.accept(ac)

	fmt.Println()
	mc := &middleCoordinates{}
	s.accept(mc)
	c.accept(mc)
	r.accept(mc)
}

/* Output:
calculator area for square
calculator area for circle
calculator area for rectangle

calculator middle point coordinates for square
calculator middle point coordinates for circle
calculator middle point coordinates for rectangle
*/
