package user

// User representa um usuario
type User struct {
	name string
}

// New constroi um novo usuario
func New(name string) *User {
	return &User{
		name: name,
	}
}
