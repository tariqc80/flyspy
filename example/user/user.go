package user

import (
)

// Model interface
type Model interface{
    Find(int) error
}

type User struct {
    data DataProvider
    name string
    email string
    password string
}

type DataProvider interface {
    Get(int) (Model, error)
    Store(Model) error
}

type UserProvider struct {
    // external database connection
}

func NewUser(p DataProvider) *User {
    return &User{
        data: p,
    }
}

func NewUserProvider() *UserProvider {
    return &UserProvider{}
}

func (u *User) Find(id int) error {

    model, err := u.data.Get(id)

    if err != nil {
        return err
    }

    m := model.(*User)

    u.name = m.name
    u.email = m.email
    u.password = m.password

    return nil
}

func (u *UserProvider) Get(id int) (*User, error) {
    return nil, nil
}

func (u UserProvider) Store(user *User) error {
    return nil
}
