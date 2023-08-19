package aggregate

import (
	"ddd-go/entity"
	"errors"

	"github.com/google/uuid"
)

var (ErrMissingValues = errors.New("missing important values") )

type Product struct {
	item *entity.Item //root entity for Product
	price float64
	quantity int
}

//factory for Product
func NewProduct(name, description string, price float64) (Product, error){
	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &entity.Item{
			ID: uuid.New(),
			Name: name,
			Description: description,
		},
		price: price,
		quantity: 0,
	},nil
}

//other functions for Product
func (p Product) GetId() uuid.UUID{
	return p.item.ID
}

func (p Product) GetItem() *entity.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}