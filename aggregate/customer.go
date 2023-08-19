// package aggregate holds out aggregates that combine many entities into a full object

package aggregate

import (
	"ddd-go/entity"
	"ddd-go/valueobject"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("A customer has to have a valid name")
)

type Customer struct {
	// person is the root entity of the customer, which means person.ID is the main identifier for the customer
	person *entity.Person
	products []*entity.Item
	transactions []valueobject.Transaction

}
// Notar lo siguiente:
// los campos están en minúscula -> no son accesibles desde afuera
// no se usan JSON, xq los aggregates no deciden cómo se formatea la data
// los entities usan punteros porque pueden cambiar de estado, y ese cambio se refleja en todos lados . Como la transaction no puede cambiar, no tiene sentido usar un puntero

//NewCustomer is a factory to create a new customer aggregate
//It will validate that the name is not empty
func NewCustomer(name string) (Customer,error){
	if name ==""{
		return Customer{}, ErrInvalidPerson 
	}
	
	person := &entity.Person{
		Name: name,
		ID: uuid.New(),
	}

	return Customer{
		person: person,
		products: make([]*entity.Item,0),
		transactions: make([]valueobject.Transaction,0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID{
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID){
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetName(name string){
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

