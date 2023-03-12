package customer

import (
	"errors"
	"github.com/google/uuid"
	"github.com/guiddodg/ddd-go/aggregate"
)

var (
	ErrCustomerNotFOund     = errors.New("The customer was not found in repository")
	ErrFailterToAddCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer       = errors.New("failed to update the customer")
)

type CustomerRepository interface {
	Get(id uuid.UUID) (aggregate.Customer, error)
	Add(customer aggregate.Customer) error
	Update(customer aggregate.Customer) error
}
