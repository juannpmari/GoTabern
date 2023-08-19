package memory

import (
	"ddd-go/aggregate"
	"ddd-go/domain/product"
	"sync"

	"github.com/google/uuid"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

//Factory MemoryProductRepository
func New() *MemoryProductRepository{
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}

//other functions for MemoryProductRepository
func (mpr *MemoryProductRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product

	for _, product := range mpr.products {
		products = append(products, product)
	}
	return products, nil
	//It's the interface, and not the implementation, that decides if we need to return an error
	// in this case we don't but for ex with Mongo, we would need to return error (e.g. connection failed)
}

func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error){
	if product, ok := mpr.products[id]; ok{
		return product, nil
	}
	return aggregate.Product{}, product.ErrProductNotFound
}

func (mpr *MemoryProductRepository) Add(newprod aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _,ok := mpr.products[newprod.GetId()]; ok {
		return product.ErrProductAlreadyExists
	}

	// si no existe
	mpr.products[newprod.GetId()] = newprod

	return nil
}

func (mpr *MemoryProductRepository) Update(update aggregate.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _,ok := mpr.products[update.GetId()]; !ok {
		return product.ErrProductNotFound
	}

	mpr.products[update.GetId()] = update
	return nil
}

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(mpr.products, id)

	return nil
}