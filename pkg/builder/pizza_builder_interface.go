package builder

import "github.com/Jefschlarski/pizza-builder-ordering/pkg/menu"

// PizzaBuilderInterface é a interface para construir um pizza utilizando o padrao builder
type PizzaBuilderInterface interface {
	SetDough(dough string) PizzaBuilderInterface
	SetSauce(sauce string) PizzaBuilderInterface
	SetTopping(topping []string) PizzaBuilderInterface
	SetSize(size string) PizzaBuilderInterface
	GetPizza() (menu.Pizza, error)
	Reset() PizzaBuilderInterface
}
