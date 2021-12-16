package route

import (
	"fmt"

	"mercadolibre.com/pattern/strategy/internal/bussines/model"
)

const (
	Purchase_routing   = "Purchase_routing"
	Chargeback_routing = "Chargeback_routing"
)

type (
	RouterStrategy interface {
		Routing(i model.Intention)
	}
)

func NewRouter(s string) (RouterStrategy, error) {
	switch s {
	case Purchase_routing:
		return &PurchaseRouting{}, nil
	case Chargeback_routing:
		return &ChargebackRouting{}, nil
	default:
		return nil, fmt.Errorf("Strategy '%s' not found", s)
	}
}
