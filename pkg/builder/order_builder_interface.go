package builder

import (
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/address"
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/order"
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/pizzabox"
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/user"
)

// OrderBuilderInterface é a interface para construir um pedido utilizando o padrao builder
type OrderBuilderInterface interface {
	SetCustomer(customer *user.User) OrderBuilderInterface
	SetPizzaBox(pizzaBox *pizzabox.PizzaBox) OrderBuilderInterface
	SetDeliveryAddress(address *address.Address) OrderBuilderInterface
	GetOrder() (order.Order, error)
	Reset() OrderBuilderInterface
}
