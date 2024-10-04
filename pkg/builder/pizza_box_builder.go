package builder

import (
	"fmt"

	"github.com/Jefschlarski/pizza-builder-ordering/pkg/menu"
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/pizzabox"
)

// PizzaBoxBuilderInterface implementa a interface PizzaBoxBuilderInterface para construir uma pizzabox
type pizzaBoxBuilder struct {
	pizzaBox *pizzabox.PizzaBox
}

// NewPizzaBoxBuilder constroi um novo PizzaBoxBuilder
func NewPizzaBoxBuilder() PizzaBoxBuilderInterface {
	return &pizzaBoxBuilder{
		pizzaBox: &pizzabox.PizzaBox{},
	}
}

// SetPizzaBox define a pizza do pizzabox
func (b *pizzaBoxBuilder) SetPizza(pizza menu.Pizza) PizzaBoxBuilderInterface {
	b.pizzaBox.Pizza = pizza
	return b
}

// SetDrink define o drink do pizzabox
func (b *pizzaBoxBuilder) SetDrink(drink menu.Drink) PizzaBoxBuilderInterface {
	b.pizzaBox.Drink = drink
	return b
}

// SetDessert define o dessert do pizzabox
func (b *pizzaBoxBuilder) SetDessert(dessert menu.Dessert) PizzaBoxBuilderInterface {
	b.pizzaBox.Dessert = dessert
	return b
}

// GetPizzaBox retorna o pizzabox construído
func (b *pizzaBoxBuilder) GetPizzaBox() (pizzabox.PizzaBox, error) {
	if b.pizzaBox == nil {
		return pizzabox.PizzaBox{}, fmt.Errorf("pizza box não pode ser nula")
	}

	return *b.pizzaBox, nil
}

// Reset redefine o PizzaBoxBuilder
func (b *pizzaBoxBuilder) Reset() PizzaBoxBuilderInterface {
	b.pizzaBox = &pizzabox.PizzaBox{}
	return b
}
