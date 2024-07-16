package main

import "fmt"

type Prototype interface {
	Clone() Prototype
	GetName() string
	SetName(name string)
}

type ConcretePrototype struct {
	Name string
}

func (p *ConcretePrototype) Clone() Prototype {
	clone := *p
	return &clone
}

func (p *ConcretePrototype) GetName() string {
	return p.Name
}

func (p *ConcretePrototype) SetName(name string) {
	p.Name = name
}

func main() {
	prototype := ConcretePrototype{Name: "prototype1"}

	clone := prototype.Clone()

	clone.SetName("clone1")

	fmt.Printf("original name:%s\n", prototype.GetName())
	fmt.Printf("clone name:%s\n", clone.GetName())
}

/* Output:
original name:prototype1
clone name:clone1
*/

/* UML
+-----------------------+
|       Prototype       |
+-----------------------+
| + Clone() : Prototype |
+-----------------------+
           ▲
           |
           |
+-----------------------------+
|     ConcretePrototype       |
+-----------------------------+
| - Name : string             |
+-----------------------------+
| + Clone() : Prototype       |
| + GetName() : string        |
| + SetName(name : string)    |
+-----------------------------+
*/
