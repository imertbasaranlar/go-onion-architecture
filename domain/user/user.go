package user

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func NewUser(name, email, password string) (*User, error) {
	return &User{
		ID:       "11",
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}

func (u *User) ChangeEmail(email string) error {
	u.Email = email
	return nil
}
