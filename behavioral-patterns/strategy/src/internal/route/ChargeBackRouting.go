package route

import (
	"fmt"
	"mercadolibre.com/pattern/strategy/internal/bussines/model"
)

type ChargebackRouting struct {
}


func (l *ChargebackRouting) Routing(i model.Intention) {
	fmt.Println("RoutingType:  " + i.Type)
}