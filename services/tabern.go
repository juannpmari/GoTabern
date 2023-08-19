package services

import (
	"log"

	"github.com/google/uuid"
)

type TabernConfiguration func(os *Tabern) error //Is a function signature

// Tabern is a service that holds subservices
type Tabern struct {
	//orderservice to take orders
	OrderService *OrderService

	// BillingService -> to receive payments
	BillingService interface{}
}

// Factory

func NewTabern(cfgs ...TabernConfiguration) (*Tabern, error) {
	t := &Tabern{}

	for _, cfg := range cfgs {
		if err := cfg(t); err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithOrderService(os *OrderService) TabernConfiguration {
	return func(t *Tabern) error {
		t.OrderService = os
		return nil
	}
}

func (t *Tabern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}

	log.Printf("\nBill the customer: %0.0f\n", price)

	return nil

}