package user

import (
)

// Model interface
type Model interface{}

type User struct {
    data DataProvider
    name string
    email string
    password string
}

type DataProvider interface {
    Get(int, Model) error
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

    var model User
    err := u.data.Get(id, &model)

    if err != nil {
        return err
    }

    u.name = model.name
    u.email = model.email
    u.password = model.password

    return nil
}

func (u *UserProvider) Get(id int, user *User) error {
    return nil
}

func (u UserProvider) Store(user *User) error {
    return nil
}
