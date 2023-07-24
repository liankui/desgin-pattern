package main

import (
	"fmt"
	"sync"
)

// https://golangbyexample.com/queue-in-golang/

type customQueue struct {
	queue []string
	lock  sync.RWMutex
}

func (c *customQueue) Enqueue(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.queue = append(c.queue, name)
}

func (c *customQueue) Dequeue() error {
	if len(c.queue) > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.queue = c.queue[1:]
		return nil
	}
	return fmt.Errorf("pop error: queue is empty")
}

// Front returns the first element of list l or nil if the list is empty.
func (c *customQueue) Front() (string, error) {
	if len(c.queue) > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.queue[0], nil
	}
	return "", fmt.Errorf("peek error: queue is empty")
}

func (c *customQueue) Size() int {
	return len(c.queue)
}

func (c *customQueue) Empty() bool {
	return len(c.queue) == 0
}

func main() {
	cq := &customQueue{queue: make([]string, 0)}
	cq.Enqueue("A")
	cq.Enqueue("B")
	cq.Enqueue("C")
	fmt.Printf("size: %d\n", cq.Size())
	for cq.Size() > 0 {
		front, _ := cq.Front()
		fmt.Printf("front: %s, dequeue: %s\n", front, front)
		cq.Dequeue()
		fmt.Printf("size: %d\n", cq.Size())
	}
}

/*
size: 3
front: A, dequeue: A
size: 2
front: B, dequeue: B
size: 1
front: C, dequeue: C
size: 0
*/
