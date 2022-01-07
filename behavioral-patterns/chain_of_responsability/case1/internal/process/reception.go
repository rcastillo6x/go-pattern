package process

import (
	"fmt"
	"mercadolibre.com/pattern/cor/internal/bussines/model"
)

type Reception struct {
	Next model.Department
}

func (r *Reception) Execute(p *model.Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.Next.Execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	if r.Next!= nil {
		r.Next.Execute(p)
	}
}

func (r *Reception) SetNext(next model.Department) {
	r.Next = next
}
