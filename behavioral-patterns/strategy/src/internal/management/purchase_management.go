package intention_mgnt

import (
	"fmt"

	"mercadolibre.com/pattern/strategy/internal/bussines/model"
)

type PurchaseManagement struct {
}

func (l *PurchaseManagement) Routing(i model.Intention) {
	fmt.Printf("Validate Intention Type: %v \n ", i.Type)
}

func (l *PurchaseManagement) Validate(i model.Intention) {
	fmt.Printf("Validate Intention Type: %v \n ", i.Type)
}

func (l *PurchaseManagement) Processing(i model.Intention) {
	fmt.Printf("Processing Intention Type: %v \n ", i.Type)
}
