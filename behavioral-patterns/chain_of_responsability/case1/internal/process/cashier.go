package process

import (
	"fmt"

	"mercadolibre.com/pattern/cor/internal/bussines/model"
)

type Cashier struct {
	Next model.Department
}

func (c *Cashier) Execute(p *model.Patient) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
	if c.Next!= nil {
		c.Next.Execute(p)
	}
}

func (c *Cashier) SetNext(next model.Department) {
	c.Next = next
}
