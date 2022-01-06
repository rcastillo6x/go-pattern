package intention_mgnt

import (
	"fmt"

	"mercadolibre.com/pattern/strategy/internal/bussines/model"
)

const (
	PurchaseIntention   = "purchase_intention"
	ChargebackIntention = "chargeback_intention"
)

type (
	IntentionStrategy interface {
		Validate(i model.Intention)
		Processing(i model.Intention)
		Routing(i model.Intention)
	}
)

func NewManagement(s string) (IntentionStrategy, error) {
	switch s {
	case PurchaseIntention:
		return &PurchaseManagement{}, nil
	case ChargebackIntention:
		return &ChargebackManagement{}, nil
	default:
		return nil, fmt.Errorf("strategy '%s' not found", s)
	}
}
