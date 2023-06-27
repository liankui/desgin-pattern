package main

import "fmt"

/* https://golangbyexample.com/observer-design-pattern-golang/
观察者设计模式是一种行为设计模式。这种模式允许实例（称为主题）将事件发布到其他多个实例（称为观察者）。这些观察员订阅了该主题，因此，如果主题发生任何变化，则会收到事件通知。

让我们举个例子。在电子商务网站上，许多商品缺货。可能会有客户对缺货的特定商品感兴趣。这个问题有三种解决方案
	客户不断以某种频率检查物品的可用性。
	电子商务用库存的所有新商品轰炸客户
	客户只订阅他感兴趣的特定项目，并在该项目可用时收到通知。此外，多个客户可以订阅同一产品
选项3是最可行的，这就是Observer Patter的全部内容。观察者模式的主要组成部分是：
	主题-当任何变化时，它是发布事件的实例。
	观察者-它订阅主题并收到事件的通知。
一般来说，主题和观察者作为一个接口实现。两者都使用了具体实施
*/

type subject interface {
	register(o observer)
	deregister(o observer)
	notifyAll()
}

type item struct {
	observerList []observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{name: name}
}

func (i *item) updateAvailability() {
	fmt.Printf("item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o observer) {
	l := len(i.observerList)
	for j := l - 1; j >= 0; j-- {
		if i.observerList[j].getId() == o.getId() {
			i.observerList = append(i.observerList[:j], i.observerList[j+1:]...)
		}
	}
}

func (i *item) notifyAll() {
	for _, obs := range i.observerList {
		obs.update(i.name)
	}
}

type observer interface {
	update(string)
	getId() string
}

type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	fmt.Printf("sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *customer) getId() string {
	return c.id
}

func main() {
	iphoneItem := newItem("iphone")
	o1 := &customer{id: "111@gamil.com"}
	o2 := &customer{id: "222@gamil.com"}

	iphoneItem.register(o1)
	iphoneItem.register(o2)
	iphoneItem.updateAvailability()
	iphoneItem.deregister(o1)
	iphoneItem.deregister(o2)
	iphoneItem.updateAvailability()
}

/* Output:
item iphone is now in stock
sending email to customer 111@gamil.com for item iphone
sending email to customer 222@gamil.com for item iphone
item iphone is now in stock
*/
