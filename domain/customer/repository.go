package customer

import (
	"ddd-go/aggregate"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer = errors.New("failed to update the customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
// repository for fetching customers and managing them
// es una abstracción, porque acá no importa si viene de memoria, mongo, sql, etc.

// Repositories are for managing aggregates
// IMPORTANTE: cada repositorio maneja un solo aggregate (loose-coupling)