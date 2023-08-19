package services

import (
	"ddd-go/aggregate"
	"testing"
	
	"github.com/google/uuid"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer","Helathy beverage",199)
	if err != nil {
		t.Fatal(err)
	}

	peanuts, err := aggregate.NewProduct("Peanuts","Snacks",0.99)

	if err != nil{
		t.Fatal(err)
	}

	wine, err := aggregate.NewProduct("Wine","nasty drink", 0.99)
	if err != nil{
		t.Fatal(err)
	}

	return []aggregate.Product{
		beer,peanuts,wine,
	}
}

func TestOrder_NewOrderService(t *testing.T){
	products := init_products(t)
	
	os,err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil{
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("Percy")
	if err != nil{
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetId(),
	}

	_,err = os.CreateOrder(cust.GetID(), order)

	if err != nil {
		t.Error(err)
	}
}