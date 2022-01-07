package process

import (
	"fmt"
	"mercadolibre.com/pattern/cor/internal/bussines/model"
)

type Doctor struct {
	Next model.Department
}

func (d *Doctor) Execute(p *model.Patient) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.Next.Execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.DoctorCheckUpDone = true
	if d.Next!= nil {
		d.Next.Execute(p)
	}
}

func (d *Doctor) SetNext(next model.Department) {
	d.Next = next
}
