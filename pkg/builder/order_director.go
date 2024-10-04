package builder

import (
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/address"
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/pizzabox"
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/user"
)

// OrderDirector é o responsável por construir um pedido com uma estrutura definida
type OrderDirector struct {
	builder OrderBuilderInterface
}

// NewOrderDirector cria um novo OrderDirector
func NewOrderDirector() *OrderDirector {
	return &OrderDirector{}
}

// SetBuilder define o builder para o OrderDirector
func (d *OrderDirector) SetBuilder(b OrderBuilderInterface) {
	d.builder = b
}

// ConstructOrder constroi um pedido utilizando o builder
func (d *OrderDirector) ConstructOrder(customer *user.User, pizzaBox *pizzabox.PizzaBox, address *address.Address) {
	d.builder.SetCustomer(customer).SetPizzaBox(pizzaBox).SetDeliveryAddress(address)
}
