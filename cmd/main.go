package main

import (
	"fmt"

	"github.com/Jefschlarski/pizza-builder-ordering/pkg/address"
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/builder"
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/menu"
	"github.com/Jefschlarski/pizza-builder-ordering/pkg/user"
)

func main() {
	// executa o programa n vezes
	qtdloop := 1
	for i := 1; i <= qtdloop; i++ {
		executar(i)
	}
}

func executar(i int) {
	fmt.Printf("Iniciando construção da pizza : %d\n", i)

	// Criando um builder e um diretor para a pizza
	PizzaBuilder := builder.NewPizzaBuilder()
	PizzaDirector := builder.NewPizzaDirector()

	dough := "massa fina"
	sauce := "molho de tomate"
	size := "grande"
	topping := []string{"mussarela", "oregano", "tomate"}

	// Setando o builder no diretor
	PizzaDirector.SetBuilder(PizzaBuilder)

	// Construindo a pizza
	PizzaDirector.ConstructPizza(dough, sauce, size, topping)

	// Recuperando a pizza construída
	pizza, err := PizzaBuilder.GetPizza()
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Pizza: %+v\n", pizza)
	}

	fmt.Printf("Finalizando construção da pizza: %d\n", i)
	// Resetando o builder da pizza
	PizzaBuilder.Reset()

	Drink := menu.Drink{Name: "agua"}
	Dessert := menu.Dessert{Name: "sorvete"}

	fmt.Printf("Iniciando construção da pizzabox: %d\n", i)

	// Criando um builder para a pizzabox
	PizzaBoxBuilder := builder.NewPizzaBoxBuilder()

	// Pizzabox não tem um diretor para ser construído

	// Construindo a pizzabox utilizando o builder diretamente
	PizzaBoxBuilder.SetPizza(pizza).SetDrink(Drink).SetDessert(Dessert)

	// Não utilizando o diretor para construir a pizzabox você pode escolher o que vai ser incluido
	// PizzaBoxBuilder.SetPizza(pizza)
	// PizzaBoxBuilder.SetPizza(pizza).SetDrink(Drink)
	// PizzaBoxBuilder.SetDrink(Drink).SetDessert(Dessert)

	// Recuperando o pizzabox construído
	pizzaBox, err := PizzaBoxBuilder.GetPizzaBox()
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("PizzaBox: %+v\n", pizzaBox)
	}

	fmt.Printf("Finalizando construção da pizzabox: %d\n", i)

	// Resetando o builder da pizzabox
	PizzaBoxBuilder.Reset()

	// Criando um customer e um endereço
	customer := user.New("Jeferson")
	deliveryAddress := address.New("Av. Teste", "São Jose", "122")

	fmt.Printf("Iniciando construção da ordem: %d\n", i)

	// Criando um builder e um diretor para a ordem
	OrderBuilder := builder.NewOrderBuilder()
	OrderDirector := builder.NewOrderDirector()

	// Setando o builder no diretor
	OrderDirector.SetBuilder(OrderBuilder)

	// Construindo a ordem
	OrderDirector.ConstructOrder(customer, &pizzaBox, deliveryAddress)

	// Recuperando a ordem
	order, err := OrderBuilder.GetOrder()
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Order: %+v\n", order)
	}

	fmt.Printf("Finalizando construção da ordem: %d\n", i)

	// Limpando o builder da ordem
	OrderBuilder.Reset()
}
