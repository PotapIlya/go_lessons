package user

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{
		name,
		age,
	}
}

func (u *User) SetName(name string) {
	u.Name = name
}
