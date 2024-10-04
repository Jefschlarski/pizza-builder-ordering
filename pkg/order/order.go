package order

import (
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/address"
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/pizzabox"
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/user"
)

// Order representa um pedido
type Order struct {
	customer user.User
	pizzaBox pizzabox.PizzaBox
	address  address.Address
}

// GetCustomer retorna o cliente do pedido
func (o *Order) GetCustomer() *user.User {
	return &o.customer
}

// GetPizzaBox retorna o pizzabox do pedido
func (o *Order) GetPizzaBox() *pizzabox.PizzaBox {
	return &o.pizzaBox
}

// GetAddress retorna o endereço de entrega do pedido
func (o *Order) GetAddress() *address.Address {
	return &o.address
}

// SetCustomer define o cliente do pedido
func (o *Order) SetCustomer(customer *user.User) {
	o.customer = *customer
}

// SetPizzaBox define o pizzabox do pedido
func (o *Order) SetPizzaBox(pizzaBox *pizzabox.PizzaBox) {
	o.pizzaBox = *pizzaBox
}

// SetAddress define o endereço de entrega do pedido
func (o *Order) SetAddress(address *address.Address) {
	o.address = *address
}
