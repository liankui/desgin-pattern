package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

/* https://golangbyexample.com/golang-object-pool/
对象池模式（Object Pool Pattern）是一种用于管理和重用一组预先分配的对象的设计模式。通过对象池，程序可以避免频繁地创建和销毁对象，从而提升性能，尤其是在对象创建和销毁开销较大的情况下。

使用场景
1.	数据库连接池：管理数据库连接，减少连接创建和销毁的开销。
2.	线程池：管理线程，避免频繁创建和销毁线程带来的开销。
3.	缓存对象：比如大对象或者频繁使用的对象，避免每次都重新创建。
4.	网络连接池：如 HTTP 客户端连接池，提升网络请求性能。

*/

type iPoolObject interface {
	getID() string // 这是可用于比较两个不同池对象的任何id
}

type connection struct {
	id string
}

func (c *connection) getID() string {
	return c.id
}

type pool struct {
	idle     []iPoolObject
	active   []iPoolObject
	capacity int
	mu       *sync.Mutex
}

// InitPool Initialize the pool
func initPool(poolObjects []iPoolObject) (*pool, error) {
	if len(poolObjects) == 0 {
		return nil, fmt.Errorf("cannot craete a pool of 0 length")
	}
	active := make([]iPoolObject, 0)
	pool := &pool{
		idle:     poolObjects,
		active:   active,
		capacity: len(poolObjects),
		mu:       new(sync.Mutex),
	}
	return pool, nil
}

// loan 取用资源持里的一个资源（切片的截取[1:]左闭右开区间）
func (p *pool) loan() (iPoolObject, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.idle) == 0 {
		return nil, fmt.Errorf("no pool object free, please request after sometime")
	}

	obj := p.idle[0]
	p.idle = p.idle[1:]
	p.active = append(p.active, obj)

	fmt.Printf("Loan Pool Object with ID: %s\n", obj.getID())
	return obj, nil
}

// receive 接收一个资源至资源池中
func (p *pool) receive(target iPoolObject) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if err := p.remove(target); err != nil {
		return err
	}

	p.idle = append(p.idle, target)

	fmt.Printf("Return Pool Object with ID: %s\n", target.getID())
	return nil
}

// remove pop出active队列中的target资源
func (p *pool) remove(target iPoolObject) error {
	currentActiveLength := len(p.active)
	for i, obj := range p.active {
		if obj.getID() == target.getID() {
			p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
			p.active = p.active[:currentActiveLength-1]
			return nil
		}
	}
	return fmt.Errorf("targe pool object doesn't belong to the pool")
}

func main() {
	connections := make([]iPoolObject, 0)
	for i := 0; i < 3; i++ {
		c := &connection{id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	pool, err := initPool(connections)
	if err != nil {
		log.Fatalf("Init Pool Error: %s", err)
	}
	conn1, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	conn2, err := pool.loan()
	if err != nil {
		log.Fatalf("Pool Loan Error: %s", err)
	}
	pool.receive(conn1)
	pool.receive(conn2)
}

/*
Loan Pool Object with ID: 0
Loan Pool Object with ID: 1
Return Pool Object with ID: 0
Return Pool Object with ID: 1
*/
