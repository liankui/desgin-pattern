package main

import (
	"fmt"
	"sync"
)

type Object struct {
	id string
}

type ObjectPool struct {
	pool sync.Pool
}

func NewObjectPool() *ObjectPool {
	return &ObjectPool{
		pool: sync.Pool{
			New: func() any { return &Object{} },
		},
	}
}

func (p *ObjectPool) Get() *Object {
	return p.pool.Get().(*Object)
}

func (p *ObjectPool) Put(obj *Object) {
	p.pool.Put(obj)
}

func main() {
	pool := NewObjectPool()

	obj := pool.Get()
	fmt.Printf("get object:%p\n", obj)
	obj2 := pool.Get()
	fmt.Printf("get object again:%p\n", obj2)

	pool.Put(obj)
	pool.Put(obj2)

	obj3 := pool.Get()
	fmt.Printf("get object again:%p\n", obj3)
	obj4 := pool.Get()
	fmt.Printf("get object again:%p\n", obj4)
}
