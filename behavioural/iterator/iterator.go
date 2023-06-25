package main

import (
	"fmt"
)

/* 迭代器设计模式 https://golangbyexample.com/go-iterator-design-pattern/
一种行为设计模式。在这种模式中，集合结构提供了一个迭代器，它允许它依次浏览集合结构中的每个元素，而不会暴露其底层实现。

基本组成部分
1.迭代器接口: 该接口提供了hasNext() 、getNext() 等基本操作。顾名思义，这些操作可让您遍历集合，重新启动迭代等
2.集合接口: 此接口表示需要遍历的集合。此接口定义了一个方法createIterator()，该方法返回迭代器类型
3.具体迭代器: 迭代器接口的具体实现
4.具体集合: 集合接口的具体实现
这种模式背后的主要思想是将集合结构的迭代逻辑公开到不同的对象 (实现迭代器接口) 中。此迭代器提供了一种独立于集合类型的迭代集合的通用方法。
*/
type user struct {
	name string
	age  int
}

type userIterator struct {
	index int
	users []*user
}

func (u *userIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *userIterator) getNext() *user {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}

type iterator interface {
	hasNext() bool
	getNext() *user
}

type collection interface {
	createIterator() iterator
}

type userCollection struct {
	users []*user
}

func (u *userCollection) createIterator() iterator {
	return &userIterator{
		users: u.users,
	}
}

func main() {
	user1 := &user{
		name: "a",
		age:  30,
	}
	user2 := &user{
		name: "b",
		age:  20,
	}
	userCollection := &userCollection{
		users: []*user{user1, user2},
	}
	iterator := userCollection.createIterator()
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}

/*
User is &{name:a age:30}
User is &{name:b age:20}
*/
