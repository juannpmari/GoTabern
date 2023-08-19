// Package entities holds all the entities that are shared across subdomains
package entity

import "github.com/google/uuid"

// Person is an entity that represent a person in all domain
type Person struct {
	// ID is the identifier of the entity
	ID uuid.UUID
	Name string
	Age int
}