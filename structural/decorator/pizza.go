package main

import "fmt"

/* https://golangbyexample.com/decorator-pattern-golang/
装饰设计模式是一种结构设计模式。它允许您提供附加功能或装饰对象，而无需更改该对象。
用一个例子可以更好地理解它。想象一下，你正在开一家披萨连锁店。你从两种披萨开始
	素食狂热披萨
	佩皮豆腐披萨
上述每个披萨都有其价格。所以你会创建一个披萨界面，如下所示
type pizza interface {
	getPrice() int
}
您还需要创建两个带有getPrice函数的披萨结构，该函数将返回价格。这两个披萨结构在定义getPrice()方法时实现了披萨接口
后来，你开始提供配料和披萨，每个配料都有一些额外的价格。因此，原来的基础披萨现在需要用浇头装饰。想象一下，您在菜单中添加了以下两个配料
	番茄浇头
	奶酪浇头
此外，请记住，披萨和配料也是披萨。客户可以通过不同的方式选择他们的披萨。对于例如
	带有番茄浇头的素食狂热
	蔬菜主菜配番茄和奶酪浇头
	没有任何配料的Peppy Paneer披萨
	奶酪浇头的Peppy Paneer披萨
	...
那么，鉴于你现在也有配料，你现在会如何设计呢。装饰图案将出现在画面中。它可以帮助附加功能，而无需实际修改任何现有结构。
在这种情况下，装饰器模式建议为每个可用的顶部创建单独的结构。每个顶部结构都将实现上面的披萨界面，并具有披萨的嵌入和实例。
我们现在有不同类型的披萨的单独结构和可用配料类型的单独结构。每个披萨和配料都有自己的价格。每当您在披萨上添加任何配料时，
该配料的价格就会添加到基础披萨的价格中，这就是您获得最终价格的方式。
因此，装饰图案让您在不改变披萨对象的情况下装饰原始基础披萨对象。披萨对象对配料一无所知。它只是知道它的价格，没有别的。
*/

type pizza interface {
	getPrice() int
}

type peppyPaneer struct{}

func (pp *peppyPaneer) getPrice() int {
	return 20
}

type veggeMania struct{}

func (vm *veggeMania) getPrice() int {
	return 15
}

type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

type tomatoTopping struct {
	pizza pizza
}

func (t *tomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 7
}

func main() {
	peppyPaneerPizza := &peppyPaneer{}
	peppyPaneerPizzaWithCheese := &cheeseTopping{pizza: peppyPaneerPizza}
	peppyPaneerPizzaWithCheeseAndTomato := &tomatoTopping{pizza: peppyPaneerPizzaWithCheese}
	fmt.Printf("price of peppyPaneer Pizza With Cheese And Tomato is %d\n", peppyPaneerPizzaWithCheeseAndTomato.getPrice())

	veggeManiaPizza := &veggeMania{}
	veggeManiaPizzaWithTomatoTopping := &tomatoTopping{pizza: veggeManiaPizza}
	fmt.Printf("price of veggeMania Pizza With Tomato is %d\n", veggeManiaPizzaWithTomatoTopping.getPrice())
}

/* Output:
price of peppyPaneer Pizza With Cheese And Tomato is 37
price of veggeMania Pizza With Tomato is 22
*/
