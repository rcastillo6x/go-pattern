package process

import (
	"fmt"
	"mercadolibre.com/pattern/cor/internal/bussines/model"
)

type Medical struct {
	Next model.Department
}

func (m *Medical) Execute(p *model.Patient) {
	if p.MedicineDone {
		fmt.Println("Medicine already given to patient")
		m.Next.Execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")

    if p!= nil {
		p.MedicineDone = true
		if m.Next!= nil {
			m.Next.Execute(p)
		}
	}
}

func (m *Medical) SetNext(next model.Department) {
	m.Next = next
}
