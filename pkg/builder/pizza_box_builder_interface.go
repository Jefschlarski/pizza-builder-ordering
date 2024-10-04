package builder

import (
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/menu"
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/pizzabox"
)

// PizzaBoxBuilderInterface implementa a interface PizzaBoxBuilderInterface para construir um pizzabox utilizando o padrao builder
type PizzaBoxBuilderInterface interface {
	SetPizza(pizza menu.Pizza) PizzaBoxBuilderInterface
	SetDrink(drink menu.Drink) PizzaBoxBuilderInterface
	SetDessert(dessert menu.Dessert) PizzaBoxBuilderInterface
	GetPizzaBox() (pizzabox.PizzaBox, error)
	Reset() PizzaBoxBuilderInterface
}
