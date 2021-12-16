package route

import (
	"fmt"
	"mercadolibre.com/pattern/strategy/internal/bussines/model"
)

type PurchaseRouting struct {
}

func (l *PurchaseRouting) Routing(i model.Intention) {
	fmt.Println("RoutingType:  " + i.Type)
}