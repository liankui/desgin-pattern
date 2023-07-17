package main

import (
	"fmt"
	"log"
)

/* https://golangbyexample.com/facade-design-pattern-in-golang/
立面图案被归类为结构设计图案。这种设计模式旨在隐藏底层系统的复杂性，并为客户提供一个简单的界面。
它为系统中底层的许多接口提供了一个统一的接口，因此从客户端的角度来看，它更容易使用。基本上，它比复杂的系统提供了更高层次的抽象。

立面一词本身意味着“建筑物的主要正面，面向街道或开放空间“，只有建筑物的正面显示所有潜在的复杂性都隐藏在后面。
让我们用一个简单的例子来了解立面设计模式。在这个数字钱包的时代，当有人实际进行钱包借记/贷记时，背景中发生了许多客户
可能没有意识到的事情。以下列表说明了信用/借记过程中发生的一些活动
	支票账户
	检查安全别针
	贷记/借记余额
	分类账条目
	发送通知
值得注意的是，单个借记/贷记操作会发生很多事情。这就是立面图案进入画面的地方。作为客户端，只需输入钱包号码、安全密码、
金额并指定操作类型。其余的事情都在后台处理。在这里，我们创建了一个WalletFacade，它为客户端提供了一个简单的界面，并负责处理所有底层操作。
*/

type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
	ledger       *ledger
}

func newWalletFacade(accountId string, code int) *walletFacade {
	fmt.Println("starting create account")
	wf := &walletFacade{
		account:      newAccount(accountId),
		wallet:       newWallet(),
		securityCode: newSecurityCode(code),
		notification: &notification{},
		ledger:       &ledger{},
	}

	fmt.Println("account created")
	return wf
}

func (w *walletFacade) addMoneyToWallet(accountId string, securityCode, amount int) error {
	fmt.Println("starting add money to wallet")
	err := w.account.checkAccount(accountId)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountId, "credit", amount)

	return nil
}

func (w *walletFacade) deductMoneyFromWallet(accountId string, securityCode, amount int) error {
	fmt.Println("starting debit money from wallet")
	err := w.account.checkAccount(accountId)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}

	w.wallet.debitBalance(amount)
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountId, "credit", amount)

	return nil
}

type account struct {
	name string
}

func newAccount(accountName string) *account {
	return &account{name: accountName}
}

func (a *account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("account name is incorrect")
	}

	fmt.Println("account verified")

	return nil
}

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{code: code}
}

func (s *securityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("security code is incorrect")
	}

	fmt.Println("securityCode verified")

	return nil
}

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{balance: 0}
}

func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("wallet balance added successfully")
}

func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("balance is not sufficient")
	}

	fmt.Println("wallet balance is sufficient")

	return nil
}

type ledger struct {
}

func (l *ledger) makeEntry(accountId, txnType string, amount int) {
	fmt.Printf("make ledger entry for accountId %q with txnType %q for amount %d", accountId, txnType, amount)
}

type notification struct {
}

func (n *notification) sendWalletCreditNotification() {
	fmt.Println("sending wallet credit notification")
}

func (n *notification) sendWalletDebitNotification() {
	fmt.Println("sending wallet debit notification")
}

func main() {
	wf := newWalletFacade("aaa", 111)
	err := wf.addMoneyToWallet("aaa", 111, 100)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	err = wf.deductMoneyFromWallet("ab", 111, 5)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}
}

/* Output:
starting create account
account created
starting add money to wallet
account verified
securityCode verified
wallet balance added successfully
sending wallet credit notification
make ledger entry for accountId "aaa" with txnType "credit" for amount 100starting debit money from wallet
2023/07/17 15:09:36 error: account name is incorrect
*/
