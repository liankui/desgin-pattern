package main

import "fmt"

/* https://golangbyexample.com/null-object-design-pattern-golang/
空对象设计模式是一种行为设计模式。当客户端代码依赖于一些可以为空的依赖项时，这很有用。使用此设计模式可防止客户端对这些依赖项
的结果进行空检查。话虽如此，还应该注意的是，对于这种空依赖项，客户端行为也很好。

空对象设计模式的主要组成部分是：
	实体-它是定义子结构必须实现的原始操作的接口
	ConcreteEntity - 它实现了实体接口
	NullEntity - 它表示空对象。它还实现了实体接口，但具有空属性
	客户端-客户端获得实体接口的实现并使用它。它并不真正关心实现是ConcreteEntity还是NullEntity。它对待他们两个都一样。

假设我们有一所学院，有很多系，每个系都有一定数量的教授。
*/

type department interface {
	getNumberOfProfessors() int
	getName() string
}

type college struct {
	departments []department
}

func (c *college) addDepartment(departmentName string, numOfProfessors int) {
	switch departmentName {
	case "computerScience":
		csDepartment := &computerScience{numberOfProfessors: numOfProfessors}
		c.departments = append(c.departments, csDepartment)
	case "mechanical":
		mechDepartment := &mechanical{numberOfProfessors: numOfProfessors}
		c.departments = append(c.departments, mechDepartment)
	}
}

func (c *college) getDepartment(departmentName string) department {
	for _, dep := range c.departments {
		if dep.getName() == departmentName {
			return dep
		}
	}
	// return a null department if the department doesn't exits
	return &nullDepartment{}
}

type computerScience struct {
	numberOfProfessors int
}

func (c *computerScience) getNumberOfProfessors() int {
	return c.numberOfProfessors
}

func (c *computerScience) getName() string {
	return "computerScience"
}

type mechanical struct {
	numberOfProfessors int
}

func (m *mechanical) getNumberOfProfessors() int {
	return m.numberOfProfessors
}

func (m *mechanical) getName() string {
	return "mechanical"
}

type nullDepartment struct {
	numberOfProfessors int
}

func (n *nullDepartment) getNumberOfProfessors() int {
	return 0
}

func (n *nullDepartment) getName() string {
	return "nullDepartment"
}

func main() {
	college1 := college{}
	college1.addDepartment("computerScience", 4)
	college1.addDepartment("mechanical", 5)

	college2 := college{}
	college2.addDepartment("computerScience", 2)

	departmentArray := []string{"computerScience", "mechanical", "civil", "electronics"}

	totalProfessors1 := 0
	for _, depName := range departmentArray {
		d := college1.getDepartment(depName)
		totalProfessors1 += d.getNumberOfProfessors()
	}
	fmt.Printf("total number of professors in college1 is %d\n", totalProfessors1)

	totalProfessors2 := 0
	for _, depName := range departmentArray {
		d := college2.getDepartment(depName)
		totalProfessors2 += d.getNumberOfProfessors()
	}
	fmt.Printf("total number of professors in college2 is %d\n", totalProfessors2)
}

/* Output:
total number of professors in college1 is 9
total number of professors in college2 is 2
*/
