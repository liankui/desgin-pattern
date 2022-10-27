package chainOfResponsibility

import "testing"

// reception -> doctor -> medical -> cashier
func TestChainOfResponsibility(t *testing.T) {
	cashier := &cashier{}
	//Set next for medical department
	medical := &medical{}
	medical.setNext(cashier)
	//Set next for doctor department
	doctor := &doctor{}
	doctor.setNext(medical)
	//Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)
	patient := &patient{name: "abc"}
	//Patient visiting
	reception.execute(patient)
}

/*
Reception registering patient
Doctor checking patient
Medical giving medicine to patient
Cashier getting money from patient patient
*/
