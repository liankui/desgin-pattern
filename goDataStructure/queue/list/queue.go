package main

import (
	"container/list"
	"fmt"
)

// https://golangbyexample.com/queue-in-golang/

type customQueue struct {
	queue *list.List
}

func (c *customQueue) Enqueue(value string) {
	c.queue.PushBack(value)
}

func (c *customQueue) Dequeue() error {
	if c.queue.Len() > 0 {
		c.queue.Remove(c.queue.Front())
	}
	return fmt.Errorf("pop error: queue is empty")
}

func (c *customQueue) Front() (string, error) {
	if c.queue.Len() > 0 {
		if val, ok := c.queue.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("peek error: queue datatype is incorrect")
	}
	return "", fmt.Errorf("peek error: queue is empty")
}

func (c *customQueue) Size() int {
	return c.queue.Len()
}

func (c *customQueue) Empty() bool {
	return c.queue.Len() == 0
}

func main() {
	cq := &customQueue{queue: list.New()}
	cq.Enqueue("A")
	cq.Enqueue("B")
	fmt.Printf("size: %d\n", cq.Size())
	for cq.Size() > 0 {
		front, _ := cq.Front()
		fmt.Printf("front: %s, dequeue: %s\n", front, front)
		cq.Dequeue()
		fmt.Printf("size: %d\n", cq.Size())
	}
}

/*
size: 2
front: A, dequeue: A
size: 1
front: B, dequeue: B
size: 0
*/
