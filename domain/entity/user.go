package entity

type User struct {
	id   int
	name string
}

func NewUser() *User {
	return &User{}
}

func (u *User) WithID(id int) *User {
	if u == nil {
		return nil
	}
	u.id = id
	return u
}

func (u *User) WithName(name string) *User {
	if u == nil {
		return nil
	}
	u.name = name
	return u
}

func (u *User) GetID() int {
	if u == nil {
		return 0
	}
	return u.id
}

func (u *User) GetName() string {
	if u == nil {
		return ""
	}
	return u.name
}
