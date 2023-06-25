package main

import "fmt"

/* https://golangbyexample.com/chain-of-responsibility-design-pattern-in-golang/
责任链设计模式是一种行为设计模式。它允许您创建一条请求处理程序链。对于每个传入的请求，它通过链和每个处理程序传递:
	- 处理请求或跳过处理。
	- 决定是否将请求传递给链中的下一个处理程序
责任链设计模式将通过一个示例得到最好的理解。让我们以医院为例。医院有多个部门，例如:
	接待
	医生
	药室
	出纳员
每当有病人到达时，他首先去接待处，然后去医生，然后去手术室，然后去收银员，依此类推。在某种程度上，患者被送到一系列部门，当完成时，将患者送到其他部门。这就是责任链模式出现的地方。
何时使用:
* 当有多个候选者处理同一请求时，该模式适用。
* 当您不希望客户端选择接收器时，因为多个对象可以处理请求。另外，您希望将客户端与接收器解耦。客户端只需要知道链中的第一个元素。
	- 与医院的示例一样，患者首先进入接待处，然后根据患者的当前状态将其发送到链中的下一个处理程序。
*/

type department interface {
	execute(*patient)
	setNext(department)
}

type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next department) {
	r.next = next
}

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *cashier) setNext(next department) {
	c.next = next
}

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func main() {
	cashier := &cashier{}

	// Set next for medical department
	medical := &medical{}
	medical.setNext(cashier)

	// Set next for doctor department
	doctor := &doctor{}
	doctor.setNext(medical)

	// Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "abc"}
	// Patient visiting
	reception.execute(patient)
}

/* Output:
Reception registering patient
Doctor checking patient
Medical giving medicine to patient
Cashier getting money from patient patient
*/
