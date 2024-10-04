package address

// Address representa um endereço
type Address struct {
	street string
	city   string
	number string
}

// New constroi um novo endereço
func New(street, city, number string) *Address {
	return &Address{
		street: street,
		city:   city,
		number: number,
	}
}
