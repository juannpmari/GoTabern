package services

import (
	"ddd-go/aggregate"
	"testing"

	"github.com/google/uuid"
)

func Test_Tabern(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		// WithMongoCustomerRepository(context.Background(),"mongodb://localhost:27017"),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	tabern, err := NewTabern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("percy")
	if err != nil {
		t.Fatal(err)
	}

	if err = os.customers.Add(cust); err != nil{
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetId(),
	}

	err = tabern.Order(cust.GetID(),order)

	if err != nil{
		t.Fatal(err)
	}
}