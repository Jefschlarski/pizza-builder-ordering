
# pizza-builder-ordering

Um simples projeto que demonstra a implementação do padrão de projeto Builder em Go num cenário fictício.

Builder: Separa a construção de um objeto complexo da sua representação, permitindo a criação passo a passo.

## Descrição

Este projeto demonstra a implementação do padrão de projeto Builder em Go, aplicado a fins de exemplo a um simples sistema de gerenciamento de pedidos de uma pizzaria. O objetivo criar a pizza, o box e o pedido utilizando o padrão proposto.

## Funcionalidades
- **Personalização de Pizzas:** pode definir a massa, molho, cobertura e tamanho para criar uma pizza.
- **Construção de PizzaBox:** Permite a construção de uma "caixa de pizza" que seria um pacote que pode incluir um ou mais dos seguintes itens pizza, bebida e sobremesa.
- **Gerenciamento de Pedidos:** Cria um pedido completo com informações do cliente e endereço de entrega.
- **Padrão Builder:** Utiliza o padrão de projeto Builder para construir objetos de pizza, pizza box e pedidos de forma incremental e flexível.

## Estrutura do Projeto
- `main.go`: Contém a função principal e a lógica para simular as escolhas do cliente e construir a pizza, pizza box e o pedido.
- `pkg/address`: Contém a estrutura relacionada ao endereço de entrega.
- `pkg/builder`: Contém as interfaces e implementações dos builders e diretores para pizza, pizza box e pedidos.
- `pkg/menu`: Contém as estruturas para itens do menu como pizza, bebidas e sobremesas.
- `pkg/user`: Contém a estrutura relacionada ao cliente.

## Como Executar
1. Clone o repositório:
```bash
 git clone https://github.com/Jefschlarski/pizza-builder-ordering.git
```

2. Entre no diretório:
```bash
 cd pizza-builder-ordering
```

3. Execute o projeto:
```bash
  go run ./cmd/main.go
```