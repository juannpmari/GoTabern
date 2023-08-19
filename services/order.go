// este servicio es para la creación de órdenes de la taberna
package services

import (
	"context"
	"ddd-go/aggregate"
	"ddd-go/domain/customer"
	"ddd-go/domain/customer/memory"
	"ddd-go/domain/customer/mongo"
	"ddd-go/domain/product"
	prodmem "ddd-go/domain/product/memory"
	"log"

	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository //We need a customer, so we neeed a CustomerRepository
	products product.ProductRepository

	// billing billing.Service -> por ahora no
}

// NewOrderService(WithCustomerRepository,WithMemoryProductRepository) -> es un ejemplo

//Factory function for our service
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error){
	// receives a variable amount of OrderConfiguration and returns a service with the cfg applied

	os := &OrderService{}
	// Loop through all the cfgs and apply them
	for _, cfg := range cfgs{
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// NewOrderService(WithMemoryCustomerRepository(),WithLogging("debug"),WithTracing("")) //Es un ejemplo
// En un futuro puedo cambiar WithMemoryCustomerRepository por WithMongoCustomerRepository, o para sql, etc.

// WitchCustomerRepository applies a customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration{
	// Return a function that matches the orderconfiguration alias
	return func(os *OrderService) error{
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration{
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(ctx context.Context, connStr string) OrderConfiguration{
	return func(os *OrderService) error{
		cr, err := mongo.New(ctx, connStr)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
	
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration{
	return func(os *OrderService) error {
		pr := prodmem.New()

		for _,p := range products{
			if err := pr.Add(p); err != nil {
				return err
			}
		}

		os.products = pr
		return nil 
	}
}

//Create a function for the OrderService service
func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (float64 , error) {
	// Fetch the customer
	c, err := o.customers.Get(customerID)
	if err != nil{
		return 0,err
	}

	// Apart from the curstomer, we need to get each Product
	var products []aggregate.Product
	var total float64

	for _,id := range productsIDs{
		p,err := o.products.GetByID(id)
		
		if err != nil {
			return 0,err
		}

		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("Customer: %s has ordered %d products",c.GetID(), len(products))
	return total,nil
}