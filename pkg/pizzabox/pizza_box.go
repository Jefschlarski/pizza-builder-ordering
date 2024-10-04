package pizzabox

import "github.com/Jefschlarski/pizza-builder-ordering/pkg/menu"

// PizzaBox representa um pizzabox
type PizzaBox struct {
	Pizza   menu.Pizza
	Drink   menu.Drink
	Dessert menu.Dessert
}
