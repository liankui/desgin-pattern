package main

import (
	"fmt"
	"sync"
)

// https://golangbyexample.com/stack-in-golang/

type customStack struct {
	stack []string
	lock  sync.RWMutex
}

func (c *customStack) Push(value string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.stack = append(c.stack, value)
}

func (c *customStack) Pop() error {
	if len(c.stack) > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.stack = c.stack[:len(c.stack)-1]
	}
	return fmt.Errorf("pop error: Stack is empty")
}

// Front 相当于查看栈顶元素（Peek）
func (c *customStack) Front() (string, error) {
	if len(c.stack) > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.stack[len(c.stack)-1], nil
	}
	return "", fmt.Errorf("peep Error: Stack is empty")
}

func (c *customStack) Size() int {
	return len(c.stack)
}

func (c *customStack) Empty() bool {
	return len(c.stack) == 0
}

func main() {
	cs := &customStack{stack: []string{}}
	fmt.Printf("push: %s\n", "A")
	cs.Push("A")
	fmt.Printf("push: %s\n", "B")
	cs.Push("B")
	fmt.Printf("size: %d\n", cs.Size())
	for cs.Size() > 0 {
		front, _ := cs.Front()
		fmt.Printf("front: %s and pop: %s\n", front, front)
		cs.Pop()
		fmt.Printf("size: %d\n", cs.Size())
	}
}

/* Output:
push: A
push: B
size: 2
front: B and pop: B
size: 1
front: A and pop: A
size: 0
*/
