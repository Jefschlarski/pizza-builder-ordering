package builder

import (
	"fmt"

	"github.com/Jefschlarski/pizza-builder-ordering/pkg/menu"
)

// PizzaBuilder implementa a interface PizzaBuilderInterface para construir uma pizza
type pizzaBuilder struct {
	pizza *menu.Pizza
}

// NewPizzaBuilder cria um novo builder para construir uma pizza
func NewPizzaBuilder() PizzaBuilderInterface {
	return &pizzaBuilder{
		pizza: &menu.Pizza{},
	}
}

// SetDough define a massa da pizza
func (b *pizzaBuilder) SetDough(dough string) PizzaBuilderInterface {
	b.pizza.Dough = dough
	return b
}

// SetSauce define o molho da pizza
func (b *pizzaBuilder) SetSauce(sauce string) PizzaBuilderInterface {
	b.pizza.Sauce = sauce
	return b
}

// SetTopping define os ingredientes da pizza
func (b *pizzaBuilder) SetTopping(topping []string) PizzaBuilderInterface {
	b.pizza.Topping = topping
	return b
}

// SetSize define o tamanho da pizza
func (b *pizzaBuilder) SetSize(size string) PizzaBuilderInterface {
	b.pizza.Size = size
	return b
}

// GetPizza retorna a pizza construída
func (b *pizzaBuilder) GetPizza() (menu.Pizza, error) {
	if b.pizza.Dough == "" || b.pizza.Sauce == "" || len(b.pizza.Topping) == 0 || b.pizza.Size == "" {
		return menu.Pizza{}, fmt.Errorf("pizza incompleta: um ou mais atributos estão faltando")
	}
	return *b.pizza, nil
}

// Reset redefine o builder
func (b *pizzaBuilder) Reset() PizzaBuilderInterface {
	b.pizza = &menu.Pizza{}
	return b
}
