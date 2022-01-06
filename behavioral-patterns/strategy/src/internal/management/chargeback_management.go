package intention_mgnt

import (
	"fmt"

	"mercadolibre.com/pattern/strategy/internal/bussines/model"
)

type ChargebackManagement struct {
}

func (l *ChargebackManagement) Routing(i model.Intention) {
	fmt.Printf("Validate Intention Type: %v \n ", i.Type)
}

func (l *ChargebackManagement) Validate(i model.Intention) {
	fmt.Printf("Validate Intention Type: %v \n ", i.Type)
}

func (l *ChargebackManagement) Processing(i model.Intention) {
	fmt.Printf("Processing Intention Type: %v \n ", i.Type)
}
