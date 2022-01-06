package main

import (
	"fmt"

	"mercadolibre.com/pattern/strategy/internal/bussines/model"
	mgnt "mercadolibre.com/pattern/strategy/internal/management"
)

func main() {
	//  Given an intention Purchase
	i := model.Intention{
		Type:    "Purchase",
		Payload: "{}",
	}
	// Define the strategy
	r, e := mgnt.NewManagement(mgnt.PurchaseIntention)
	if e != nil {
		fmt.Println(e.Error())
	}
	// Process Purchase
	fmt.Println("*** Process Purchase ***")
	r.Validate(i)
	r.Processing(i)
	r.Routing(i)

	//  Given an intention ChargeBack
	i = model.Intention{
		Type:    "ChargeBack",
		Payload: "{}",
	}

	// Define the strategy
	r, e = mgnt.NewManagement(mgnt.ChargebackIntention)
	if e != nil {
		fmt.Println(e.Error())
	}

	// Process Chargeback
	fmt.Println("*** Process ChargeBack ***")
	r.Validate(i)
	r.Processing(i)
	r.Routing(i)

}
