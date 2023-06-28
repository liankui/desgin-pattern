package main

import (
	"fmt"
	"log"
)

/* https://golangbyexample.com/state-design-pattern-go/
状态设计模式是基于有限状态机的行为设计模式。我们将在自动售货机示例中解释国家设计模式。为了简单起见，
让我们假设自动售货机只有一种物品或产品。此外，为了简单起见，让我们假设自动售货机可以处于4种不同的状态
	hasItem
	noItem
	itemRequested
	hasMoney
自动售货机也会有不同的动作。同样为了简单起见，让我们假设只有四个操作：
	Select the item
	Add the item
	Insert Money
	Dispense Item
何时使用
	当对象可以处于许多不同的状态时，请使用状态设计模式。根据当前请求，对象需要更改其当前状态
 	  	在上面的例子中，自动售货机可以处于许多不同的状态。自动售货机将从一个状态移动到另一个状态。
		假设自动售货机在itemRequested中，那么一旦“插入钱”操作完成，它将移动到hasMoney状态
	当对象根据当前状态对同一请求有不同的响应时使用。在这里使用状态设计模式将防止许多条件语句
	  	例如，在自动售货机的情况下，如果用户想要购买物品，那么机器如果是hasItemState，则会继续，如果是innoItemState，
		则会拒绝。如果您在这里注意到，自动售货机在购买物品的请求上会给出两种不同的响应，这取决于它是在hasItemState还是noItemState中。
		请注意下面的 vendingMachine.go 文件，它没有任何条件语句。所有逻辑都由具体的状态实现来处理。
*/

type state interface {
	getStateName() string
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

type vendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state

	currentState state

	itemCount int
	itemPrice int
}

func newVendingMachine(itemCount int, itemPrice int) *vendingMachine {
	v := &vendingMachine{itemCount: itemCount, itemPrice: itemPrice}
	hasItemState := &hasItemState{vendingMachine: v}
	itemRequestedState := &itemRequestedState{vendingMachine: v}
	hasMoneyState := &hasMoneyState{vendingMachine: v}
	noItemState := &noItemState{vendingMachine: v}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) setState(s state) {
	stateName := ""
	if v.currentState != nil {
		stateName = v.currentState.getStateName()
	}
	fmt.Printf("\tset state, old state:%q, new state:%q\n", stateName, s.getStateName())
	v.currentState = s
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("adding %d items\n", count)
	v.itemCount += count
}

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (i *hasItemState) getStateName() string {
	if i != nil {
		return "hasItemState"
	}
	return ""
}

func (i *hasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *hasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("no item present")
	}
	fmt.Printf("item requested\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *hasItemState) insertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (i *hasItemState) dispenseItem() error {
	return fmt.Errorf("please select item first")
}

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) getStateName() string {
	if i != nil {
		return "itemRequestedState"
	}
	return ""
}

func (i *itemRequestedState) addItem(count int) error {
	return fmt.Errorf("item Dispense in progress")
}

func (i *itemRequestedState) requestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *itemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less. Please intert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *itemRequestedState) dispenseItem() error {
	return fmt.Errorf("please insert money first")
}

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (i *hasMoneyState) getStateName() string {
	if i != nil {
		return "hasMoneyState"
	}
	return ""
}

func (i *hasMoneyState) addItem(count int) error {
	return fmt.Errorf("item Dispense in progress")
}

func (i *hasMoneyState) requestItem() error {
	return fmt.Errorf("item Dispense in progress")
}

func (i *hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *hasMoneyState) dispenseItem() error {
	fmt.Println("dispensing item")
	i.vendingMachine.itemCount -= 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else if i.vendingMachine.itemCount > 0 {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}

type noItemState struct {
	vendingMachine *vendingMachine
}

func (i *noItemState) getStateName() string {
	if i != nil {
		return "noItemState"
	}
	return ""
}

func (i *noItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *noItemState) requestItem() error {
	return fmt.Errorf("item out of stock")
}

func (i *noItemState) insertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *noItemState) dispenseItem() error {
	return fmt.Errorf("item out of stock")
}

func main() {
	vm := newVendingMachine(1, 10)
	if err := vm.requestItem(); err != nil {
		log.Fatal(err)
	}
	if err := vm.insertMoney(10); err != nil {
		log.Fatal(err)
	}
	if err := vm.dispenseItem(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----")
	if err := vm.addItem(2); err != nil {
		log.Fatal(err)
	}

	fmt.Println("-----")
	if err := vm.requestItem(); err != nil {
		log.Fatal(err)
	}
	if err := vm.insertMoney(10); err != nil {
		log.Fatal(err)
	}
	if err := vm.dispenseItem(); err != nil {
		log.Fatal(err)
	}
}

/* Output:
        set state, old state:"", new state:"hasItemState"
item requested
        set state, old state:"hasItemState", new state:"itemRequestedState"
money entered is ok
        set state, old state:"itemRequestedState", new state:"hasMoneyState"
dispensing item
        set state, old state:"hasMoneyState", new state:"noItemState"
-----
adding 2 items
        set state, old state:"noItemState", new state:"hasItemState"
-----
item requested
        set state, old state:"hasItemState", new state:"itemRequestedState"
money entered is ok
        set state, old state:"itemRequestedState", new state:"hasMoneyState"
dispensing item
        set state, old state:"hasMoneyState", new state:"hasItemState"
*/
