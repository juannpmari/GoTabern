package product

import (
	"ddd-go/aggregate"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound = errors.New("No such product")
	ErrProductAlreadyExists = errors.New("there is already such a product")
)

type ProductRepository interface { //manages and handles Product aggregates
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
