package iterator

/* https://golangbyexample.com/go-iterator-design-pattern/
迭代器设计模式是一种行为设计模式。在这种模式中，集合结构提供了一个迭代器，它允许它依次浏览集合结构中的每个元素，而不会暴露其底层实现。

以下是迭代器设计模式的基本组成部分
迭代器接口: 该接口提供了hasNext() 、getNext() 等基本操作。顾名思义，这些操作可让您遍历集合，重新启动迭代等
集合接口: 此接口表示需要遍历的集合。此接口定义了一个方法createIterator()，该方法返回迭代器类型
具体迭代器: 迭代器接口的具体实现
具体集合: 集合接口的具体实现

这种模式背后的主要思想是将集合结构的迭代逻辑公开到不同的对象 (实现迭代器接口) 中。此迭代器提供了一种独立于集合类型的迭代集合的通用方法。
*/

type iterator interface {
    hasNext() bool
    getNext() *user
}
