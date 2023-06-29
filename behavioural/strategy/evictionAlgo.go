package main

import "fmt"

/* https://golangbyexample.com/strategy-design-pattern-golang/
策略设计模式是一种行为设计模式。此设计模式允许您在运行时更改对象的行为，而无需更改该对象的类。

让我们用一个例子来了解策略模式。假设您正在构建内存缓存。由于它是内存缓存，所以大小有限。每当它达到最大大小时，
就需要从缓存中驱逐一些旧条目。这种驱逐可以通过几种算法发生。一些流行的算法是
	LRU - 最少最近使用：删除最近最少使用的条目。
	FIFO - 先入先出：删除先创建的条目。
	LFU - 最少使用：删除最不经常使用的条目。
现在的问题是如何将我们的缓存类与算法解耦，以便我们应该能够在运行时更改算法。此外，在添加新算法时，缓存类不应更改。
这就是策略模式的出现。策略模式建议创建一个算法家族，每个算法都有自己的类。这些类中的每一个都遵循相同的接口，这使得
算法在家族中可以互换。假设常见的接口名称是 evictionAlgo。

现在，我们的主缓存类将嵌入evictionAlgo接口。我们的Cache类不会自行实现所有类型的驱逐算法，而是将所有方法委托给
evictionAlgo接口。由于evictionAlgo是一个接口，我们可以运行时间将算法更改为LRU、FIFO、LFU，而无需在缓存类中进行任何更改。

何时使用
	当对象需要支持不同的行为，并且您想在运行时更改行为时。
	当您想避免选择运行时行为的很多条件时。
	当您有不同的算法相似时，它们只是在执行某些行为的方式上有所不同。
*/

type evictionAlgo interface {
	evict(c *cache)
}

type lru struct{}

func (l *lru) evict(c *cache) {
	fmt.Println("evicting by lru strategy")
}

type fifo struct{}

func (l *fifo) evict(c *cache) {
	fmt.Println("evicting by fifo strategy")
}

type lfu struct{}

func (l *lfu) evict(c *cache) {
	fmt.Println("evicting by lfu strategy")
}

type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	cap          int
	maxCap       int
}

func initCache(e evictionAlgo) *cache {
	return &cache{
		storage:      make(map[string]string),
		evictionAlgo: e,
		cap:          0,
		maxCap:       2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(k, v string) {
	if c.cap == c.maxCap {
		c.evict()
	}
	c.cap++
	c.storage[k] = v
}

func (c *cache) del(k string) {
	delete(c.storage, k)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.cap--
}

func main() {
	lfu := &lfu{}
	c := initCache(lfu)
	c.add("a", "1")
	c.add("b", "2")
	c.add("c", "3")
	lru := &lru{}
	c.setEvictionAlgo(lru)
	c.add("d", "4")
	fifo := &fifo{}
	c.setEvictionAlgo(fifo)
	c.add("e", "5")
}
