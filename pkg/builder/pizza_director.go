package builder

// PizzaDirector é o responsável por construir um pizza com uma estrutura definida
type pizzaDirector struct {
	builder PizzaBuilderInterface
}

// NewPizzaDirector cria um novo director para construir uma pizza
func NewPizzaDirector() *pizzaDirector {
	return &pizzaDirector{}
}

// SetBuilder define o builder para o director
func (d *pizzaDirector) SetBuilder(b PizzaBuilderInterface) {
	d.builder = b
}

// ConstructPizza constroi uma pizza utilizando o builder
func (d *pizzaDirector) ConstructPizza(dough, sauce, size string, topping []string) {
	d.builder.SetDough(dough).SetSauce(sauce).SetTopping(topping).SetSize(size)
}
