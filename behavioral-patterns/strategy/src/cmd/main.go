package main

import (
	"fmt"
	"mercadolibre.com/pattern/strategy/internal/bussines/model"
	"mercadolibre.com/pattern/strategy/internal/route"
)

func main() {
	//  Given an intention Purchase
	i := model.Intention{
		Type: "Purchase",
		Payload:"{}",
	}
	// Define the strategy
	r, e := route.NewRouter(route.Purchase_routing)
	if e != nil{
		fmt.Println (e.Error())
	}
	r.Routing(i)

	//  Given an intention ChargeBack
	i = model.Intention{
		Type: "ChargeBack",
		Payload:"{}",
	}
	r, e = route.NewRouter(route.Chargeback_routing)
	if e != nil{
		fmt.Println (e.Error())
	}
	r.Routing(i)

}
