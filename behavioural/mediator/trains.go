package main

import (
	"fmt"
	"sync"
)

/* https://golangbyexample.com/mediator-design-pattern-golang/
调解员设计模式是一种行为设计模式。这种模式建议创建一个中介对象，以防止对象之间的直接通信，从而避免它们之间的直接依赖。

调解人模式的一个很好的例子是铁路系统平台。两列火车为了站台的可用性而从不相互通信。车站经理充当调解人，只为其中一列火车提供站台。
火车与车站经理连接，并采取相应行动。它保持着排队等候的火车。如果任何火车离开站台，它会通知其中一列火车接下来到达站台。

在下面的代码中，注意stationManger如何充当火车和站台之间的调解人。
	乘客火车和货物火车实现了火车接口。
	stationManger实现了调解员接口。
*/

type train interface {
	requestArrival()
	departure()
	permitArrival()
}

type passengerTrain struct {
	mediator
}

func (p *passengerTrain) requestArrival() {
	if p.canLand(p) {
		fmt.Println("passengerTrain landing")
	} else {
		fmt.Println("passengerTrain waiting")
	}
}

func (p *passengerTrain) departure() {
	fmt.Println("passengerTrain leaving")
	p.notifyFree()
}

func (p *passengerTrain) permitArrival() {
	fmt.Println("passengerTrain arrival permitted. Landing...")
}

type goodsTrain struct {
	mediator
}

func (g *goodsTrain) requestArrival() {
	if g.canLand(g) {
		fmt.Println("goodsTrain landing")
	} else {
		fmt.Println("goodsTrain waiting")
	}
}

func (g *goodsTrain) departure() {
	fmt.Println("goodsTrain leaving")
	g.notifyFree()
}

func (g *goodsTrain) permitArrival() {
	fmt.Println("goodsTrain arrival permitted. Landing...")
}

type mediator interface {
	canLand(train) bool
	notifyFree()
}

type stationManager struct {
	isPlatformFree    bool
	lock              *sync.Mutex
	trainWaitingQueue []train
}

func NewStationManager() *stationManager {
	return &stationManager{
		isPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}

func (s *stationManager) canLand(t train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainWaitingQueue = append(s.trainWaitingQueue, t)
	return false
}

func (s *stationManager) notifyFree() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainWaitingQueue) > 0 {
		firstTrainInQueue := s.trainWaitingQueue[0]
		s.trainWaitingQueue = s.trainWaitingQueue[1:]
		firstTrainInQueue.permitArrival()
	}
}

func main() {
	sm := NewStationManager()
	pt := &passengerTrain{sm}
	gt := &goodsTrain{sm}
	pt.requestArrival()
	gt.requestArrival()
	pt.departure()
}

/* Output:
passengerTrain landing
goodsTrain waiting
passengerTrain leaving
goodsTrain arrival permitted. Landing...
*/
