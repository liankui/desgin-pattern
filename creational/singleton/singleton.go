package main

import (
	"fmt"
	"sync"
)

/* https://golangbyexample.com/singleton-design-pattern-go/
单例模式适用于：
	1. DB实例-我们只想创建DB对象的一个实例，并且该实例将在整个应用程序中使用。
*/

type single struct{}

// 全局实例
var singleInstance *single
var lock sync.Mutex

// 创建实例（双重校验nil）
func getInstance() *single {
	// 在开始时检查nil单实例，防止每次调用时昂贵的锁操作。
	if singleInstance == nil {
		// 单实例在锁内部创建
		lock.Lock()
		defer lock.Unlock()
		// 另一个nil singleInstance检查，为了确保多个goroutine绕过第一次检查，只有一个goroutine能够创建单实例。
		if singleInstance == nil {
			fmt.Println("Creating Single Instance Now")
			singleInstance = new(single)
		} else {
			fmt.Println("Single Instance already created-1")
		}
	} else {
		fmt.Println("Single Instance already created-2")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 100; i++ {
		go getInstance2()
	}

	fmt.Scanln() // Scanln类似于Scan，但是在换行符处停止扫描，并且在最终项目之后必须有一个换行符或EOF。
}

/*
有一个 “立即创建单个实例” 的输出，这意味着只有一个goroutine能够创建单个实例。
有几个 “单个实例已经创建-1” 的输出，这意味着一些goroutines在第一次检查中发现了单实例的值为nil，并绕过了该值。
有几个输出 “单个实例已经创建-2”，这意味着当他们到达单个实例时，单个实例已经创建，并且如果检查，他们无法绕过第一个
*/

// 重构，使用sync.Once也能实现只创建一次的结果。
var once sync.Once

func getInstance2() *single {
	once.Do(func() {
		fmt.Println("Creting Single Instance Now")
		singleInstance = new(single)
	})
	return singleInstance
}

/*
有一个 “立即创建单个实例” 的输出，这意味着只有一个goroutine能够创建单个实例
有几个输出 “单个实例已经创建-2”，这意味着当他们到达单个实例时，单个实例已经创建，并且如果检查，他们无法绕过第一个
*/

// 最佳实践1
type Engine interface {
	// Name get the name of the engine
	Name() string
}

var engines map[string]func() Engine
var enginesOnce sync.Once

func RegisterEngine(name string, engine func() Engine) {
	enginesOnce.Do(func() {
		engines = make(map[string]func() Engine)
	})
	engines[name] = engine
}
