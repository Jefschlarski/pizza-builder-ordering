package builder

import (
	"fmt"

	"github.com/Jefschlarski/pizza-builder-ordering/pkg/address"
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/order"
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/pizzabox"
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/user"
)

// OrderBuilder implementa a interface OrderBuilderInterface para construir um pedido
type OrderBuilder struct {
	order order.Order
}

// NewOrderBuilder cria um novo OrderBuilder
func NewOrderBuilder() OrderBuilderInterface {
	return &OrderBuilder{}
}

// SetCustomer define o cliente do pedido
func (b *OrderBuilder) SetCustomer(customer *user.User) OrderBuilderInterface {
	b.order.SetCustomer(customer)
	return b
}

// SetPizzaBox define o pizzabox do pedido
func (b *OrderBuilder) SetPizzaBox(pizzaBox *pizzabox.PizzaBox) OrderBuilderInterface {
	b.order.SetPizzaBox(pizzaBox)
	return b
}

// SetDeliveryAddress define o endereço de entrega do pedido
func (b *OrderBuilder) SetDeliveryAddress(address *address.Address) OrderBuilderInterface {
	b.order.SetAddress(address)
	return b
}

// GetOrder retorna o pedido construído
func (b *OrderBuilder) GetOrder() (order.Order, error) {
	if b.order.GetCustomer() == nil || b.order.GetPizzaBox() == nil || b.order.GetAddress() == nil {
		return order.Order{}, fmt.Errorf("ordem incompleta: um ou mais atributos estão faltando")
	}
	return b.order, nil
}

// Reset redefine o OrderBuilder
func (b *OrderBuilder) Reset() OrderBuilderInterface {
	b.order = order.Order{}
	return b
}
