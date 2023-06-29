package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

/*
	https://golangbyexample.com/golang-object-pool/

对象池设计模式是一种创造性的设计模式，其中预先初始化和创建对象池，并将其保存在池中。
根据需要，客户端可以从池中请求对象，使用它并将其返回到池中。池中的物体从未被摧毁。

何时使用:
1.当创建类对象的成本很高，并且在特定时间需要的此类对象的数量不多时。

	-以DB连接为例。每个连接对象的创建成本都很高，因为涉及网络调用，并且一次可能不需要超过某个连接。对象池设计模式非常适合这种情况。

2.当池对象是不可变对象时

	-以DB连接为例。DB连接是一个不可变的对象。几乎没有任何财产需要更改

3.出于性能原因。由于已经创建了池，因此它将显着提高应用程序的性能
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
