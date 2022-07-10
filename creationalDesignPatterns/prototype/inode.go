package prototype

/* https://golangbyexample.com/prototype-pattern-go/
原型模式
这是一种创造性的设计模式，可让您创建对象的副本。在此模式中，创建克隆对象的责任被委派给要克隆的实际对象。
要克隆的对象公开了一个克隆方法，该方法返回该对象的克隆副本。

何时使用：
- 当要克隆的对象的创建过程很复杂时，我们使用原型模式，即克隆可能涉及处理深副本，分层副本等的情况。此外，也可能有一些私人成员无法直接访问。
- 创建对象的副本，而不是从头开始创建新实例。这样可以防止在创建新对象 (例如数据库操作) 时涉及昂贵的操作。
- 当您想要创建一个新对象的副本时，但它仅作为接口提供给您。因此，您不能直接创建该对象的副本。

demo：
在golang的上下文中，让我们尝试通过os文件系统的示例来理解它。os文件系统有文件和文件夹，文件夹本身包含文件和文件夹。
每个文件和文件夹都可以用inode接口表示。inode接口也有clone() 函数。
*/

type inode interface {
	print(string)
	clone() inode
}