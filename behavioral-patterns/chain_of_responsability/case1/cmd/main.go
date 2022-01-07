package main

import (
	"mercadolibre.com/pattern/cor/internal/bussines/model"
	"mercadolibre.com/pattern/cor/internal/process"
)

func createPipelineProcess() *process.Reception{
	medical := &process.Medical{}

	//Set next for cashier department
	cashier := &process.Cashier{}
	cashier.SetNext(medical)
	//Set next for doctor department
	doctor := &process.Doctor{}
	doctor.SetNext(cashier)
	//Set next for reception department
	reception := &process.Reception{}
	reception.SetNext(doctor)
	return  reception
}

func main() {

	// Given a Patient
	patient := &model.Patient{Name: "abc"}

	// Create pipeline of process
	reception:= createPipelineProcess()

	//Patient visiting
	reception.Execute(patient)
}
