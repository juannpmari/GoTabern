// Package memory is a in-memory implementation of Customer repository
// Esta implementación es para crear el customer en memoria. Podría por ej. tener otra para crearlo en una mongo o una postgres
package memory

import (
	"ddd-go/aggregate"
	"ddd-go/domain/customer"
	"fmt"
	"sync"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository{ //factory function to create a MemoryRepository
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

// Other functions for MemoryRepository
func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer,error){
	//receives id and return Customer, error
	if customer, ok := mr.customers[id]; ok{
		return customer, nil
	}
	return aggregate.Customer{},customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(c aggregate.Customer) error{
	//Receives a customer and saves it into the MemoryRepository's map
	if mr.customers == nil { //In case there's no map
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}
	// Make sure customer is already in repo
	if _,ok := mr.customers[c.GetID()]; ok{
		return fmt.Errorf("customer already exists: %w",customer.ErrFailedToAddCustomer)
	}

	//if it doesn't exist yet, we add it
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil 
}

func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	if _,ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}

	//if it exists
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}


